// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/shared"
	"github.com/tidwall/gjson"
)

// WebhookHeaderTimestampFormat is the format of the header X-Orb-Timestamp for webhook requests sent by Orb.
const WebhookHeaderTimestampFormat = "2006-01-02T15:04:05.999999999"

// WebhookService contains methods and other services that help with interacting
// with the Orb API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption

	// webhookSecret is the secret defined at the client level
	webhookSecret string
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts

	// This is a dummy response object. We need to build a request config to be able to check the webhook secret defined
	// at the client level.
	_res := struct{}{}
	cfg, err := requestconfig.NewRequestConfig(context.TODO(), http.MethodPost, "/webhooks", nil, &_res, opts...)
	if err != nil {
		panic(err)
	}

	if cfg.WebhookSecret != "" {
		r.webhookSecret = cfg.WebhookSecret
	}
	return
}

func (r *WebhookService) Unwrap(payload []byte, opts ...option.RequestOption) (*UnwrapWebhookEvent, error) {
	res := &UnwrapWebhookEvent{}
	err := res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Issued when a backfill is closed and its events are reflected into usage.
type BackfillReflectedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// A backfill represents an update to historical usage data, adding or replacing
	// events in a timeframe.
	Backfill BackfillReflectedWebhookEventBackfill `json:"backfill" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time   `json:"created_at" api:"required" format:"date-time"`
	Properties interface{} `json:"properties" api:"required"`
	// The event this payload describes.
	Type BackfillReflectedWebhookEventType `json:"type" api:"required"`
	JSON backfillReflectedWebhookEventJSON `json:"-"`
}

// backfillReflectedWebhookEventJSON contains the JSON metadata for the struct
// [BackfillReflectedWebhookEvent]
type backfillReflectedWebhookEventJSON struct {
	ID          apijson.Field
	Backfill    apijson.Field
	CreatedAt   apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BackfillReflectedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r backfillReflectedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r BackfillReflectedWebhookEvent) implementsUnwrapWebhookEvent() {}

// A backfill represents an update to historical usage data, adding or replacing
// events in a timeframe.
type BackfillReflectedWebhookEventBackfill struct {
	ID string `json:"id" api:"required"`
	// If in the future, the time at which the backfill will automatically close. If in
	// the past, the time at which the backfill was closed.
	CloseTime time.Time `json:"close_time" api:"required,nullable" format:"date-time"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The Orb-generated ID of the customer to which this backfill is scoped. If
	// `null`, this backfill is scoped to all customers.
	CustomerID string `json:"customer_id" api:"required,nullable"`
	// The number of events ingested in this backfill.
	EventsIngested int64 `json:"events_ingested" api:"required"`
	// If `true`, existing events in the backfill's timeframe will be replaced with the
	// newly ingested events associated with the backfill. If `false`, newly ingested
	// events will be added to the existing events.
	ReplaceExistingEvents bool `json:"replace_existing_events" api:"required"`
	// The time at which this backfill was reverted.
	RevertedAt time.Time `json:"reverted_at" api:"required,nullable" format:"date-time"`
	// The status of the backfill.
	Status         BackfillReflectedWebhookEventBackfillStatus `json:"status" api:"required"`
	TimeframeEnd   time.Time                                   `json:"timeframe_end" api:"required" format:"date-time"`
	TimeframeStart time.Time                                   `json:"timeframe_start" api:"required" format:"date-time"`
	// A boolean
	// [computed property](/extensibility/advanced-metrics#computed-properties) used to
	// filter the set of events to deprecate
	DeprecationFilter string                                    `json:"deprecation_filter" api:"nullable"`
	JSON              backfillReflectedWebhookEventBackfillJSON `json:"-"`
}

// backfillReflectedWebhookEventBackfillJSON contains the JSON metadata for the
// struct [BackfillReflectedWebhookEventBackfill]
type backfillReflectedWebhookEventBackfillJSON struct {
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

func (r *BackfillReflectedWebhookEventBackfill) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r backfillReflectedWebhookEventBackfillJSON) RawJSON() string {
	return r.raw
}

// The status of the backfill.
type BackfillReflectedWebhookEventBackfillStatus string

const (
	BackfillReflectedWebhookEventBackfillStatusPending       BackfillReflectedWebhookEventBackfillStatus = "pending"
	BackfillReflectedWebhookEventBackfillStatusReflected     BackfillReflectedWebhookEventBackfillStatus = "reflected"
	BackfillReflectedWebhookEventBackfillStatusPendingRevert BackfillReflectedWebhookEventBackfillStatus = "pending_revert"
	BackfillReflectedWebhookEventBackfillStatusReverted      BackfillReflectedWebhookEventBackfillStatus = "reverted"
)

func (r BackfillReflectedWebhookEventBackfillStatus) IsKnown() bool {
	switch r {
	case BackfillReflectedWebhookEventBackfillStatusPending, BackfillReflectedWebhookEventBackfillStatusReflected, BackfillReflectedWebhookEventBackfillStatusPendingRevert, BackfillReflectedWebhookEventBackfillStatusReverted:
		return true
	}
	return false
}

// The event this payload describes.
type BackfillReflectedWebhookEventType string

const (
	BackfillReflectedWebhookEventTypeBackfillReflected BackfillReflectedWebhookEventType = "backfill.reflected"
)

func (r BackfillReflectedWebhookEventType) IsKnown() bool {
	switch r {
	case BackfillReflectedWebhookEventTypeBackfillReflected:
		return true
	}
	return false
}

// Issued when a backfill is reverted, removing its events from usage.
type BackfillRevertedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// A backfill represents an update to historical usage data, adding or replacing
	// events in a timeframe.
	Backfill BackfillRevertedWebhookEventBackfill `json:"backfill" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time   `json:"created_at" api:"required" format:"date-time"`
	Properties interface{} `json:"properties" api:"required"`
	// The event this payload describes.
	Type BackfillRevertedWebhookEventType `json:"type" api:"required"`
	JSON backfillRevertedWebhookEventJSON `json:"-"`
}

// backfillRevertedWebhookEventJSON contains the JSON metadata for the struct
// [BackfillRevertedWebhookEvent]
type backfillRevertedWebhookEventJSON struct {
	ID          apijson.Field
	Backfill    apijson.Field
	CreatedAt   apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BackfillRevertedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r backfillRevertedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r BackfillRevertedWebhookEvent) implementsUnwrapWebhookEvent() {}

// A backfill represents an update to historical usage data, adding or replacing
// events in a timeframe.
type BackfillRevertedWebhookEventBackfill struct {
	ID string `json:"id" api:"required"`
	// If in the future, the time at which the backfill will automatically close. If in
	// the past, the time at which the backfill was closed.
	CloseTime time.Time `json:"close_time" api:"required,nullable" format:"date-time"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The Orb-generated ID of the customer to which this backfill is scoped. If
	// `null`, this backfill is scoped to all customers.
	CustomerID string `json:"customer_id" api:"required,nullable"`
	// The number of events ingested in this backfill.
	EventsIngested int64 `json:"events_ingested" api:"required"`
	// If `true`, existing events in the backfill's timeframe will be replaced with the
	// newly ingested events associated with the backfill. If `false`, newly ingested
	// events will be added to the existing events.
	ReplaceExistingEvents bool `json:"replace_existing_events" api:"required"`
	// The time at which this backfill was reverted.
	RevertedAt time.Time `json:"reverted_at" api:"required,nullable" format:"date-time"`
	// The status of the backfill.
	Status         BackfillRevertedWebhookEventBackfillStatus `json:"status" api:"required"`
	TimeframeEnd   time.Time                                  `json:"timeframe_end" api:"required" format:"date-time"`
	TimeframeStart time.Time                                  `json:"timeframe_start" api:"required" format:"date-time"`
	// A boolean
	// [computed property](/extensibility/advanced-metrics#computed-properties) used to
	// filter the set of events to deprecate
	DeprecationFilter string                                   `json:"deprecation_filter" api:"nullable"`
	JSON              backfillRevertedWebhookEventBackfillJSON `json:"-"`
}

// backfillRevertedWebhookEventBackfillJSON contains the JSON metadata for the
// struct [BackfillRevertedWebhookEventBackfill]
type backfillRevertedWebhookEventBackfillJSON struct {
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

func (r *BackfillRevertedWebhookEventBackfill) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r backfillRevertedWebhookEventBackfillJSON) RawJSON() string {
	return r.raw
}

// The status of the backfill.
type BackfillRevertedWebhookEventBackfillStatus string

const (
	BackfillRevertedWebhookEventBackfillStatusPending       BackfillRevertedWebhookEventBackfillStatus = "pending"
	BackfillRevertedWebhookEventBackfillStatusReflected     BackfillRevertedWebhookEventBackfillStatus = "reflected"
	BackfillRevertedWebhookEventBackfillStatusPendingRevert BackfillRevertedWebhookEventBackfillStatus = "pending_revert"
	BackfillRevertedWebhookEventBackfillStatusReverted      BackfillRevertedWebhookEventBackfillStatus = "reverted"
)

func (r BackfillRevertedWebhookEventBackfillStatus) IsKnown() bool {
	switch r {
	case BackfillRevertedWebhookEventBackfillStatusPending, BackfillRevertedWebhookEventBackfillStatusReflected, BackfillRevertedWebhookEventBackfillStatusPendingRevert, BackfillRevertedWebhookEventBackfillStatusReverted:
		return true
	}
	return false
}

// The event this payload describes.
type BackfillRevertedWebhookEventType string

const (
	BackfillRevertedWebhookEventTypeBackfillReverted BackfillRevertedWebhookEventType = "backfill.reverted"
)

func (r BackfillRevertedWebhookEventType) IsKnown() bool {
	switch r {
	case BackfillRevertedWebhookEventTypeBackfillReverted:
		return true
	}
	return false
}

// Issued when a billable metric is edited.
type BillableMetricEditedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The Metric resource represents a calculation of a quantity based on events.
	// Metrics are defined by the query that transforms raw usage events into
	// meaningful values for your customers.
	BillableMetric BillableMetric `json:"billable_metric" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                  `json:"created_at" api:"required" format:"date-time"`
	Properties BillableMetricEditedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type BillableMetricEditedWebhookEventType `json:"type" api:"required"`
	JSON billableMetricEditedWebhookEventJSON `json:"-"`
}

// billableMetricEditedWebhookEventJSON contains the JSON metadata for the struct
// [BillableMetricEditedWebhookEvent]
type billableMetricEditedWebhookEventJSON struct {
	ID             apijson.Field
	BillableMetric apijson.Field
	CreatedAt      apijson.Field
	Properties     apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BillableMetricEditedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billableMetricEditedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r BillableMetricEditedWebhookEvent) implementsUnwrapWebhookEvent() {}

type BillableMetricEditedWebhookEventProperties struct {
	// metadata values are non-null on the wire, as on the price event.
	PreviousAttributes BillableMetricEditedWebhookEventPropertiesPreviousAttributes `json:"previous_attributes" api:"required"`
	JSON               billableMetricEditedWebhookEventPropertiesJSON               `json:"-"`
}

// billableMetricEditedWebhookEventPropertiesJSON contains the JSON metadata for
// the struct [BillableMetricEditedWebhookEventProperties]
type billableMetricEditedWebhookEventPropertiesJSON struct {
	PreviousAttributes apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BillableMetricEditedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billableMetricEditedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// metadata values are non-null on the wire, as on the price event.
type BillableMetricEditedWebhookEventPropertiesPreviousAttributes struct {
	Description string                                                           `json:"description" api:"nullable"`
	Metadata    map[string]string                                                `json:"metadata" api:"nullable"`
	Name        string                                                           `json:"name" api:"nullable"`
	JSON        billableMetricEditedWebhookEventPropertiesPreviousAttributesJSON `json:"-"`
}

// billableMetricEditedWebhookEventPropertiesPreviousAttributesJSON contains the
// JSON metadata for the struct
// [BillableMetricEditedWebhookEventPropertiesPreviousAttributes]
type billableMetricEditedWebhookEventPropertiesPreviousAttributesJSON struct {
	Description apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillableMetricEditedWebhookEventPropertiesPreviousAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billableMetricEditedWebhookEventPropertiesPreviousAttributesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type BillableMetricEditedWebhookEventType string

const (
	BillableMetricEditedWebhookEventTypeBillableMetricEdited BillableMetricEditedWebhookEventType = "billable_metric.edited"
)

func (r BillableMetricEditedWebhookEventType) IsKnown() bool {
	switch r {
	case BillableMetricEditedWebhookEventTypeBillableMetricEdited:
		return true
	}
	return false
}

// Issued when a credit block accounting sync fails.
type CreditBlockAccountingSyncFailedWebhookEvent struct {
	// The ID of this webhook event.
	ID                   string                                                          `json:"id" api:"required"`
	AccountingSyncRecord CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecord `json:"accounting_sync_record" api:"required"`
	// The Credit Block resource models prepaid credits within Orb.
	Block CreditBlockAccountingSyncFailedWebhookEventBlock `json:"block" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                             `json:"created_at" api:"required" format:"date-time"`
	Properties CreditBlockAccountingSyncFailedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type CreditBlockAccountingSyncFailedWebhookEventType `json:"type" api:"required"`
	JSON creditBlockAccountingSyncFailedWebhookEventJSON `json:"-"`
}

// creditBlockAccountingSyncFailedWebhookEventJSON contains the JSON metadata for
// the struct [CreditBlockAccountingSyncFailedWebhookEvent]
type creditBlockAccountingSyncFailedWebhookEventJSON struct {
	ID                   apijson.Field
	AccountingSyncRecord apijson.Field
	Block                apijson.Field
	CreatedAt            apijson.Field
	Properties           apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CreditBlockAccountingSyncFailedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockAccountingSyncFailedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CreditBlockAccountingSyncFailedWebhookEvent) implementsUnwrapWebhookEvent() {}

type CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecord struct {
	ID                 string                                                                    `json:"id" api:"required"`
	CustomerID         string                                                                    `json:"customer_id" api:"required"`
	RecordType         CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType `json:"record_type" api:"required"`
	BlockID            string                                                                    `json:"block_id" api:"nullable"`
	ErrorDetails       map[string]interface{}                                                    `json:"error_details" api:"nullable"`
	InvoiceID          string                                                                    `json:"invoice_id" api:"nullable"`
	ProviderCustomerID string                                                                    `json:"provider_customer_id" api:"nullable"`
	Status             string                                                                    `json:"status" api:"nullable"`
	SyncAction         string                                                                    `json:"sync_action" api:"nullable"`
	JSON               creditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordJSON       `json:"-"`
}

// creditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordJSON contains the
// JSON metadata for the struct
// [CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecord]
type creditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordJSON struct {
	ID                 apijson.Field
	CustomerID         apijson.Field
	RecordType         apijson.Field
	BlockID            apijson.Field
	ErrorDetails       apijson.Field
	InvoiceID          apijson.Field
	ProviderCustomerID apijson.Field
	Status             apijson.Field
	SyncAction         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordJSON) RawJSON() string {
	return r.raw
}

type CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType string

const (
	CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomer                   CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "customer"
	CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeInvoice                    CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "invoice"
	CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeTransaction                CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "transaction"
	CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "customer_balance_transaction"
	CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCreditNote                 CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "credit_note"
	CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSubscription               CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "subscription"
	CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSalesOrder                 CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "sales_order"
	CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeBlock                      CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "block"
)

func (r CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType) IsKnown() bool {
	switch r {
	case CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomer, CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeInvoice, CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeTransaction, CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction, CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCreditNote, CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSubscription, CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSalesOrder, CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeBlock:
		return true
	}
	return false
}

// The Credit Block resource models prepaid credits within Orb.
type CreditBlockAccountingSyncFailedWebhookEventBlock struct {
	ID      string `json:"id" api:"required"`
	Balance string `json:"balance" api:"required"`
	// How this credit block was created: `allocation` (a subscription's recurring
	// credit allocation), `top_up` (an automatic balance-threshold top-up),
	// `commitment` (a subscription commitment true-up rolled forward as credit), or
	// `manual` (a manual credit ledger increment, including credits voided or expired
	// off another block).
	CreditBlockSource     CreditBlockAccountingSyncFailedWebhookEventBlockCreditBlockSource `json:"credit_block_source" api:"required"`
	EffectiveDate         time.Time                                                         `json:"effective_date" api:"required,nullable" format:"date-time"`
	ExpiryDate            time.Time                                                         `json:"expiry_date" api:"required,nullable" format:"date-time"`
	Filters               []CreditBlockAccountingSyncFailedWebhookEventBlockFilter          `json:"filters" api:"required"`
	MaximumInitialBalance string                                                            `json:"maximum_initial_balance" api:"required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata         map[string]string                                      `json:"metadata" api:"required"`
	PerUnitCostBasis string                                                 `json:"per_unit_cost_basis" api:"required,nullable"`
	Status           CreditBlockAccountingSyncFailedWebhookEventBlockStatus `json:"status" api:"required"`
	// The credit allocation that funded a block. Extends the allocation resource
	// serialized on prices with the catalog-item attribution of the funding price.
	CreditAllocation CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocation `json:"credit_allocation" api:"nullable"`
	// The subscription commitment whose true-up rolled forward into this credit block.
	// Present only when `credit_block_source` is `commitment`.
	CreditCommitment CreditBlockAccountingSyncFailedWebhookEventBlockCreditCommitment `json:"credit_commitment" api:"nullable"`
	JSON             creditBlockAccountingSyncFailedWebhookEventBlockJSON             `json:"-"`
}

// creditBlockAccountingSyncFailedWebhookEventBlockJSON contains the JSON metadata
// for the struct [CreditBlockAccountingSyncFailedWebhookEventBlock]
type creditBlockAccountingSyncFailedWebhookEventBlockJSON struct {
	ID                    apijson.Field
	Balance               apijson.Field
	CreditBlockSource     apijson.Field
	EffectiveDate         apijson.Field
	ExpiryDate            apijson.Field
	Filters               apijson.Field
	MaximumInitialBalance apijson.Field
	Metadata              apijson.Field
	PerUnitCostBasis      apijson.Field
	Status                apijson.Field
	CreditAllocation      apijson.Field
	CreditCommitment      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CreditBlockAccountingSyncFailedWebhookEventBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockAccountingSyncFailedWebhookEventBlockJSON) RawJSON() string {
	return r.raw
}

// How this credit block was created: `allocation` (a subscription's recurring
// credit allocation), `top_up` (an automatic balance-threshold top-up),
// `commitment` (a subscription commitment true-up rolled forward as credit), or
// `manual` (a manual credit ledger increment, including credits voided or expired
// off another block).
type CreditBlockAccountingSyncFailedWebhookEventBlockCreditBlockSource string

const (
	CreditBlockAccountingSyncFailedWebhookEventBlockCreditBlockSourceAllocation CreditBlockAccountingSyncFailedWebhookEventBlockCreditBlockSource = "allocation"
	CreditBlockAccountingSyncFailedWebhookEventBlockCreditBlockSourceTopUp      CreditBlockAccountingSyncFailedWebhookEventBlockCreditBlockSource = "top_up"
	CreditBlockAccountingSyncFailedWebhookEventBlockCreditBlockSourceCommitment CreditBlockAccountingSyncFailedWebhookEventBlockCreditBlockSource = "commitment"
	CreditBlockAccountingSyncFailedWebhookEventBlockCreditBlockSourceManual     CreditBlockAccountingSyncFailedWebhookEventBlockCreditBlockSource = "manual"
)

func (r CreditBlockAccountingSyncFailedWebhookEventBlockCreditBlockSource) IsKnown() bool {
	switch r {
	case CreditBlockAccountingSyncFailedWebhookEventBlockCreditBlockSourceAllocation, CreditBlockAccountingSyncFailedWebhookEventBlockCreditBlockSourceTopUp, CreditBlockAccountingSyncFailedWebhookEventBlockCreditBlockSourceCommitment, CreditBlockAccountingSyncFailedWebhookEventBlockCreditBlockSourceManual:
		return true
	}
	return false
}

type CreditBlockAccountingSyncFailedWebhookEventBlockFilter struct {
	// The property of the price to filter on.
	Field CreditBlockAccountingSyncFailedWebhookEventBlockFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator CreditBlockAccountingSyncFailedWebhookEventBlockFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                   `json:"values" api:"required"`
	JSON   creditBlockAccountingSyncFailedWebhookEventBlockFilterJSON `json:"-"`
}

// creditBlockAccountingSyncFailedWebhookEventBlockFilterJSON contains the JSON
// metadata for the struct [CreditBlockAccountingSyncFailedWebhookEventBlockFilter]
type creditBlockAccountingSyncFailedWebhookEventBlockFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditBlockAccountingSyncFailedWebhookEventBlockFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockAccountingSyncFailedWebhookEventBlockFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type CreditBlockAccountingSyncFailedWebhookEventBlockFiltersField string

const (
	CreditBlockAccountingSyncFailedWebhookEventBlockFiltersFieldPriceID       CreditBlockAccountingSyncFailedWebhookEventBlockFiltersField = "price_id"
	CreditBlockAccountingSyncFailedWebhookEventBlockFiltersFieldItemID        CreditBlockAccountingSyncFailedWebhookEventBlockFiltersField = "item_id"
	CreditBlockAccountingSyncFailedWebhookEventBlockFiltersFieldPriceType     CreditBlockAccountingSyncFailedWebhookEventBlockFiltersField = "price_type"
	CreditBlockAccountingSyncFailedWebhookEventBlockFiltersFieldCurrency      CreditBlockAccountingSyncFailedWebhookEventBlockFiltersField = "currency"
	CreditBlockAccountingSyncFailedWebhookEventBlockFiltersFieldPricingUnitID CreditBlockAccountingSyncFailedWebhookEventBlockFiltersField = "pricing_unit_id"
)

func (r CreditBlockAccountingSyncFailedWebhookEventBlockFiltersField) IsKnown() bool {
	switch r {
	case CreditBlockAccountingSyncFailedWebhookEventBlockFiltersFieldPriceID, CreditBlockAccountingSyncFailedWebhookEventBlockFiltersFieldItemID, CreditBlockAccountingSyncFailedWebhookEventBlockFiltersFieldPriceType, CreditBlockAccountingSyncFailedWebhookEventBlockFiltersFieldCurrency, CreditBlockAccountingSyncFailedWebhookEventBlockFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type CreditBlockAccountingSyncFailedWebhookEventBlockFiltersOperator string

const (
	CreditBlockAccountingSyncFailedWebhookEventBlockFiltersOperatorIncludes CreditBlockAccountingSyncFailedWebhookEventBlockFiltersOperator = "includes"
	CreditBlockAccountingSyncFailedWebhookEventBlockFiltersOperatorExcludes CreditBlockAccountingSyncFailedWebhookEventBlockFiltersOperator = "excludes"
)

func (r CreditBlockAccountingSyncFailedWebhookEventBlockFiltersOperator) IsKnown() bool {
	switch r {
	case CreditBlockAccountingSyncFailedWebhookEventBlockFiltersOperatorIncludes, CreditBlockAccountingSyncFailedWebhookEventBlockFiltersOperatorExcludes:
		return true
	}
	return false
}

type CreditBlockAccountingSyncFailedWebhookEventBlockStatus string

const (
	CreditBlockAccountingSyncFailedWebhookEventBlockStatusActive         CreditBlockAccountingSyncFailedWebhookEventBlockStatus = "active"
	CreditBlockAccountingSyncFailedWebhookEventBlockStatusPendingPayment CreditBlockAccountingSyncFailedWebhookEventBlockStatus = "pending_payment"
)

func (r CreditBlockAccountingSyncFailedWebhookEventBlockStatus) IsKnown() bool {
	switch r {
	case CreditBlockAccountingSyncFailedWebhookEventBlockStatusActive, CreditBlockAccountingSyncFailedWebhookEventBlockStatusPendingPayment:
		return true
	}
	return false
}

// The credit allocation that funded a block. Extends the allocation resource
// serialized on prices with the catalog-item attribution of the funding price.
type CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocation struct {
	AllowsRollover   bool                    `json:"allows_rollover" api:"required"`
	Currency         string                  `json:"currency" api:"required"`
	CustomExpiration shared.CustomExpiration `json:"custom_expiration" api:"required,nullable"`
	// The ID of the catalog item this block was allocated from, derived from the
	// allocation's price.
	ItemID        string                                                                   `json:"item_id" api:"required"`
	Filters       []CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFilter `json:"filters"`
	LicenseTypeID string                                                                   `json:"license_type_id" api:"nullable"`
	JSON          creditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationJSON     `json:"-"`
}

// creditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationJSON contains
// the JSON metadata for the struct
// [CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocation]
type creditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationJSON struct {
	AllowsRollover   apijson.Field
	Currency         apijson.Field
	CustomExpiration apijson.Field
	ItemID           apijson.Field
	Filters          apijson.Field
	LicenseTypeID    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFilter struct {
	// The property of the price to filter on.
	Field CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                                   `json:"values" api:"required"`
	JSON   creditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFilterJSON `json:"-"`
}

// creditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFilterJSON
// contains the JSON metadata for the struct
// [CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFilter]
type creditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersField string

const (
	CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersFieldPriceID       CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersField = "price_id"
	CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersFieldItemID        CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersField = "item_id"
	CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersFieldPriceType     CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersField = "price_type"
	CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersFieldCurrency      CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersField = "currency"
	CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersFieldPricingUnitID CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersField = "pricing_unit_id"
)

func (r CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersField) IsKnown() bool {
	switch r {
	case CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersFieldPriceID, CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersFieldItemID, CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersFieldPriceType, CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersFieldCurrency, CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersOperator string

const (
	CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersOperatorIncludes CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersOperator = "includes"
	CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersOperatorExcludes CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersOperator = "excludes"
)

func (r CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersOperator) IsKnown() bool {
	switch r {
	case CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersOperatorIncludes, CreditBlockAccountingSyncFailedWebhookEventBlockCreditAllocationFiltersOperatorExcludes:
		return true
	}
	return false
}

// The subscription commitment whose true-up rolled forward into this credit block.
// Present only when `credit_block_source` is `commitment`.
type CreditBlockAccountingSyncFailedWebhookEventBlockCreditCommitment struct {
	// The ID of the subscription commitment this block was rolled forward from.
	ID string `json:"id" api:"required"`
	// The subscription the commitment belongs to.
	SubscriptionID string                                                               `json:"subscription_id" api:"nullable"`
	JSON           creditBlockAccountingSyncFailedWebhookEventBlockCreditCommitmentJSON `json:"-"`
}

// creditBlockAccountingSyncFailedWebhookEventBlockCreditCommitmentJSON contains
// the JSON metadata for the struct
// [CreditBlockAccountingSyncFailedWebhookEventBlockCreditCommitment]
type creditBlockAccountingSyncFailedWebhookEventBlockCreditCommitmentJSON struct {
	ID             apijson.Field
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CreditBlockAccountingSyncFailedWebhookEventBlockCreditCommitment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockAccountingSyncFailedWebhookEventBlockCreditCommitmentJSON) RawJSON() string {
	return r.raw
}

type CreditBlockAccountingSyncFailedWebhookEventProperties struct {
	ConnectionType string                                                    `json:"connection_type" api:"required"`
	FailureReason  string                                                    `json:"failure_reason" api:"required"`
	JSON           creditBlockAccountingSyncFailedWebhookEventPropertiesJSON `json:"-"`
}

// creditBlockAccountingSyncFailedWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [CreditBlockAccountingSyncFailedWebhookEventProperties]
type creditBlockAccountingSyncFailedWebhookEventPropertiesJSON struct {
	ConnectionType apijson.Field
	FailureReason  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CreditBlockAccountingSyncFailedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockAccountingSyncFailedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type CreditBlockAccountingSyncFailedWebhookEventType string

const (
	CreditBlockAccountingSyncFailedWebhookEventTypeCreditBlockAccountingSyncFailed CreditBlockAccountingSyncFailedWebhookEventType = "credit_block.accounting_sync_failed"
)

func (r CreditBlockAccountingSyncFailedWebhookEventType) IsKnown() bool {
	switch r {
	case CreditBlockAccountingSyncFailedWebhookEventTypeCreditBlockAccountingSyncFailed:
		return true
	}
	return false
}

// Issued when a credit block accounting sync succeeds.
type CreditBlockAccountingSyncSucceededWebhookEvent struct {
	// The ID of this webhook event.
	ID                   string                                                             `json:"id" api:"required"`
	AccountingSyncRecord CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecord `json:"accounting_sync_record" api:"required"`
	// The Credit Block resource models prepaid credits within Orb.
	Block CreditBlockAccountingSyncSucceededWebhookEventBlock `json:"block" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                                `json:"created_at" api:"required" format:"date-time"`
	Properties CreditBlockAccountingSyncSucceededWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type CreditBlockAccountingSyncSucceededWebhookEventType `json:"type" api:"required"`
	JSON creditBlockAccountingSyncSucceededWebhookEventJSON `json:"-"`
}

// creditBlockAccountingSyncSucceededWebhookEventJSON contains the JSON metadata
// for the struct [CreditBlockAccountingSyncSucceededWebhookEvent]
type creditBlockAccountingSyncSucceededWebhookEventJSON struct {
	ID                   apijson.Field
	AccountingSyncRecord apijson.Field
	Block                apijson.Field
	CreatedAt            apijson.Field
	Properties           apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CreditBlockAccountingSyncSucceededWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockAccountingSyncSucceededWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CreditBlockAccountingSyncSucceededWebhookEvent) implementsUnwrapWebhookEvent() {}

type CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecord struct {
	ID                 string                                                                       `json:"id" api:"required"`
	CustomerID         string                                                                       `json:"customer_id" api:"required"`
	RecordType         CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType `json:"record_type" api:"required"`
	BlockID            string                                                                       `json:"block_id" api:"nullable"`
	ErrorDetails       map[string]interface{}                                                       `json:"error_details" api:"nullable"`
	InvoiceID          string                                                                       `json:"invoice_id" api:"nullable"`
	ProviderCustomerID string                                                                       `json:"provider_customer_id" api:"nullable"`
	Status             string                                                                       `json:"status" api:"nullable"`
	SyncAction         string                                                                       `json:"sync_action" api:"nullable"`
	JSON               creditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON       `json:"-"`
}

// creditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON contains
// the JSON metadata for the struct
// [CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecord]
type creditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON struct {
	ID                 apijson.Field
	CustomerID         apijson.Field
	RecordType         apijson.Field
	BlockID            apijson.Field
	ErrorDetails       apijson.Field
	InvoiceID          apijson.Field
	ProviderCustomerID apijson.Field
	Status             apijson.Field
	SyncAction         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON) RawJSON() string {
	return r.raw
}

type CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType string

const (
	CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomer                   CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "customer"
	CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeInvoice                    CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "invoice"
	CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeTransaction                CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "transaction"
	CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "customer_balance_transaction"
	CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCreditNote                 CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "credit_note"
	CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSubscription               CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "subscription"
	CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSalesOrder                 CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "sales_order"
	CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeBlock                      CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "block"
)

func (r CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType) IsKnown() bool {
	switch r {
	case CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomer, CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeInvoice, CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeTransaction, CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction, CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCreditNote, CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSubscription, CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSalesOrder, CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeBlock:
		return true
	}
	return false
}

// The Credit Block resource models prepaid credits within Orb.
type CreditBlockAccountingSyncSucceededWebhookEventBlock struct {
	ID      string `json:"id" api:"required"`
	Balance string `json:"balance" api:"required"`
	// How this credit block was created: `allocation` (a subscription's recurring
	// credit allocation), `top_up` (an automatic balance-threshold top-up),
	// `commitment` (a subscription commitment true-up rolled forward as credit), or
	// `manual` (a manual credit ledger increment, including credits voided or expired
	// off another block).
	CreditBlockSource     CreditBlockAccountingSyncSucceededWebhookEventBlockCreditBlockSource `json:"credit_block_source" api:"required"`
	EffectiveDate         time.Time                                                            `json:"effective_date" api:"required,nullable" format:"date-time"`
	ExpiryDate            time.Time                                                            `json:"expiry_date" api:"required,nullable" format:"date-time"`
	Filters               []CreditBlockAccountingSyncSucceededWebhookEventBlockFilter          `json:"filters" api:"required"`
	MaximumInitialBalance string                                                               `json:"maximum_initial_balance" api:"required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata         map[string]string                                         `json:"metadata" api:"required"`
	PerUnitCostBasis string                                                    `json:"per_unit_cost_basis" api:"required,nullable"`
	Status           CreditBlockAccountingSyncSucceededWebhookEventBlockStatus `json:"status" api:"required"`
	// The credit allocation that funded a block. Extends the allocation resource
	// serialized on prices with the catalog-item attribution of the funding price.
	CreditAllocation CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocation `json:"credit_allocation" api:"nullable"`
	// The subscription commitment whose true-up rolled forward into this credit block.
	// Present only when `credit_block_source` is `commitment`.
	CreditCommitment CreditBlockAccountingSyncSucceededWebhookEventBlockCreditCommitment `json:"credit_commitment" api:"nullable"`
	JSON             creditBlockAccountingSyncSucceededWebhookEventBlockJSON             `json:"-"`
}

// creditBlockAccountingSyncSucceededWebhookEventBlockJSON contains the JSON
// metadata for the struct [CreditBlockAccountingSyncSucceededWebhookEventBlock]
type creditBlockAccountingSyncSucceededWebhookEventBlockJSON struct {
	ID                    apijson.Field
	Balance               apijson.Field
	CreditBlockSource     apijson.Field
	EffectiveDate         apijson.Field
	ExpiryDate            apijson.Field
	Filters               apijson.Field
	MaximumInitialBalance apijson.Field
	Metadata              apijson.Field
	PerUnitCostBasis      apijson.Field
	Status                apijson.Field
	CreditAllocation      apijson.Field
	CreditCommitment      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CreditBlockAccountingSyncSucceededWebhookEventBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockAccountingSyncSucceededWebhookEventBlockJSON) RawJSON() string {
	return r.raw
}

// How this credit block was created: `allocation` (a subscription's recurring
// credit allocation), `top_up` (an automatic balance-threshold top-up),
// `commitment` (a subscription commitment true-up rolled forward as credit), or
// `manual` (a manual credit ledger increment, including credits voided or expired
// off another block).
type CreditBlockAccountingSyncSucceededWebhookEventBlockCreditBlockSource string

const (
	CreditBlockAccountingSyncSucceededWebhookEventBlockCreditBlockSourceAllocation CreditBlockAccountingSyncSucceededWebhookEventBlockCreditBlockSource = "allocation"
	CreditBlockAccountingSyncSucceededWebhookEventBlockCreditBlockSourceTopUp      CreditBlockAccountingSyncSucceededWebhookEventBlockCreditBlockSource = "top_up"
	CreditBlockAccountingSyncSucceededWebhookEventBlockCreditBlockSourceCommitment CreditBlockAccountingSyncSucceededWebhookEventBlockCreditBlockSource = "commitment"
	CreditBlockAccountingSyncSucceededWebhookEventBlockCreditBlockSourceManual     CreditBlockAccountingSyncSucceededWebhookEventBlockCreditBlockSource = "manual"
)

func (r CreditBlockAccountingSyncSucceededWebhookEventBlockCreditBlockSource) IsKnown() bool {
	switch r {
	case CreditBlockAccountingSyncSucceededWebhookEventBlockCreditBlockSourceAllocation, CreditBlockAccountingSyncSucceededWebhookEventBlockCreditBlockSourceTopUp, CreditBlockAccountingSyncSucceededWebhookEventBlockCreditBlockSourceCommitment, CreditBlockAccountingSyncSucceededWebhookEventBlockCreditBlockSourceManual:
		return true
	}
	return false
}

type CreditBlockAccountingSyncSucceededWebhookEventBlockFilter struct {
	// The property of the price to filter on.
	Field CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                      `json:"values" api:"required"`
	JSON   creditBlockAccountingSyncSucceededWebhookEventBlockFilterJSON `json:"-"`
}

// creditBlockAccountingSyncSucceededWebhookEventBlockFilterJSON contains the JSON
// metadata for the struct
// [CreditBlockAccountingSyncSucceededWebhookEventBlockFilter]
type creditBlockAccountingSyncSucceededWebhookEventBlockFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditBlockAccountingSyncSucceededWebhookEventBlockFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockAccountingSyncSucceededWebhookEventBlockFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersField string

const (
	CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersFieldPriceID       CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersField = "price_id"
	CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersFieldItemID        CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersField = "item_id"
	CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersFieldPriceType     CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersField = "price_type"
	CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersFieldCurrency      CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersField = "currency"
	CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersFieldPricingUnitID CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersField = "pricing_unit_id"
)

func (r CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersField) IsKnown() bool {
	switch r {
	case CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersFieldPriceID, CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersFieldItemID, CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersFieldPriceType, CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersFieldCurrency, CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersOperator string

const (
	CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersOperatorIncludes CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersOperator = "includes"
	CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersOperatorExcludes CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersOperator = "excludes"
)

func (r CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersOperator) IsKnown() bool {
	switch r {
	case CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersOperatorIncludes, CreditBlockAccountingSyncSucceededWebhookEventBlockFiltersOperatorExcludes:
		return true
	}
	return false
}

type CreditBlockAccountingSyncSucceededWebhookEventBlockStatus string

const (
	CreditBlockAccountingSyncSucceededWebhookEventBlockStatusActive         CreditBlockAccountingSyncSucceededWebhookEventBlockStatus = "active"
	CreditBlockAccountingSyncSucceededWebhookEventBlockStatusPendingPayment CreditBlockAccountingSyncSucceededWebhookEventBlockStatus = "pending_payment"
)

func (r CreditBlockAccountingSyncSucceededWebhookEventBlockStatus) IsKnown() bool {
	switch r {
	case CreditBlockAccountingSyncSucceededWebhookEventBlockStatusActive, CreditBlockAccountingSyncSucceededWebhookEventBlockStatusPendingPayment:
		return true
	}
	return false
}

// The credit allocation that funded a block. Extends the allocation resource
// serialized on prices with the catalog-item attribution of the funding price.
type CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocation struct {
	AllowsRollover   bool                    `json:"allows_rollover" api:"required"`
	Currency         string                  `json:"currency" api:"required"`
	CustomExpiration shared.CustomExpiration `json:"custom_expiration" api:"required,nullable"`
	// The ID of the catalog item this block was allocated from, derived from the
	// allocation's price.
	ItemID        string                                                                      `json:"item_id" api:"required"`
	Filters       []CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFilter `json:"filters"`
	LicenseTypeID string                                                                      `json:"license_type_id" api:"nullable"`
	JSON          creditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationJSON     `json:"-"`
}

// creditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationJSON contains
// the JSON metadata for the struct
// [CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocation]
type creditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationJSON struct {
	AllowsRollover   apijson.Field
	Currency         apijson.Field
	CustomExpiration apijson.Field
	ItemID           apijson.Field
	Filters          apijson.Field
	LicenseTypeID    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFilter struct {
	// The property of the price to filter on.
	Field CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                                      `json:"values" api:"required"`
	JSON   creditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFilterJSON `json:"-"`
}

// creditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFilterJSON
// contains the JSON metadata for the struct
// [CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFilter]
type creditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersField string

const (
	CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersFieldPriceID       CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersField = "price_id"
	CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersFieldItemID        CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersField = "item_id"
	CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersFieldPriceType     CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersField = "price_type"
	CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersFieldCurrency      CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersField = "currency"
	CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersFieldPricingUnitID CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersField = "pricing_unit_id"
)

func (r CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersField) IsKnown() bool {
	switch r {
	case CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersFieldPriceID, CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersFieldItemID, CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersFieldPriceType, CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersFieldCurrency, CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersOperator string

const (
	CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersOperatorIncludes CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersOperator = "includes"
	CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersOperatorExcludes CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersOperator = "excludes"
)

func (r CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersOperator) IsKnown() bool {
	switch r {
	case CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersOperatorIncludes, CreditBlockAccountingSyncSucceededWebhookEventBlockCreditAllocationFiltersOperatorExcludes:
		return true
	}
	return false
}

// The subscription commitment whose true-up rolled forward into this credit block.
// Present only when `credit_block_source` is `commitment`.
type CreditBlockAccountingSyncSucceededWebhookEventBlockCreditCommitment struct {
	// The ID of the subscription commitment this block was rolled forward from.
	ID string `json:"id" api:"required"`
	// The subscription the commitment belongs to.
	SubscriptionID string                                                                  `json:"subscription_id" api:"nullable"`
	JSON           creditBlockAccountingSyncSucceededWebhookEventBlockCreditCommitmentJSON `json:"-"`
}

// creditBlockAccountingSyncSucceededWebhookEventBlockCreditCommitmentJSON contains
// the JSON metadata for the struct
// [CreditBlockAccountingSyncSucceededWebhookEventBlockCreditCommitment]
type creditBlockAccountingSyncSucceededWebhookEventBlockCreditCommitmentJSON struct {
	ID             apijson.Field
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CreditBlockAccountingSyncSucceededWebhookEventBlockCreditCommitment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockAccountingSyncSucceededWebhookEventBlockCreditCommitmentJSON) RawJSON() string {
	return r.raw
}

type CreditBlockAccountingSyncSucceededWebhookEventProperties struct {
	ConnectionType string                                                       `json:"connection_type" api:"required"`
	JSON           creditBlockAccountingSyncSucceededWebhookEventPropertiesJSON `json:"-"`
}

// creditBlockAccountingSyncSucceededWebhookEventPropertiesJSON contains the JSON
// metadata for the struct
// [CreditBlockAccountingSyncSucceededWebhookEventProperties]
type creditBlockAccountingSyncSucceededWebhookEventPropertiesJSON struct {
	ConnectionType apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CreditBlockAccountingSyncSucceededWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockAccountingSyncSucceededWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type CreditBlockAccountingSyncSucceededWebhookEventType string

const (
	CreditBlockAccountingSyncSucceededWebhookEventTypeCreditBlockAccountingSyncSucceeded CreditBlockAccountingSyncSucceededWebhookEventType = "credit_block.accounting_sync_succeeded"
)

func (r CreditBlockAccountingSyncSucceededWebhookEventType) IsKnown() bool {
	switch r {
	case CreditBlockAccountingSyncSucceededWebhookEventTypeCreditBlockAccountingSyncSucceeded:
		return true
	}
	return false
}

// Issued when a credit note accounting sync fails.
type CreditNoteAccountingSyncFailedWebhookEvent struct {
	// The ID of this webhook event.
	ID                   string                                                         `json:"id" api:"required"`
	AccountingSyncRecord CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecord `json:"accounting_sync_record" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The [Credit Note](/invoicing/credit-notes) resource represents a credit that has
	// been applied to a particular invoice.
	CreditNote shared.CreditNote                                    `json:"credit_note" api:"required"`
	Properties CreditNoteAccountingSyncFailedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type CreditNoteAccountingSyncFailedWebhookEventType `json:"type" api:"required"`
	JSON creditNoteAccountingSyncFailedWebhookEventJSON `json:"-"`
}

// creditNoteAccountingSyncFailedWebhookEventJSON contains the JSON metadata for
// the struct [CreditNoteAccountingSyncFailedWebhookEvent]
type creditNoteAccountingSyncFailedWebhookEventJSON struct {
	ID                   apijson.Field
	AccountingSyncRecord apijson.Field
	CreatedAt            apijson.Field
	CreditNote           apijson.Field
	Properties           apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CreditNoteAccountingSyncFailedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteAccountingSyncFailedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CreditNoteAccountingSyncFailedWebhookEvent) implementsUnwrapWebhookEvent() {}

type CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecord struct {
	ID                 string                                                                   `json:"id" api:"required"`
	CustomerID         string                                                                   `json:"customer_id" api:"required"`
	RecordType         CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType `json:"record_type" api:"required"`
	CreditNoteID       string                                                                   `json:"credit_note_id" api:"nullable"`
	ErrorDetails       map[string]interface{}                                                   `json:"error_details" api:"nullable"`
	ProviderCustomerID string                                                                   `json:"provider_customer_id" api:"nullable"`
	Status             string                                                                   `json:"status" api:"nullable"`
	SyncAction         string                                                                   `json:"sync_action" api:"nullable"`
	JSON               creditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordJSON       `json:"-"`
}

// creditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordJSON contains the
// JSON metadata for the struct
// [CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecord]
type creditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordJSON struct {
	ID                 apijson.Field
	CustomerID         apijson.Field
	RecordType         apijson.Field
	CreditNoteID       apijson.Field
	ErrorDetails       apijson.Field
	ProviderCustomerID apijson.Field
	Status             apijson.Field
	SyncAction         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordJSON) RawJSON() string {
	return r.raw
}

type CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType string

const (
	CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomer                   CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "customer"
	CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeInvoice                    CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "invoice"
	CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeTransaction                CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "transaction"
	CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "customer_balance_transaction"
	CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCreditNote                 CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "credit_note"
	CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSubscription               CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "subscription"
	CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSalesOrder                 CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "sales_order"
	CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeBlock                      CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "block"
)

func (r CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType) IsKnown() bool {
	switch r {
	case CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomer, CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeInvoice, CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeTransaction, CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction, CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCreditNote, CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSubscription, CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSalesOrder, CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeBlock:
		return true
	}
	return false
}

type CreditNoteAccountingSyncFailedWebhookEventProperties struct {
	ConnectionType string                                                   `json:"connection_type" api:"required"`
	FailureReason  string                                                   `json:"failure_reason" api:"required"`
	JSON           creditNoteAccountingSyncFailedWebhookEventPropertiesJSON `json:"-"`
}

// creditNoteAccountingSyncFailedWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [CreditNoteAccountingSyncFailedWebhookEventProperties]
type creditNoteAccountingSyncFailedWebhookEventPropertiesJSON struct {
	ConnectionType apijson.Field
	FailureReason  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CreditNoteAccountingSyncFailedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteAccountingSyncFailedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type CreditNoteAccountingSyncFailedWebhookEventType string

const (
	CreditNoteAccountingSyncFailedWebhookEventTypeCreditNoteAccountingSyncFailed CreditNoteAccountingSyncFailedWebhookEventType = "credit_note.accounting_sync_failed"
)

func (r CreditNoteAccountingSyncFailedWebhookEventType) IsKnown() bool {
	switch r {
	case CreditNoteAccountingSyncFailedWebhookEventTypeCreditNoteAccountingSyncFailed:
		return true
	}
	return false
}

// Issued when a credit note accounting sync succeeds.
type CreditNoteAccountingSyncSucceededWebhookEvent struct {
	// The ID of this webhook event.
	ID                   string                                                            `json:"id" api:"required"`
	AccountingSyncRecord CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecord `json:"accounting_sync_record" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The [Credit Note](/invoicing/credit-notes) resource represents a credit that has
	// been applied to a particular invoice.
	CreditNote shared.CreditNote                                       `json:"credit_note" api:"required"`
	Properties CreditNoteAccountingSyncSucceededWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type CreditNoteAccountingSyncSucceededWebhookEventType `json:"type" api:"required"`
	JSON creditNoteAccountingSyncSucceededWebhookEventJSON `json:"-"`
}

// creditNoteAccountingSyncSucceededWebhookEventJSON contains the JSON metadata for
// the struct [CreditNoteAccountingSyncSucceededWebhookEvent]
type creditNoteAccountingSyncSucceededWebhookEventJSON struct {
	ID                   apijson.Field
	AccountingSyncRecord apijson.Field
	CreatedAt            apijson.Field
	CreditNote           apijson.Field
	Properties           apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CreditNoteAccountingSyncSucceededWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteAccountingSyncSucceededWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CreditNoteAccountingSyncSucceededWebhookEvent) implementsUnwrapWebhookEvent() {}

type CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecord struct {
	ID                 string                                                                      `json:"id" api:"required"`
	CustomerID         string                                                                      `json:"customer_id" api:"required"`
	RecordType         CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType `json:"record_type" api:"required"`
	CreditNoteID       string                                                                      `json:"credit_note_id" api:"nullable"`
	ErrorDetails       map[string]interface{}                                                      `json:"error_details" api:"nullable"`
	ProviderCustomerID string                                                                      `json:"provider_customer_id" api:"nullable"`
	Status             string                                                                      `json:"status" api:"nullable"`
	SyncAction         string                                                                      `json:"sync_action" api:"nullable"`
	JSON               creditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON       `json:"-"`
}

// creditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON contains
// the JSON metadata for the struct
// [CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecord]
type creditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON struct {
	ID                 apijson.Field
	CustomerID         apijson.Field
	RecordType         apijson.Field
	CreditNoteID       apijson.Field
	ErrorDetails       apijson.Field
	ProviderCustomerID apijson.Field
	Status             apijson.Field
	SyncAction         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON) RawJSON() string {
	return r.raw
}

type CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType string

const (
	CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomer                   CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "customer"
	CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeInvoice                    CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "invoice"
	CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeTransaction                CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "transaction"
	CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "customer_balance_transaction"
	CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCreditNote                 CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "credit_note"
	CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSubscription               CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "subscription"
	CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSalesOrder                 CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "sales_order"
	CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeBlock                      CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "block"
)

func (r CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType) IsKnown() bool {
	switch r {
	case CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomer, CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeInvoice, CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeTransaction, CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction, CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCreditNote, CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSubscription, CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSalesOrder, CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeBlock:
		return true
	}
	return false
}

type CreditNoteAccountingSyncSucceededWebhookEventProperties struct {
	ConnectionType string                                                      `json:"connection_type" api:"required"`
	JSON           creditNoteAccountingSyncSucceededWebhookEventPropertiesJSON `json:"-"`
}

// creditNoteAccountingSyncSucceededWebhookEventPropertiesJSON contains the JSON
// metadata for the struct
// [CreditNoteAccountingSyncSucceededWebhookEventProperties]
type creditNoteAccountingSyncSucceededWebhookEventPropertiesJSON struct {
	ConnectionType apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CreditNoteAccountingSyncSucceededWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteAccountingSyncSucceededWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type CreditNoteAccountingSyncSucceededWebhookEventType string

const (
	CreditNoteAccountingSyncSucceededWebhookEventTypeCreditNoteAccountingSyncSucceeded CreditNoteAccountingSyncSucceededWebhookEventType = "credit_note.accounting_sync_succeeded"
)

func (r CreditNoteAccountingSyncSucceededWebhookEventType) IsKnown() bool {
	switch r {
	case CreditNoteAccountingSyncSucceededWebhookEventTypeCreditNoteAccountingSyncSucceeded:
		return true
	}
	return false
}

// Issued when a credit note is created.
type CreditNoteIssuedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The [Credit Note](/invoicing/credit-notes) resource represents a credit that has
	// been applied to a particular invoice.
	CreditNote shared.CreditNote `json:"credit_note" api:"required"`
	Properties interface{}       `json:"properties" api:"required"`
	// The event this payload describes.
	Type CreditNoteIssuedWebhookEventType `json:"type" api:"required"`
	JSON creditNoteIssuedWebhookEventJSON `json:"-"`
}

// creditNoteIssuedWebhookEventJSON contains the JSON metadata for the struct
// [CreditNoteIssuedWebhookEvent]
type creditNoteIssuedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	CreditNote  apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditNoteIssuedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteIssuedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CreditNoteIssuedWebhookEvent) implementsUnwrapWebhookEvent() {}

// The event this payload describes.
type CreditNoteIssuedWebhookEventType string

const (
	CreditNoteIssuedWebhookEventTypeCreditNoteIssued CreditNoteIssuedWebhookEventType = "credit_note.issued"
)

func (r CreditNoteIssuedWebhookEventType) IsKnown() bool {
	switch r {
	case CreditNoteIssuedWebhookEventTypeCreditNoteIssued:
		return true
	}
	return false
}

// Issued when a credit note is marked as void.
type CreditNoteMarkedAsVoidWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The [Credit Note](/invoicing/credit-notes) resource represents a credit that has
	// been applied to a particular invoice.
	CreditNote shared.CreditNote `json:"credit_note" api:"required"`
	Properties interface{}       `json:"properties" api:"required"`
	// The event this payload describes.
	Type CreditNoteMarkedAsVoidWebhookEventType `json:"type" api:"required"`
	JSON creditNoteMarkedAsVoidWebhookEventJSON `json:"-"`
}

// creditNoteMarkedAsVoidWebhookEventJSON contains the JSON metadata for the struct
// [CreditNoteMarkedAsVoidWebhookEvent]
type creditNoteMarkedAsVoidWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	CreditNote  apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditNoteMarkedAsVoidWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteMarkedAsVoidWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CreditNoteMarkedAsVoidWebhookEvent) implementsUnwrapWebhookEvent() {}

// The event this payload describes.
type CreditNoteMarkedAsVoidWebhookEventType string

const (
	CreditNoteMarkedAsVoidWebhookEventTypeCreditNoteMarkedAsVoid CreditNoteMarkedAsVoidWebhookEventType = "credit_note.marked_as_void"
)

func (r CreditNoteMarkedAsVoidWebhookEventType) IsKnown() bool {
	switch r {
	case CreditNoteMarkedAsVoidWebhookEventTypeCreditNoteMarkedAsVoid:
		return true
	}
	return false
}

// Issued when a customer accounting sync fails.
type CustomerAccountingSyncFailedWebhookEvent struct {
	// The ID of this webhook event.
	ID                   string                                                       `json:"id" api:"required"`
	AccountingSyncRecord CustomerAccountingSyncFailedWebhookEventAccountingSyncRecord `json:"accounting_sync_record" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
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
	Customer   Customer                                           `json:"customer" api:"required"`
	Properties CustomerAccountingSyncFailedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type CustomerAccountingSyncFailedWebhookEventType `json:"type" api:"required"`
	JSON customerAccountingSyncFailedWebhookEventJSON `json:"-"`
}

// customerAccountingSyncFailedWebhookEventJSON contains the JSON metadata for the
// struct [CustomerAccountingSyncFailedWebhookEvent]
type customerAccountingSyncFailedWebhookEventJSON struct {
	ID                   apijson.Field
	AccountingSyncRecord apijson.Field
	CreatedAt            apijson.Field
	Customer             apijson.Field
	Properties           apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CustomerAccountingSyncFailedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAccountingSyncFailedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CustomerAccountingSyncFailedWebhookEvent) implementsUnwrapWebhookEvent() {}

type CustomerAccountingSyncFailedWebhookEventAccountingSyncRecord struct {
	ID                 string                                                                 `json:"id" api:"required"`
	CustomerID         string                                                                 `json:"customer_id" api:"required"`
	RecordType         CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType `json:"record_type" api:"required"`
	ErrorDetails       map[string]interface{}                                                 `json:"error_details" api:"nullable"`
	ProviderCustomerID string                                                                 `json:"provider_customer_id" api:"nullable"`
	Status             string                                                                 `json:"status" api:"nullable"`
	SyncAction         string                                                                 `json:"sync_action" api:"nullable"`
	JSON               customerAccountingSyncFailedWebhookEventAccountingSyncRecordJSON       `json:"-"`
}

// customerAccountingSyncFailedWebhookEventAccountingSyncRecordJSON contains the
// JSON metadata for the struct
// [CustomerAccountingSyncFailedWebhookEventAccountingSyncRecord]
type customerAccountingSyncFailedWebhookEventAccountingSyncRecordJSON struct {
	ID                 apijson.Field
	CustomerID         apijson.Field
	RecordType         apijson.Field
	ErrorDetails       apijson.Field
	ProviderCustomerID apijson.Field
	Status             apijson.Field
	SyncAction         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerAccountingSyncFailedWebhookEventAccountingSyncRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAccountingSyncFailedWebhookEventAccountingSyncRecordJSON) RawJSON() string {
	return r.raw
}

type CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType string

const (
	CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomer                   CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "customer"
	CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeInvoice                    CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "invoice"
	CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeTransaction                CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "transaction"
	CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "customer_balance_transaction"
	CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCreditNote                 CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "credit_note"
	CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSubscription               CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "subscription"
	CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSalesOrder                 CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "sales_order"
	CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeBlock                      CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "block"
)

func (r CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType) IsKnown() bool {
	switch r {
	case CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomer, CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeInvoice, CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeTransaction, CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction, CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCreditNote, CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSubscription, CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSalesOrder, CustomerAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeBlock:
		return true
	}
	return false
}

type CustomerAccountingSyncFailedWebhookEventProperties struct {
	ConnectionType string                                                 `json:"connection_type" api:"required"`
	FailureReason  string                                                 `json:"failure_reason" api:"required"`
	JSON           customerAccountingSyncFailedWebhookEventPropertiesJSON `json:"-"`
}

// customerAccountingSyncFailedWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [CustomerAccountingSyncFailedWebhookEventProperties]
type customerAccountingSyncFailedWebhookEventPropertiesJSON struct {
	ConnectionType apijson.Field
	FailureReason  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerAccountingSyncFailedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAccountingSyncFailedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type CustomerAccountingSyncFailedWebhookEventType string

const (
	CustomerAccountingSyncFailedWebhookEventTypeCustomerAccountingSyncFailed CustomerAccountingSyncFailedWebhookEventType = "customer.accounting_sync_failed"
)

func (r CustomerAccountingSyncFailedWebhookEventType) IsKnown() bool {
	switch r {
	case CustomerAccountingSyncFailedWebhookEventTypeCustomerAccountingSyncFailed:
		return true
	}
	return false
}

// Issued when a customer accounting sync succeeds.
type CustomerAccountingSyncSucceededWebhookEvent struct {
	// The ID of this webhook event.
	ID                   string                                                          `json:"id" api:"required"`
	AccountingSyncRecord CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecord `json:"accounting_sync_record" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
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
	Customer   Customer                                              `json:"customer" api:"required"`
	Properties CustomerAccountingSyncSucceededWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type CustomerAccountingSyncSucceededWebhookEventType `json:"type" api:"required"`
	JSON customerAccountingSyncSucceededWebhookEventJSON `json:"-"`
}

// customerAccountingSyncSucceededWebhookEventJSON contains the JSON metadata for
// the struct [CustomerAccountingSyncSucceededWebhookEvent]
type customerAccountingSyncSucceededWebhookEventJSON struct {
	ID                   apijson.Field
	AccountingSyncRecord apijson.Field
	CreatedAt            apijson.Field
	Customer             apijson.Field
	Properties           apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CustomerAccountingSyncSucceededWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAccountingSyncSucceededWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CustomerAccountingSyncSucceededWebhookEvent) implementsUnwrapWebhookEvent() {}

type CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecord struct {
	ID                 string                                                                    `json:"id" api:"required"`
	CustomerID         string                                                                    `json:"customer_id" api:"required"`
	RecordType         CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType `json:"record_type" api:"required"`
	ErrorDetails       map[string]interface{}                                                    `json:"error_details" api:"nullable"`
	ProviderCustomerID string                                                                    `json:"provider_customer_id" api:"nullable"`
	Status             string                                                                    `json:"status" api:"nullable"`
	SyncAction         string                                                                    `json:"sync_action" api:"nullable"`
	JSON               customerAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON       `json:"-"`
}

// customerAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON contains the
// JSON metadata for the struct
// [CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecord]
type customerAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON struct {
	ID                 apijson.Field
	CustomerID         apijson.Field
	RecordType         apijson.Field
	ErrorDetails       apijson.Field
	ProviderCustomerID apijson.Field
	Status             apijson.Field
	SyncAction         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON) RawJSON() string {
	return r.raw
}

type CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType string

const (
	CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomer                   CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "customer"
	CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeInvoice                    CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "invoice"
	CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeTransaction                CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "transaction"
	CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "customer_balance_transaction"
	CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCreditNote                 CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "credit_note"
	CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSubscription               CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "subscription"
	CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSalesOrder                 CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "sales_order"
	CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeBlock                      CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "block"
)

func (r CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType) IsKnown() bool {
	switch r {
	case CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomer, CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeInvoice, CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeTransaction, CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction, CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCreditNote, CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSubscription, CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSalesOrder, CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeBlock:
		return true
	}
	return false
}

type CustomerAccountingSyncSucceededWebhookEventProperties struct {
	ConnectionType string                                                    `json:"connection_type" api:"required"`
	JSON           customerAccountingSyncSucceededWebhookEventPropertiesJSON `json:"-"`
}

// customerAccountingSyncSucceededWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [CustomerAccountingSyncSucceededWebhookEventProperties]
type customerAccountingSyncSucceededWebhookEventPropertiesJSON struct {
	ConnectionType apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerAccountingSyncSucceededWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerAccountingSyncSucceededWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type CustomerAccountingSyncSucceededWebhookEventType string

const (
	CustomerAccountingSyncSucceededWebhookEventTypeCustomerAccountingSyncSucceeded CustomerAccountingSyncSucceededWebhookEventType = "customer.accounting_sync_succeeded"
)

func (r CustomerAccountingSyncSucceededWebhookEventType) IsKnown() bool {
	switch r {
	case CustomerAccountingSyncSucceededWebhookEventTypeCustomerAccountingSyncSucceeded:
		return true
	}
	return false
}

// Issued when a customer balance transaction is created.
type CustomerBalanceTransactionCreatedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
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
	Customer   Customer                                                `json:"customer" api:"required"`
	Properties CustomerBalanceTransactionCreatedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type CustomerBalanceTransactionCreatedWebhookEventType `json:"type" api:"required"`
	JSON customerBalanceTransactionCreatedWebhookEventJSON `json:"-"`
}

// customerBalanceTransactionCreatedWebhookEventJSON contains the JSON metadata for
// the struct [CustomerBalanceTransactionCreatedWebhookEvent]
type customerBalanceTransactionCreatedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Customer    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerBalanceTransactionCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerBalanceTransactionCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CustomerBalanceTransactionCreatedWebhookEvent) implementsUnwrapWebhookEvent() {}

type CustomerBalanceTransactionCreatedWebhookEventProperties struct {
	BalanceTransaction CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransaction `json:"balance_transaction" api:"required"`
	JSON               customerBalanceTransactionCreatedWebhookEventPropertiesJSON               `json:"-"`
}

// customerBalanceTransactionCreatedWebhookEventPropertiesJSON contains the JSON
// metadata for the struct
// [CustomerBalanceTransactionCreatedWebhookEventProperties]
type customerBalanceTransactionCreatedWebhookEventPropertiesJSON struct {
	BalanceTransaction apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerBalanceTransactionCreatedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerBalanceTransactionCreatedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

type CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransaction struct {
	// A unique id for this transaction.
	ID     string                                                                          `json:"id" api:"required"`
	Action CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionAction `json:"action" api:"required"`
	// The value of the amount changed in the transaction.
	Amount string `json:"amount" api:"required"`
	// The creation time of this transaction.
	CreatedAt  time.Time             `json:"created_at" api:"required" format:"date-time"`
	CreditNote shared.CreditNoteTiny `json:"credit_note" api:"required,nullable"`
	// An optional description provided for manual customer balance adjustments.
	Description string `json:"description" api:"required,nullable"`
	// The new value of the customer's balance prior to the transaction, in the
	// customer's currency.
	EndingBalance string             `json:"ending_balance" api:"required"`
	Invoice       shared.InvoiceTiny `json:"invoice" api:"required,nullable"`
	// The original value of the customer's balance prior to the transaction, in the
	// customer's currency.
	StartingBalance string                                                                        `json:"starting_balance" api:"required"`
	Type            CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionType `json:"type" api:"required"`
	JSON            customerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionJSON `json:"-"`
}

// customerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionJSON
// contains the JSON metadata for the struct
// [CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransaction]
type customerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionJSON struct {
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

func (r *CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionJSON) RawJSON() string {
	return r.raw
}

type CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionAction string

const (
	CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionAppliedToInvoice      CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionAction = "applied_to_invoice"
	CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionManualAdjustment      CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionAction = "manual_adjustment"
	CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionProratedRefund        CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionAction = "prorated_refund"
	CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionRevertProratedRefund  CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionAction = "revert_prorated_refund"
	CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionReturnFromVoiding     CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionAction = "return_from_voiding"
	CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionCreditNoteApplied     CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionAction = "credit_note_applied"
	CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionCreditNoteVoided      CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionAction = "credit_note_voided"
	CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionOverpaymentRefund     CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionAction = "overpayment_refund"
	CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionExternalPayment       CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionAction = "external_payment"
	CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionSmallInvoiceCarryover CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionAction = "small_invoice_carryover"
	CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionPrepaidCommitCancel   CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionAction = "prepaid_commit_cancel"
)

func (r CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionAction) IsKnown() bool {
	switch r {
	case CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionAppliedToInvoice, CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionManualAdjustment, CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionProratedRefund, CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionRevertProratedRefund, CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionReturnFromVoiding, CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionCreditNoteApplied, CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionCreditNoteVoided, CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionOverpaymentRefund, CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionExternalPayment, CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionSmallInvoiceCarryover, CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionActionPrepaidCommitCancel:
		return true
	}
	return false
}

type CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionType string

const (
	CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionTypeIncrement CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionType = "increment"
	CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionTypeDecrement CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionType = "decrement"
)

func (r CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionType) IsKnown() bool {
	switch r {
	case CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionTypeIncrement, CustomerBalanceTransactionCreatedWebhookEventPropertiesBalanceTransactionTypeDecrement:
		return true
	}
	return false
}

// The event this payload describes.
type CustomerBalanceTransactionCreatedWebhookEventType string

const (
	CustomerBalanceTransactionCreatedWebhookEventTypeCustomerBalanceTransactionCreated CustomerBalanceTransactionCreatedWebhookEventType = "customer.balance_transaction_created"
)

func (r CustomerBalanceTransactionCreatedWebhookEventType) IsKnown() bool {
	switch r {
	case CustomerBalanceTransactionCreatedWebhookEventTypeCustomerBalanceTransactionCreated:
		return true
	}
	return false
}

// Issued when a customer resource is created.
type CustomerCreatedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
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
	Customer   Customer    `json:"customer" api:"required"`
	Properties interface{} `json:"properties" api:"required"`
	// The event this payload describes.
	Type CustomerCreatedWebhookEventType `json:"type" api:"required"`
	JSON customerCreatedWebhookEventJSON `json:"-"`
}

// customerCreatedWebhookEventJSON contains the JSON metadata for the struct
// [CustomerCreatedWebhookEvent]
type customerCreatedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Customer    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreatedWebhookEvent) implementsUnwrapWebhookEvent() {}

// The event this payload describes.
type CustomerCreatedWebhookEventType string

const (
	CustomerCreatedWebhookEventTypeCustomerCreated CustomerCreatedWebhookEventType = "customer.created"
)

func (r CustomerCreatedWebhookEventType) IsKnown() bool {
	switch r {
	case CustomerCreatedWebhookEventTypeCustomerCreated:
		return true
	}
	return false
}

// Issued when a customer's prepaid credits balance is depleted.
type CustomerCreditBalanceDepletedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
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
	Customer   Customer                                            `json:"customer" api:"required"`
	Properties CustomerCreditBalanceDepletedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type CustomerCreditBalanceDepletedWebhookEventType `json:"type" api:"required"`
	JSON customerCreditBalanceDepletedWebhookEventJSON `json:"-"`
}

// customerCreditBalanceDepletedWebhookEventJSON contains the JSON metadata for the
// struct [CustomerCreditBalanceDepletedWebhookEvent]
type customerCreditBalanceDepletedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Customer    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceDepletedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDepletedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditBalanceDepletedWebhookEvent) implementsUnwrapWebhookEvent() {}

type CustomerCreditBalanceDepletedWebhookEventProperties struct {
	// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
	// usage, or credit balance and trigger webhooks when a threshold is exceeded.
	//
	// Alerts created through the API can be scoped to either customers or
	// subscriptions.
	AlertConfiguration CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfiguration `json:"alert_configuration" api:"required"`
	// A currency or custom credit unit, as embedded in webhook payloads.
	PricingUnit CustomerCreditBalanceDepletedWebhookEventPropertiesPricingUnit `json:"pricing_unit" api:"required,nullable"`
	JSON        customerCreditBalanceDepletedWebhookEventPropertiesJSON        `json:"-"`
}

// customerCreditBalanceDepletedWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [CustomerCreditBalanceDepletedWebhookEventProperties]
type customerCreditBalanceDepletedWebhookEventPropertiesJSON struct {
	AlertConfiguration apijson.Field
	PricingUnit        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditBalanceDepletedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDepletedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
// usage, or credit balance and trigger webhooks when a threshold is exceeded.
//
// Alerts created through the API can be scoped to either customers or
// subscriptions.
type CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfiguration struct {
	// Also referred to as alert_id in this documentation.
	ID string `json:"id" api:"required"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The name of the currency the credit balance or invoice cost is denominated in.
	Currency string `json:"currency" api:"required,nullable"`
	// The customer the alert applies to.
	Customer shared.CustomerMinified `json:"customer" api:"required,nullable"`
	// Whether the alert is enabled or disabled.
	Enabled bool `json:"enabled" api:"required"`
	// The metric the alert applies to.
	Metric CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationMetric `json:"metric" api:"required,nullable"`
	// The plan the alert applies to.
	Plan CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPlan `json:"plan" api:"required,nullable"`
	// The subscription the alert applies to.
	Subscription shared.SubscriptionMinified `json:"subscription" api:"required,nullable"`
	// The thresholds that define the conditions under which the alert will be
	// triggered.
	Thresholds []CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThreshold `json:"thresholds" api:"required,nullable"`
	// The type of alert. This must be a valid alert type.
	Type CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationType `json:"type" api:"required"`
	// The current status of the alert. This field is only present for credit balance
	// alerts.
	BalanceAlertStatus []CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationBalanceAlertStatus `json:"balance_alert_status" api:"nullable"`
	// The property keys to group cost alerts by. Only present for cost alerts with
	// grouping enabled.
	GroupingKeys []string `json:"grouping_keys" api:"nullable"`
	// Minified license type for alert serialization.
	LicenseType CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationLicenseType `json:"license_type" api:"nullable"`
	// Filters scoping which prices are included in spend and grouped cost alert
	// evaluation. Alerts use the price_id, item_id, and price_type fields only; the
	// alert's pricing unit is reported by currency.
	PriceFilters []CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFilter `json:"price_filters" api:"nullable"`
	// Per-group threshold overrides. Each override maps a specific combination of
	// grouping_keys values to a replacement threshold list. Only present for grouped
	// cost alerts that have at least one override.
	ThresholdOverrides []CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdOverride `json:"threshold_overrides" api:"nullable"`
	JSON               customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationJSON                `json:"-"`
}

// customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfiguration]
type customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Customer           apijson.Field
	Enabled            apijson.Field
	Metric             apijson.Field
	Plan               apijson.Field
	Subscription       apijson.Field
	Thresholds         apijson.Field
	Type               apijson.Field
	BalanceAlertStatus apijson.Field
	GroupingKeys       apijson.Field
	LicenseType        apijson.Field
	PriceFilters       apijson.Field
	ThresholdOverrides apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationJSON) RawJSON() string {
	return r.raw
}

// The metric the alert applies to.
type CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationMetric struct {
	ID   string                                                                          `json:"id" api:"required"`
	JSON customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationMetricJSON `json:"-"`
}

// customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationMetricJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationMetric]
type customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationMetricJSON) RawJSON() string {
	return r.raw
}

// The plan the alert applies to.
type CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPlan struct {
	ID string `json:"id" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                                                                        `json:"external_plan_id" api:"required,nullable"`
	Name           string                                                                        `json:"name" api:"required,nullable"`
	PlanVersion    string                                                                        `json:"plan_version" api:"required"`
	JSON           customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPlanJSON `json:"-"`
}

// customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPlanJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPlan]
type customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	PlanVersion    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPlanJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                             `json:"value" api:"required"`
	JSON  customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdJSON `json:"-"`
}

// customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThreshold]
type customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdJSON) RawJSON() string {
	return r.raw
}

// The type of alert. This must be a valid alert type.
type CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationType string

const (
	CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted          CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationType = "credit_balance_depleted"
	CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped           CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationType = "credit_balance_dropped"
	CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered         CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationType = "credit_balance_recovered"
	CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationTypeUsageExceeded                  CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationType = "usage_exceeded"
	CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationTypeCostExceeded                   CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationType = "cost_exceeded"
	CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationTypeSpendExceeded                  CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationType = "spend_exceeded"
	CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationType = "license_balance_threshold_reached"
)

func (r CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationType) IsKnown() bool {
	switch r {
	case CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted, CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped, CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered, CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationTypeUsageExceeded, CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationTypeCostExceeded, CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationTypeSpendExceeded, CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached:
		return true
	}
	return false
}

// Alert status is used to determine if an alert is currently in-alert or not.
type CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationBalanceAlertStatus struct {
	// Whether the alert is currently in-alert or not.
	InAlert bool `json:"in_alert" api:"required"`
	// The value of the threshold that defines the alert status.
	ThresholdValue string                                                                                      `json:"threshold_value" api:"required"`
	JSON           customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON `json:"-"`
}

// customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationBalanceAlertStatus]
type customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON struct {
	InAlert        apijson.Field
	ThresholdValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationBalanceAlertStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON) RawJSON() string {
	return r.raw
}

// Minified license type for alert serialization.
type CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationLicenseType struct {
	ID   string                                                                               `json:"id" api:"required"`
	JSON customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationLicenseTypeJSON `json:"-"`
}

// customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationLicenseTypeJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationLicenseType]
type customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationLicenseTypeJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationLicenseType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationLicenseTypeJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFilter struct {
	// The property of the price to filter on.
	Field CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                                             `json:"values" api:"required"`
	JSON   customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFilterJSON `json:"-"`
}

// customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFilterJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFilter]
type customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersField string

const (
	CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID       CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_id"
	CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID        CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersField = "item_id"
	CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType     CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_type"
	CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency      CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersField = "currency"
	CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersField = "pricing_unit_id"
)

func (r CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersField) IsKnown() bool {
	switch r {
	case CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID, CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID, CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType, CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency, CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersOperator string

const (
	CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "includes"
	CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "excludes"
)

func (r CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersOperator) IsKnown() bool {
	switch r {
	case CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes, CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes:
		return true
	}
	return false
}

// A per-group threshold override on a grouped cost alert.
//
// An empty `thresholds` list means the group is silenced (never fires). A
// non-empty list fully replaces the default thresholds for that group.
type CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdOverride struct {
	// The values identifying this group, ordered to match group_keys when set and the
	// alert's grouping_keys otherwise.
	GroupValues []string `json:"group_values" api:"required"`
	// The thresholds applied to this group. An empty list means the group is silenced.
	Thresholds []CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold `json:"thresholds" api:"required"`
	// The subset of the alert's grouping_keys this override binds. Null when the
	// override targets one exact group across every grouping key.
	GroupKeys []string                                                                                   `json:"group_keys" api:"nullable"`
	JSON      customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON `json:"-"`
}

// customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdOverride]
type customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON struct {
	GroupValues apijson.Field
	Thresholds  apijson.Field
	GroupKeys   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                                               `json:"value" api:"required"`
	JSON  customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON `json:"-"`
}

// customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold]
type customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDepletedWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON) RawJSON() string {
	return r.raw
}

// A currency or custom credit unit, as embedded in webhook payloads.
type CustomerCreditBalanceDepletedWebhookEventPropertiesPricingUnit struct {
	ID          string                                                             `json:"id" api:"required"`
	DisplayName string                                                             `json:"display_name" api:"required,nullable"`
	Name        string                                                             `json:"name" api:"required"`
	Symbol      string                                                             `json:"symbol" api:"required,nullable"`
	JSON        customerCreditBalanceDepletedWebhookEventPropertiesPricingUnitJSON `json:"-"`
}

// customerCreditBalanceDepletedWebhookEventPropertiesPricingUnitJSON contains the
// JSON metadata for the struct
// [CustomerCreditBalanceDepletedWebhookEventPropertiesPricingUnit]
type customerCreditBalanceDepletedWebhookEventPropertiesPricingUnitJSON struct {
	ID          apijson.Field
	DisplayName apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceDepletedWebhookEventPropertiesPricingUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDepletedWebhookEventPropertiesPricingUnitJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type CustomerCreditBalanceDepletedWebhookEventType string

const (
	CustomerCreditBalanceDepletedWebhookEventTypeCustomerCreditBalanceDepleted CustomerCreditBalanceDepletedWebhookEventType = "customer.credit_balance_depleted"
)

func (r CustomerCreditBalanceDepletedWebhookEventType) IsKnown() bool {
	switch r {
	case CustomerCreditBalanceDepletedWebhookEventTypeCustomerCreditBalanceDepleted:
		return true
	}
	return false
}

// Issued when a customer's prepaid credits balance is depleted to a configured
// threshold.
type CustomerCreditBalanceDroppedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
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
	Customer   Customer                                           `json:"customer" api:"required"`
	Properties CustomerCreditBalanceDroppedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type CustomerCreditBalanceDroppedWebhookEventType `json:"type" api:"required"`
	JSON customerCreditBalanceDroppedWebhookEventJSON `json:"-"`
}

// customerCreditBalanceDroppedWebhookEventJSON contains the JSON metadata for the
// struct [CustomerCreditBalanceDroppedWebhookEvent]
type customerCreditBalanceDroppedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Customer    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceDroppedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDroppedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditBalanceDroppedWebhookEvent) implementsUnwrapWebhookEvent() {}

type CustomerCreditBalanceDroppedWebhookEventProperties struct {
	// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
	// usage, or credit balance and trigger webhooks when a threshold is exceeded.
	//
	// Alerts created through the API can be scoped to either customers or
	// subscriptions.
	AlertConfiguration CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfiguration `json:"alert_configuration" api:"required"`
	BalanceThreshold   string                                                               `json:"balance_threshold" api:"required"`
	// A currency or custom credit unit, as embedded in webhook payloads.
	PricingUnit CustomerCreditBalanceDroppedWebhookEventPropertiesPricingUnit `json:"pricing_unit" api:"required,nullable"`
	JSON        customerCreditBalanceDroppedWebhookEventPropertiesJSON        `json:"-"`
}

// customerCreditBalanceDroppedWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [CustomerCreditBalanceDroppedWebhookEventProperties]
type customerCreditBalanceDroppedWebhookEventPropertiesJSON struct {
	AlertConfiguration apijson.Field
	BalanceThreshold   apijson.Field
	PricingUnit        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditBalanceDroppedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDroppedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
// usage, or credit balance and trigger webhooks when a threshold is exceeded.
//
// Alerts created through the API can be scoped to either customers or
// subscriptions.
type CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfiguration struct {
	// Also referred to as alert_id in this documentation.
	ID string `json:"id" api:"required"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The name of the currency the credit balance or invoice cost is denominated in.
	Currency string `json:"currency" api:"required,nullable"`
	// The customer the alert applies to.
	Customer shared.CustomerMinified `json:"customer" api:"required,nullable"`
	// Whether the alert is enabled or disabled.
	Enabled bool `json:"enabled" api:"required"`
	// The metric the alert applies to.
	Metric CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationMetric `json:"metric" api:"required,nullable"`
	// The plan the alert applies to.
	Plan CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPlan `json:"plan" api:"required,nullable"`
	// The subscription the alert applies to.
	Subscription shared.SubscriptionMinified `json:"subscription" api:"required,nullable"`
	// The thresholds that define the conditions under which the alert will be
	// triggered.
	Thresholds []CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThreshold `json:"thresholds" api:"required,nullable"`
	// The type of alert. This must be a valid alert type.
	Type CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationType `json:"type" api:"required"`
	// The current status of the alert. This field is only present for credit balance
	// alerts.
	BalanceAlertStatus []CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationBalanceAlertStatus `json:"balance_alert_status" api:"nullable"`
	// The property keys to group cost alerts by. Only present for cost alerts with
	// grouping enabled.
	GroupingKeys []string `json:"grouping_keys" api:"nullable"`
	// Minified license type for alert serialization.
	LicenseType CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationLicenseType `json:"license_type" api:"nullable"`
	// Filters scoping which prices are included in spend and grouped cost alert
	// evaluation. Alerts use the price_id, item_id, and price_type fields only; the
	// alert's pricing unit is reported by currency.
	PriceFilters []CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFilter `json:"price_filters" api:"nullable"`
	// Per-group threshold overrides. Each override maps a specific combination of
	// grouping_keys values to a replacement threshold list. Only present for grouped
	// cost alerts that have at least one override.
	ThresholdOverrides []CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdOverride `json:"threshold_overrides" api:"nullable"`
	JSON               customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationJSON                `json:"-"`
}

// customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfiguration]
type customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Customer           apijson.Field
	Enabled            apijson.Field
	Metric             apijson.Field
	Plan               apijson.Field
	Subscription       apijson.Field
	Thresholds         apijson.Field
	Type               apijson.Field
	BalanceAlertStatus apijson.Field
	GroupingKeys       apijson.Field
	LicenseType        apijson.Field
	PriceFilters       apijson.Field
	ThresholdOverrides apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationJSON) RawJSON() string {
	return r.raw
}

// The metric the alert applies to.
type CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationMetric struct {
	ID   string                                                                         `json:"id" api:"required"`
	JSON customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationMetricJSON `json:"-"`
}

// customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationMetricJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationMetric]
type customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationMetricJSON) RawJSON() string {
	return r.raw
}

// The plan the alert applies to.
type CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPlan struct {
	ID string `json:"id" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                                                                       `json:"external_plan_id" api:"required,nullable"`
	Name           string                                                                       `json:"name" api:"required,nullable"`
	PlanVersion    string                                                                       `json:"plan_version" api:"required"`
	JSON           customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPlanJSON `json:"-"`
}

// customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPlanJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPlan]
type customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	PlanVersion    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPlanJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                            `json:"value" api:"required"`
	JSON  customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdJSON `json:"-"`
}

// customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThreshold]
type customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdJSON) RawJSON() string {
	return r.raw
}

// The type of alert. This must be a valid alert type.
type CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationType string

const (
	CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted          CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationType = "credit_balance_depleted"
	CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped           CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationType = "credit_balance_dropped"
	CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered         CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationType = "credit_balance_recovered"
	CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationTypeUsageExceeded                  CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationType = "usage_exceeded"
	CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationTypeCostExceeded                   CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationType = "cost_exceeded"
	CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationTypeSpendExceeded                  CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationType = "spend_exceeded"
	CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationType = "license_balance_threshold_reached"
)

func (r CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationType) IsKnown() bool {
	switch r {
	case CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted, CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped, CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered, CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationTypeUsageExceeded, CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationTypeCostExceeded, CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationTypeSpendExceeded, CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached:
		return true
	}
	return false
}

// Alert status is used to determine if an alert is currently in-alert or not.
type CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationBalanceAlertStatus struct {
	// Whether the alert is currently in-alert or not.
	InAlert bool `json:"in_alert" api:"required"`
	// The value of the threshold that defines the alert status.
	ThresholdValue string                                                                                     `json:"threshold_value" api:"required"`
	JSON           customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON `json:"-"`
}

// customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationBalanceAlertStatus]
type customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON struct {
	InAlert        apijson.Field
	ThresholdValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationBalanceAlertStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON) RawJSON() string {
	return r.raw
}

// Minified license type for alert serialization.
type CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationLicenseType struct {
	ID   string                                                                              `json:"id" api:"required"`
	JSON customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationLicenseTypeJSON `json:"-"`
}

// customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationLicenseTypeJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationLicenseType]
type customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationLicenseTypeJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationLicenseType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationLicenseTypeJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFilter struct {
	// The property of the price to filter on.
	Field CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                                            `json:"values" api:"required"`
	JSON   customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFilterJSON `json:"-"`
}

// customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFilterJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFilter]
type customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersField string

const (
	CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID       CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_id"
	CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID        CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersField = "item_id"
	CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType     CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_type"
	CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency      CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersField = "currency"
	CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersField = "pricing_unit_id"
)

func (r CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersField) IsKnown() bool {
	switch r {
	case CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID, CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID, CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType, CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency, CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersOperator string

const (
	CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "includes"
	CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "excludes"
)

func (r CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersOperator) IsKnown() bool {
	switch r {
	case CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes, CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes:
		return true
	}
	return false
}

// A per-group threshold override on a grouped cost alert.
//
// An empty `thresholds` list means the group is silenced (never fires). A
// non-empty list fully replaces the default thresholds for that group.
type CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdOverride struct {
	// The values identifying this group, ordered to match group_keys when set and the
	// alert's grouping_keys otherwise.
	GroupValues []string `json:"group_values" api:"required"`
	// The thresholds applied to this group. An empty list means the group is silenced.
	Thresholds []CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold `json:"thresholds" api:"required"`
	// The subset of the alert's grouping_keys this override binds. Null when the
	// override targets one exact group across every grouping key.
	GroupKeys []string                                                                                  `json:"group_keys" api:"nullable"`
	JSON      customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON `json:"-"`
}

// customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdOverride]
type customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON struct {
	GroupValues apijson.Field
	Thresholds  apijson.Field
	GroupKeys   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                                              `json:"value" api:"required"`
	JSON  customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON `json:"-"`
}

// customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold]
type customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDroppedWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON) RawJSON() string {
	return r.raw
}

// A currency or custom credit unit, as embedded in webhook payloads.
type CustomerCreditBalanceDroppedWebhookEventPropertiesPricingUnit struct {
	ID          string                                                            `json:"id" api:"required"`
	DisplayName string                                                            `json:"display_name" api:"required,nullable"`
	Name        string                                                            `json:"name" api:"required"`
	Symbol      string                                                            `json:"symbol" api:"required,nullable"`
	JSON        customerCreditBalanceDroppedWebhookEventPropertiesPricingUnitJSON `json:"-"`
}

// customerCreditBalanceDroppedWebhookEventPropertiesPricingUnitJSON contains the
// JSON metadata for the struct
// [CustomerCreditBalanceDroppedWebhookEventPropertiesPricingUnit]
type customerCreditBalanceDroppedWebhookEventPropertiesPricingUnitJSON struct {
	ID          apijson.Field
	DisplayName apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceDroppedWebhookEventPropertiesPricingUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceDroppedWebhookEventPropertiesPricingUnitJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type CustomerCreditBalanceDroppedWebhookEventType string

const (
	CustomerCreditBalanceDroppedWebhookEventTypeCustomerCreditBalanceDropped CustomerCreditBalanceDroppedWebhookEventType = "customer.credit_balance_dropped"
)

func (r CustomerCreditBalanceDroppedWebhookEventType) IsKnown() bool {
	switch r {
	case CustomerCreditBalanceDroppedWebhookEventTypeCustomerCreditBalanceDropped:
		return true
	}
	return false
}

// Issued when a customer's credit balance recovers from depleted.
type CustomerCreditBalanceRecoveredWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
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
	Customer   Customer                                             `json:"customer" api:"required"`
	Properties CustomerCreditBalanceRecoveredWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type CustomerCreditBalanceRecoveredWebhookEventType `json:"type" api:"required"`
	JSON customerCreditBalanceRecoveredWebhookEventJSON `json:"-"`
}

// customerCreditBalanceRecoveredWebhookEventJSON contains the JSON metadata for
// the struct [CustomerCreditBalanceRecoveredWebhookEvent]
type customerCreditBalanceRecoveredWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Customer    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceRecoveredWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceRecoveredWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditBalanceRecoveredWebhookEvent) implementsUnwrapWebhookEvent() {}

type CustomerCreditBalanceRecoveredWebhookEventProperties struct {
	// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
	// usage, or credit balance and trigger webhooks when a threshold is exceeded.
	//
	// Alerts created through the API can be scoped to either customers or
	// subscriptions.
	AlertConfiguration CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfiguration `json:"alert_configuration" api:"required"`
	// A currency or custom credit unit, as embedded in webhook payloads.
	PricingUnit CustomerCreditBalanceRecoveredWebhookEventPropertiesPricingUnit `json:"pricing_unit" api:"required,nullable"`
	JSON        customerCreditBalanceRecoveredWebhookEventPropertiesJSON        `json:"-"`
}

// customerCreditBalanceRecoveredWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [CustomerCreditBalanceRecoveredWebhookEventProperties]
type customerCreditBalanceRecoveredWebhookEventPropertiesJSON struct {
	AlertConfiguration apijson.Field
	PricingUnit        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditBalanceRecoveredWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceRecoveredWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
// usage, or credit balance and trigger webhooks when a threshold is exceeded.
//
// Alerts created through the API can be scoped to either customers or
// subscriptions.
type CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfiguration struct {
	// Also referred to as alert_id in this documentation.
	ID string `json:"id" api:"required"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The name of the currency the credit balance or invoice cost is denominated in.
	Currency string `json:"currency" api:"required,nullable"`
	// The customer the alert applies to.
	Customer shared.CustomerMinified `json:"customer" api:"required,nullable"`
	// Whether the alert is enabled or disabled.
	Enabled bool `json:"enabled" api:"required"`
	// The metric the alert applies to.
	Metric CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationMetric `json:"metric" api:"required,nullable"`
	// The plan the alert applies to.
	Plan CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPlan `json:"plan" api:"required,nullable"`
	// The subscription the alert applies to.
	Subscription shared.SubscriptionMinified `json:"subscription" api:"required,nullable"`
	// The thresholds that define the conditions under which the alert will be
	// triggered.
	Thresholds []CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThreshold `json:"thresholds" api:"required,nullable"`
	// The type of alert. This must be a valid alert type.
	Type CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationType `json:"type" api:"required"`
	// The current status of the alert. This field is only present for credit balance
	// alerts.
	BalanceAlertStatus []CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationBalanceAlertStatus `json:"balance_alert_status" api:"nullable"`
	// The property keys to group cost alerts by. Only present for cost alerts with
	// grouping enabled.
	GroupingKeys []string `json:"grouping_keys" api:"nullable"`
	// Minified license type for alert serialization.
	LicenseType CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationLicenseType `json:"license_type" api:"nullable"`
	// Filters scoping which prices are included in spend and grouped cost alert
	// evaluation. Alerts use the price_id, item_id, and price_type fields only; the
	// alert's pricing unit is reported by currency.
	PriceFilters []CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFilter `json:"price_filters" api:"nullable"`
	// Per-group threshold overrides. Each override maps a specific combination of
	// grouping_keys values to a replacement threshold list. Only present for grouped
	// cost alerts that have at least one override.
	ThresholdOverrides []CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdOverride `json:"threshold_overrides" api:"nullable"`
	JSON               customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationJSON                `json:"-"`
}

// customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfiguration]
type customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Customer           apijson.Field
	Enabled            apijson.Field
	Metric             apijson.Field
	Plan               apijson.Field
	Subscription       apijson.Field
	Thresholds         apijson.Field
	Type               apijson.Field
	BalanceAlertStatus apijson.Field
	GroupingKeys       apijson.Field
	LicenseType        apijson.Field
	PriceFilters       apijson.Field
	ThresholdOverrides apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationJSON) RawJSON() string {
	return r.raw
}

// The metric the alert applies to.
type CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationMetric struct {
	ID   string                                                                           `json:"id" api:"required"`
	JSON customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationMetricJSON `json:"-"`
}

// customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationMetricJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationMetric]
type customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationMetricJSON) RawJSON() string {
	return r.raw
}

// The plan the alert applies to.
type CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPlan struct {
	ID string `json:"id" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                                                                         `json:"external_plan_id" api:"required,nullable"`
	Name           string                                                                         `json:"name" api:"required,nullable"`
	PlanVersion    string                                                                         `json:"plan_version" api:"required"`
	JSON           customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPlanJSON `json:"-"`
}

// customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPlanJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPlan]
type customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	PlanVersion    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPlanJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                              `json:"value" api:"required"`
	JSON  customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdJSON `json:"-"`
}

// customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThreshold]
type customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdJSON) RawJSON() string {
	return r.raw
}

// The type of alert. This must be a valid alert type.
type CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationType string

const (
	CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted          CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationType = "credit_balance_depleted"
	CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped           CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationType = "credit_balance_dropped"
	CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered         CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationType = "credit_balance_recovered"
	CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationTypeUsageExceeded                  CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationType = "usage_exceeded"
	CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationTypeCostExceeded                   CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationType = "cost_exceeded"
	CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationTypeSpendExceeded                  CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationType = "spend_exceeded"
	CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationType = "license_balance_threshold_reached"
)

func (r CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationType) IsKnown() bool {
	switch r {
	case CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted, CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped, CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered, CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationTypeUsageExceeded, CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationTypeCostExceeded, CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationTypeSpendExceeded, CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached:
		return true
	}
	return false
}

// Alert status is used to determine if an alert is currently in-alert or not.
type CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationBalanceAlertStatus struct {
	// Whether the alert is currently in-alert or not.
	InAlert bool `json:"in_alert" api:"required"`
	// The value of the threshold that defines the alert status.
	ThresholdValue string                                                                                       `json:"threshold_value" api:"required"`
	JSON           customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON `json:"-"`
}

// customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationBalanceAlertStatus]
type customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON struct {
	InAlert        apijson.Field
	ThresholdValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationBalanceAlertStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON) RawJSON() string {
	return r.raw
}

// Minified license type for alert serialization.
type CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationLicenseType struct {
	ID   string                                                                                `json:"id" api:"required"`
	JSON customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationLicenseTypeJSON `json:"-"`
}

// customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationLicenseTypeJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationLicenseType]
type customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationLicenseTypeJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationLicenseType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationLicenseTypeJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFilter struct {
	// The property of the price to filter on.
	Field CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                                              `json:"values" api:"required"`
	JSON   customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFilterJSON `json:"-"`
}

// customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFilterJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFilter]
type customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersField string

const (
	CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID       CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_id"
	CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID        CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersField = "item_id"
	CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType     CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_type"
	CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency      CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersField = "currency"
	CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersField = "pricing_unit_id"
)

func (r CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersField) IsKnown() bool {
	switch r {
	case CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID, CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID, CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType, CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency, CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersOperator string

const (
	CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "includes"
	CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "excludes"
)

func (r CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersOperator) IsKnown() bool {
	switch r {
	case CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes, CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes:
		return true
	}
	return false
}

// A per-group threshold override on a grouped cost alert.
//
// An empty `thresholds` list means the group is silenced (never fires). A
// non-empty list fully replaces the default thresholds for that group.
type CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdOverride struct {
	// The values identifying this group, ordered to match group_keys when set and the
	// alert's grouping_keys otherwise.
	GroupValues []string `json:"group_values" api:"required"`
	// The thresholds applied to this group. An empty list means the group is silenced.
	Thresholds []CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold `json:"thresholds" api:"required"`
	// The subset of the alert's grouping_keys this override binds. Null when the
	// override targets one exact group across every grouping key.
	GroupKeys []string                                                                                    `json:"group_keys" api:"nullable"`
	JSON      customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON `json:"-"`
}

// customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdOverride]
type customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON struct {
	GroupValues apijson.Field
	Thresholds  apijson.Field
	GroupKeys   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                                                `json:"value" api:"required"`
	JSON  customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON `json:"-"`
}

// customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON
// contains the JSON metadata for the struct
// [CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold]
type customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceRecoveredWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON) RawJSON() string {
	return r.raw
}

// A currency or custom credit unit, as embedded in webhook payloads.
type CustomerCreditBalanceRecoveredWebhookEventPropertiesPricingUnit struct {
	ID          string                                                              `json:"id" api:"required"`
	DisplayName string                                                              `json:"display_name" api:"required,nullable"`
	Name        string                                                              `json:"name" api:"required"`
	Symbol      string                                                              `json:"symbol" api:"required,nullable"`
	JSON        customerCreditBalanceRecoveredWebhookEventPropertiesPricingUnitJSON `json:"-"`
}

// customerCreditBalanceRecoveredWebhookEventPropertiesPricingUnitJSON contains the
// JSON metadata for the struct
// [CustomerCreditBalanceRecoveredWebhookEventPropertiesPricingUnit]
type customerCreditBalanceRecoveredWebhookEventPropertiesPricingUnitJSON struct {
	ID          apijson.Field
	DisplayName apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditBalanceRecoveredWebhookEventPropertiesPricingUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditBalanceRecoveredWebhookEventPropertiesPricingUnitJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type CustomerCreditBalanceRecoveredWebhookEventType string

const (
	CustomerCreditBalanceRecoveredWebhookEventTypeCustomerCreditBalanceRecovered CustomerCreditBalanceRecoveredWebhookEventType = "customer.credit_balance_recovered"
)

func (r CustomerCreditBalanceRecoveredWebhookEventType) IsKnown() bool {
	switch r {
	case CustomerCreditBalanceRecoveredWebhookEventTypeCustomerCreditBalanceRecovered:
		return true
	}
	return false
}

// Issued when a customer's credit ledger is incremented.
type CustomerCreditLedgerIncrementedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
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
	Customer   Customer                                              `json:"customer" api:"required"`
	Properties CustomerCreditLedgerIncrementedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type CustomerCreditLedgerIncrementedWebhookEventType `json:"type" api:"required"`
	JSON customerCreditLedgerIncrementedWebhookEventJSON `json:"-"`
}

// customerCreditLedgerIncrementedWebhookEventJSON contains the JSON metadata for
// the struct [CustomerCreditLedgerIncrementedWebhookEvent]
type customerCreditLedgerIncrementedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Customer    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditLedgerIncrementedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerIncrementedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CustomerCreditLedgerIncrementedWebhookEvent) implementsUnwrapWebhookEvent() {}

type CustomerCreditLedgerIncrementedWebhookEventProperties struct {
	// The Credit Block resource models prepaid credits within Orb.
	Block CustomerCreditLedgerIncrementedWebhookEventPropertiesBlock `json:"block" api:"required"`
	// A currency or custom credit unit, as embedded in webhook payloads.
	PricingUnit CustomerCreditLedgerIncrementedWebhookEventPropertiesPricingUnit `json:"pricing_unit" api:"required"`
	JSON        customerCreditLedgerIncrementedWebhookEventPropertiesJSON        `json:"-"`
}

// customerCreditLedgerIncrementedWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [CustomerCreditLedgerIncrementedWebhookEventProperties]
type customerCreditLedgerIncrementedWebhookEventPropertiesJSON struct {
	Block       apijson.Field
	PricingUnit apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditLedgerIncrementedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerIncrementedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The Credit Block resource models prepaid credits within Orb.
type CustomerCreditLedgerIncrementedWebhookEventPropertiesBlock struct {
	ID      string `json:"id" api:"required"`
	Balance string `json:"balance" api:"required"`
	// How this credit block was created: `allocation` (a subscription's recurring
	// credit allocation), `top_up` (an automatic balance-threshold top-up),
	// `commitment` (a subscription commitment true-up rolled forward as credit), or
	// `manual` (a manual credit ledger increment, including credits voided or expired
	// off another block).
	CreditBlockSource     CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditBlockSource `json:"credit_block_source" api:"required"`
	EffectiveDate         time.Time                                                                   `json:"effective_date" api:"required,nullable" format:"date-time"`
	ExpiryDate            time.Time                                                                   `json:"expiry_date" api:"required,nullable" format:"date-time"`
	Filters               []CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFilter          `json:"filters" api:"required"`
	MaximumInitialBalance string                                                                      `json:"maximum_initial_balance" api:"required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata         map[string]string                                                `json:"metadata" api:"required"`
	PerUnitCostBasis string                                                           `json:"per_unit_cost_basis" api:"required,nullable"`
	Status           CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockStatus `json:"status" api:"required"`
	// The credit allocation that funded a block. Extends the allocation resource
	// serialized on prices with the catalog-item attribution of the funding price.
	CreditAllocation CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocation `json:"credit_allocation" api:"nullable"`
	// The subscription commitment whose true-up rolled forward into this credit block.
	// Present only when `credit_block_source` is `commitment`.
	CreditCommitment CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditCommitment `json:"credit_commitment" api:"nullable"`
	JSON             customerCreditLedgerIncrementedWebhookEventPropertiesBlockJSON             `json:"-"`
}

// customerCreditLedgerIncrementedWebhookEventPropertiesBlockJSON contains the JSON
// metadata for the struct
// [CustomerCreditLedgerIncrementedWebhookEventPropertiesBlock]
type customerCreditLedgerIncrementedWebhookEventPropertiesBlockJSON struct {
	ID                    apijson.Field
	Balance               apijson.Field
	CreditBlockSource     apijson.Field
	EffectiveDate         apijson.Field
	ExpiryDate            apijson.Field
	Filters               apijson.Field
	MaximumInitialBalance apijson.Field
	Metadata              apijson.Field
	PerUnitCostBasis      apijson.Field
	Status                apijson.Field
	CreditAllocation      apijson.Field
	CreditCommitment      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CustomerCreditLedgerIncrementedWebhookEventPropertiesBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerIncrementedWebhookEventPropertiesBlockJSON) RawJSON() string {
	return r.raw
}

// How this credit block was created: `allocation` (a subscription's recurring
// credit allocation), `top_up` (an automatic balance-threshold top-up),
// `commitment` (a subscription commitment true-up rolled forward as credit), or
// `manual` (a manual credit ledger increment, including credits voided or expired
// off another block).
type CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditBlockSource string

const (
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditBlockSourceAllocation CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditBlockSource = "allocation"
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditBlockSourceTopUp      CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditBlockSource = "top_up"
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditBlockSourceCommitment CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditBlockSource = "commitment"
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditBlockSourceManual     CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditBlockSource = "manual"
)

func (r CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditBlockSource) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditBlockSourceAllocation, CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditBlockSourceTopUp, CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditBlockSourceCommitment, CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditBlockSourceManual:
		return true
	}
	return false
}

type CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFilter struct {
	// The property of the price to filter on.
	Field CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                             `json:"values" api:"required"`
	JSON   customerCreditLedgerIncrementedWebhookEventPropertiesBlockFilterJSON `json:"-"`
}

// customerCreditLedgerIncrementedWebhookEventPropertiesBlockFilterJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFilter]
type customerCreditLedgerIncrementedWebhookEventPropertiesBlockFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerIncrementedWebhookEventPropertiesBlockFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersField string

const (
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersFieldPriceID       CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersField = "price_id"
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersFieldItemID        CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersField = "item_id"
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersFieldPriceType     CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersField = "price_type"
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersFieldCurrency      CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersField = "currency"
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersFieldPricingUnitID CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersField = "pricing_unit_id"
)

func (r CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersField) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersFieldPriceID, CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersFieldItemID, CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersFieldPriceType, CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersFieldCurrency, CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersOperator string

const (
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersOperatorIncludes CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersOperator = "includes"
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersOperatorExcludes CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersOperator = "excludes"
)

func (r CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersOperator) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersOperatorIncludes, CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockFiltersOperatorExcludes:
		return true
	}
	return false
}

type CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockStatus string

const (
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockStatusActive         CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockStatus = "active"
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockStatusPendingPayment CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockStatus = "pending_payment"
)

func (r CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockStatus) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockStatusActive, CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockStatusPendingPayment:
		return true
	}
	return false
}

// The credit allocation that funded a block. Extends the allocation resource
// serialized on prices with the catalog-item attribution of the funding price.
type CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocation struct {
	AllowsRollover   bool                    `json:"allows_rollover" api:"required"`
	Currency         string                  `json:"currency" api:"required"`
	CustomExpiration shared.CustomExpiration `json:"custom_expiration" api:"required,nullable"`
	// The ID of the catalog item this block was allocated from, derived from the
	// allocation's price.
	ItemID        string                                                                             `json:"item_id" api:"required"`
	Filters       []CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFilter `json:"filters"`
	LicenseTypeID string                                                                             `json:"license_type_id" api:"nullable"`
	JSON          customerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationJSON     `json:"-"`
}

// customerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocation]
type customerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationJSON struct {
	AllowsRollover   apijson.Field
	Currency         apijson.Field
	CustomExpiration apijson.Field
	ItemID           apijson.Field
	Filters          apijson.Field
	LicenseTypeID    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFilter struct {
	// The property of the price to filter on.
	Field CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                                             `json:"values" api:"required"`
	JSON   customerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFilterJSON `json:"-"`
}

// customerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFilterJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFilter]
type customerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersField string

const (
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersFieldPriceID       CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersField = "price_id"
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersFieldItemID        CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersField = "item_id"
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersFieldPriceType     CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersField = "price_type"
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersFieldCurrency      CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersField = "currency"
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersFieldPricingUnitID CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersField = "pricing_unit_id"
)

func (r CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersField) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersFieldPriceID, CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersFieldItemID, CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersFieldPriceType, CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersFieldCurrency, CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersOperator string

const (
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersOperatorIncludes CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersOperator = "includes"
	CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersOperatorExcludes CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersOperator = "excludes"
)

func (r CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersOperator) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersOperatorIncludes, CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditAllocationFiltersOperatorExcludes:
		return true
	}
	return false
}

// The subscription commitment whose true-up rolled forward into this credit block.
// Present only when `credit_block_source` is `commitment`.
type CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditCommitment struct {
	// The ID of the subscription commitment this block was rolled forward from.
	ID string `json:"id" api:"required"`
	// The subscription the commitment belongs to.
	SubscriptionID string                                                                         `json:"subscription_id" api:"nullable"`
	JSON           customerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditCommitmentJSON `json:"-"`
}

// customerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditCommitmentJSON
// contains the JSON metadata for the struct
// [CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditCommitment]
type customerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditCommitmentJSON struct {
	ID             apijson.Field
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditCommitment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerIncrementedWebhookEventPropertiesBlockCreditCommitmentJSON) RawJSON() string {
	return r.raw
}

// A currency or custom credit unit, as embedded in webhook payloads.
type CustomerCreditLedgerIncrementedWebhookEventPropertiesPricingUnit struct {
	ID          string                                                               `json:"id" api:"required"`
	DisplayName string                                                               `json:"display_name" api:"required,nullable"`
	Name        string                                                               `json:"name" api:"required"`
	Symbol      string                                                               `json:"symbol" api:"required,nullable"`
	JSON        customerCreditLedgerIncrementedWebhookEventPropertiesPricingUnitJSON `json:"-"`
}

// customerCreditLedgerIncrementedWebhookEventPropertiesPricingUnitJSON contains
// the JSON metadata for the struct
// [CustomerCreditLedgerIncrementedWebhookEventPropertiesPricingUnit]
type customerCreditLedgerIncrementedWebhookEventPropertiesPricingUnitJSON struct {
	ID          apijson.Field
	DisplayName apijson.Field
	Name        apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditLedgerIncrementedWebhookEventPropertiesPricingUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditLedgerIncrementedWebhookEventPropertiesPricingUnitJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type CustomerCreditLedgerIncrementedWebhookEventType string

const (
	CustomerCreditLedgerIncrementedWebhookEventTypeCustomerCreditLedgerIncremented CustomerCreditLedgerIncrementedWebhookEventType = "customer.credit_ledger_incremented"
)

func (r CustomerCreditLedgerIncrementedWebhookEventType) IsKnown() bool {
	switch r {
	case CustomerCreditLedgerIncrementedWebhookEventTypeCustomerCreditLedgerIncremented:
		return true
	}
	return false
}

// Issued when a customer is updated.
type CustomerEditedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
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
	Customer   Customer                             `json:"customer" api:"required"`
	Properties CustomerEditedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type CustomerEditedWebhookEventType `json:"type" api:"required"`
	JSON customerEditedWebhookEventJSON `json:"-"`
}

// customerEditedWebhookEventJSON contains the JSON metadata for the struct
// [CustomerEditedWebhookEvent]
type customerEditedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Customer    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerEditedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerEditedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r CustomerEditedWebhookEvent) implementsUnwrapWebhookEvent() {}

type CustomerEditedWebhookEventProperties struct {
	// metadata values are non-null on the wire (deleting a key removes it from
	// storage); the Optional[str] values only exist on the new half of a metadata
	// FieldChange.
	PreviousAttributes CustomerEditedWebhookEventPropertiesPreviousAttributes `json:"previous_attributes" api:"required"`
	JSON               customerEditedWebhookEventPropertiesJSON               `json:"-"`
}

// customerEditedWebhookEventPropertiesJSON contains the JSON metadata for the
// struct [CustomerEditedWebhookEventProperties]
type customerEditedWebhookEventPropertiesJSON struct {
	PreviousAttributes apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerEditedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerEditedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// metadata values are non-null on the wire (deleting a key removes it from
// storage); the Optional[str] values only exist on the new half of a metadata
// FieldChange.
type CustomerEditedWebhookEventPropertiesPreviousAttributes struct {
	AutoCollection         bool              `json:"auto_collection" api:"nullable"`
	BillingAddress         shared.Address    `json:"billing_address" api:"nullable"`
	DefaultPaymentMethodID string            `json:"default_payment_method_id" api:"nullable"`
	Email                  string            `json:"email" api:"nullable"`
	EmailDelivery          bool              `json:"email_delivery" api:"nullable"`
	ExternalCustomerID     string            `json:"external_customer_id" api:"nullable"`
	Metadata               map[string]string `json:"metadata" api:"nullable"`
	Name                   string            `json:"name" api:"nullable"`
	PaymentProvider        string            `json:"payment_provider" api:"nullable"`
	PaymentProviderID      string            `json:"payment_provider_id" api:"nullable"`
	ShippingAddress        shared.Address    `json:"shipping_address" api:"nullable"`
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
	// | Faroe Islands          | `fo_vat`     | Faroe Islands VAT Number                                                                                |
	// | Finland                | `eu_vat`     | European VAT Number                                                                                     |
	// | France                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Georgia                | `ge_vat`     | Georgian VAT                                                                                            |
	// | Germany                | `de_stn`     | German Tax Number (Steuernummer)                                                                        |
	// | Germany                | `eu_vat`     | European VAT Number                                                                                     |
	// | Gibraltar              | `gi_tin`     | Gibraltar Tax Identification Number                                                                     |
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
	// | Italy                  | `it_cf`      | Italian Codice Fiscale Number                                                                           |
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
	// | Paraguay               | `py_ruc`     | Paraguayan RUC Number                                                                                   |
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
	// | Sri Lanka              | `lk_vat`     | Sri Lanka VAT Number                                                                                    |
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
	TaxID shared.CustomerTaxID                                       `json:"tax_id" api:"nullable"`
	JSON  customerEditedWebhookEventPropertiesPreviousAttributesJSON `json:"-"`
}

// customerEditedWebhookEventPropertiesPreviousAttributesJSON contains the JSON
// metadata for the struct [CustomerEditedWebhookEventPropertiesPreviousAttributes]
type customerEditedWebhookEventPropertiesPreviousAttributesJSON struct {
	AutoCollection         apijson.Field
	BillingAddress         apijson.Field
	DefaultPaymentMethodID apijson.Field
	Email                  apijson.Field
	EmailDelivery          apijson.Field
	ExternalCustomerID     apijson.Field
	Metadata               apijson.Field
	Name                   apijson.Field
	PaymentProvider        apijson.Field
	PaymentProviderID      apijson.Field
	ShippingAddress        apijson.Field
	TaxID                  apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CustomerEditedWebhookEventPropertiesPreviousAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerEditedWebhookEventPropertiesPreviousAttributesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type CustomerEditedWebhookEventType string

const (
	CustomerEditedWebhookEventTypeCustomerEdited CustomerEditedWebhookEventType = "customer.edited"
)

func (r CustomerEditedWebhookEventType) IsKnown() bool {
	switch r {
	case CustomerEditedWebhookEventTypeCustomerEdited:
		return true
	}
	return false
}

// Issued when a data export transfer fails.
type DataExportsTransferErrorWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                      `json:"created_at" api:"required" format:"date-time"`
	Properties DataExportsTransferErrorWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type DataExportsTransferErrorWebhookEventType `json:"type" api:"required"`
	JSON dataExportsTransferErrorWebhookEventJSON `json:"-"`
}

// dataExportsTransferErrorWebhookEventJSON contains the JSON metadata for the
// struct [DataExportsTransferErrorWebhookEvent]
type dataExportsTransferErrorWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DataExportsTransferErrorWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportsTransferErrorWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DataExportsTransferErrorWebhookEvent) implementsUnwrapWebhookEvent() {}

type DataExportsTransferErrorWebhookEventProperties struct {
	Description         string                                             `json:"description" api:"required"`
	DestinationName     string                                             `json:"destination_name" api:"required"`
	Resources           []string                                           `json:"resources" api:"required"`
	RowsTransferred     int64                                              `json:"rows_transferred" api:"required"`
	TransferBlamedParty string                                             `json:"transfer_blamed_party" api:"required"`
	TransferEndedAt     time.Time                                          `json:"transfer_ended_at" api:"required" format:"date-time"`
	TransferStartedAt   time.Time                                          `json:"transfer_started_at" api:"required" format:"date-time"`
	JSON                dataExportsTransferErrorWebhookEventPropertiesJSON `json:"-"`
}

// dataExportsTransferErrorWebhookEventPropertiesJSON contains the JSON metadata
// for the struct [DataExportsTransferErrorWebhookEventProperties]
type dataExportsTransferErrorWebhookEventPropertiesJSON struct {
	Description         apijson.Field
	DestinationName     apijson.Field
	Resources           apijson.Field
	RowsTransferred     apijson.Field
	TransferBlamedParty apijson.Field
	TransferEndedAt     apijson.Field
	TransferStartedAt   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DataExportsTransferErrorWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportsTransferErrorWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type DataExportsTransferErrorWebhookEventType string

const (
	DataExportsTransferErrorWebhookEventTypeDataExportsTransferError DataExportsTransferErrorWebhookEventType = "data_exports.transfer_error"
)

func (r DataExportsTransferErrorWebhookEventType) IsKnown() bool {
	switch r {
	case DataExportsTransferErrorWebhookEventTypeDataExportsTransferError:
		return true
	}
	return false
}

// Issued when a data export transfer succeeds.
type DataExportsTransferSuccessWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                        `json:"created_at" api:"required" format:"date-time"`
	Properties DataExportsTransferSuccessWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type DataExportsTransferSuccessWebhookEventType `json:"type" api:"required"`
	JSON dataExportsTransferSuccessWebhookEventJSON `json:"-"`
}

// dataExportsTransferSuccessWebhookEventJSON contains the JSON metadata for the
// struct [DataExportsTransferSuccessWebhookEvent]
type dataExportsTransferSuccessWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DataExportsTransferSuccessWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportsTransferSuccessWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r DataExportsTransferSuccessWebhookEvent) implementsUnwrapWebhookEvent() {}

type DataExportsTransferSuccessWebhookEventProperties struct {
	Description       string                                               `json:"description" api:"required"`
	DestinationName   string                                               `json:"destination_name" api:"required"`
	Resources         []string                                             `json:"resources" api:"required"`
	RowsTransferred   int64                                                `json:"rows_transferred" api:"required"`
	TransferEndedAt   time.Time                                            `json:"transfer_ended_at" api:"required" format:"date-time"`
	TransferStartedAt time.Time                                            `json:"transfer_started_at" api:"required" format:"date-time"`
	JSON              dataExportsTransferSuccessWebhookEventPropertiesJSON `json:"-"`
}

// dataExportsTransferSuccessWebhookEventPropertiesJSON contains the JSON metadata
// for the struct [DataExportsTransferSuccessWebhookEventProperties]
type dataExportsTransferSuccessWebhookEventPropertiesJSON struct {
	Description       apijson.Field
	DestinationName   apijson.Field
	Resources         apijson.Field
	RowsTransferred   apijson.Field
	TransferEndedAt   apijson.Field
	TransferStartedAt apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DataExportsTransferSuccessWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportsTransferSuccessWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type DataExportsTransferSuccessWebhookEventType string

const (
	DataExportsTransferSuccessWebhookEventTypeDataExportsTransferSuccess DataExportsTransferSuccessWebhookEventType = "data_exports.transfer_success"
)

func (r DataExportsTransferSuccessWebhookEventType) IsKnown() bool {
	switch r {
	case DataExportsTransferSuccessWebhookEventTypeDataExportsTransferSuccess:
		return true
	}
	return false
}

// Issued when an event does not match any customer.
type EventUnmatchedEventWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                 `json:"created_at" api:"required" format:"date-time"`
	Properties EventUnmatchedEventWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type EventUnmatchedEventWebhookEventType `json:"type" api:"required"`
	JSON eventUnmatchedEventWebhookEventJSON `json:"-"`
}

// eventUnmatchedEventWebhookEventJSON contains the JSON metadata for the struct
// [EventUnmatchedEventWebhookEvent]
type eventUnmatchedEventWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventUnmatchedEventWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventUnmatchedEventWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r EventUnmatchedEventWebhookEvent) implementsUnwrapWebhookEvent() {}

type EventUnmatchedEventWebhookEventProperties struct {
	Event EventUnmatchedEventWebhookEventPropertiesEvent `json:"event" api:"required"`
	JSON  eventUnmatchedEventWebhookEventPropertiesJSON  `json:"-"`
}

// eventUnmatchedEventWebhookEventPropertiesJSON contains the JSON metadata for the
// struct [EventUnmatchedEventWebhookEventProperties]
type eventUnmatchedEventWebhookEventPropertiesJSON struct {
	Event       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventUnmatchedEventWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventUnmatchedEventWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

type EventUnmatchedEventWebhookEventPropertiesEvent struct {
	CustomerID         string                                             `json:"customer_id" api:"required,nullable"`
	EventName          string                                             `json:"event_name" api:"required"`
	ExternalCustomerID string                                             `json:"external_customer_id" api:"required,nullable"`
	IdempotencyKey     string                                             `json:"idempotency_key" api:"required"`
	Properties         map[string]interface{}                             `json:"properties" api:"required"`
	Timestamp          time.Time                                          `json:"timestamp" api:"required" format:"date-time"`
	JSON               eventUnmatchedEventWebhookEventPropertiesEventJSON `json:"-"`
}

// eventUnmatchedEventWebhookEventPropertiesEventJSON contains the JSON metadata
// for the struct [EventUnmatchedEventWebhookEventPropertiesEvent]
type eventUnmatchedEventWebhookEventPropertiesEventJSON struct {
	CustomerID         apijson.Field
	EventName          apijson.Field
	ExternalCustomerID apijson.Field
	IdempotencyKey     apijson.Field
	Properties         apijson.Field
	Timestamp          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *EventUnmatchedEventWebhookEventPropertiesEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventUnmatchedEventWebhookEventPropertiesEventJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type EventUnmatchedEventWebhookEventType string

const (
	EventUnmatchedEventWebhookEventTypeEventUnmatchedEvent EventUnmatchedEventWebhookEventType = "event.unmatched_event"
)

func (r EventUnmatchedEventWebhookEventType) IsKnown() bool {
	switch r {
	case EventUnmatchedEventWebhookEventTypeEventUnmatchedEvent:
		return true
	}
	return false
}

// Issued when ingestion events reference unmatched customer IDs.
type IngestionUnmatchedCustomerIDsWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                           `json:"created_at" api:"required" format:"date-time"`
	Properties IngestionUnmatchedCustomerIDsWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type IngestionUnmatchedCustomerIDsWebhookEventType `json:"type" api:"required"`
	JSON ingestionUnmatchedCustomerIDsWebhookEventJSON `json:"-"`
}

// ingestionUnmatchedCustomerIDsWebhookEventJSON contains the JSON metadata for the
// struct [IngestionUnmatchedCustomerIDsWebhookEvent]
type ingestionUnmatchedCustomerIDsWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IngestionUnmatchedCustomerIDsWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ingestionUnmatchedCustomerIDsWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r IngestionUnmatchedCustomerIDsWebhookEvent) implementsUnwrapWebhookEvent() {}

type IngestionUnmatchedCustomerIDsWebhookEventProperties struct {
	ExternalCustomerIDs []string                                                `json:"external_customer_ids" api:"required"`
	JSON                ingestionUnmatchedCustomerIDsWebhookEventPropertiesJSON `json:"-"`
}

// ingestionUnmatchedCustomerIDsWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [IngestionUnmatchedCustomerIDsWebhookEventProperties]
type ingestionUnmatchedCustomerIDsWebhookEventPropertiesJSON struct {
	ExternalCustomerIDs apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *IngestionUnmatchedCustomerIDsWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ingestionUnmatchedCustomerIDsWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type IngestionUnmatchedCustomerIDsWebhookEventType string

const (
	IngestionUnmatchedCustomerIDsWebhookEventTypeIngestionUnmatchedCustomerIDs IngestionUnmatchedCustomerIDsWebhookEventType = "ingestion.unmatched_customer_ids"
)

func (r IngestionUnmatchedCustomerIDsWebhookEventType) IsKnown() bool {
	switch r {
	case IngestionUnmatchedCustomerIDsWebhookEventTypeIngestionUnmatchedCustomerIDs:
		return true
	}
	return false
}

// Issued when an invoice accounting sync fails.
type InvoiceAccountingSyncFailedWebhookEvent struct {
	// The ID of this webhook event.
	ID                   string                                                      `json:"id" api:"required"`
	AccountingSyncRecord InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecord `json:"accounting_sync_record" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice                                    `json:"invoice" api:"required"`
	Properties InvoiceAccountingSyncFailedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceAccountingSyncFailedWebhookEventType `json:"type" api:"required"`
	JSON invoiceAccountingSyncFailedWebhookEventJSON `json:"-"`
}

// invoiceAccountingSyncFailedWebhookEventJSON contains the JSON metadata for the
// struct [InvoiceAccountingSyncFailedWebhookEvent]
type invoiceAccountingSyncFailedWebhookEventJSON struct {
	ID                   apijson.Field
	AccountingSyncRecord apijson.Field
	CreatedAt            apijson.Field
	Invoice              apijson.Field
	Properties           apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InvoiceAccountingSyncFailedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceAccountingSyncFailedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceAccountingSyncFailedWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecord struct {
	ID                 string                                                                `json:"id" api:"required"`
	CustomerID         string                                                                `json:"customer_id" api:"required"`
	RecordType         InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType `json:"record_type" api:"required"`
	ErrorDetails       map[string]interface{}                                                `json:"error_details" api:"nullable"`
	InvoiceID          string                                                                `json:"invoice_id" api:"nullable"`
	ProviderCustomerID string                                                                `json:"provider_customer_id" api:"nullable"`
	Status             string                                                                `json:"status" api:"nullable"`
	SyncAction         string                                                                `json:"sync_action" api:"nullable"`
	JSON               invoiceAccountingSyncFailedWebhookEventAccountingSyncRecordJSON       `json:"-"`
}

// invoiceAccountingSyncFailedWebhookEventAccountingSyncRecordJSON contains the
// JSON metadata for the struct
// [InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecord]
type invoiceAccountingSyncFailedWebhookEventAccountingSyncRecordJSON struct {
	ID                 apijson.Field
	CustomerID         apijson.Field
	RecordType         apijson.Field
	ErrorDetails       apijson.Field
	InvoiceID          apijson.Field
	ProviderCustomerID apijson.Field
	Status             apijson.Field
	SyncAction         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceAccountingSyncFailedWebhookEventAccountingSyncRecordJSON) RawJSON() string {
	return r.raw
}

type InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType string

const (
	InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomer                   InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "customer"
	InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeInvoice                    InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "invoice"
	InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeTransaction                InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "transaction"
	InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "customer_balance_transaction"
	InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCreditNote                 InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "credit_note"
	InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSubscription               InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "subscription"
	InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSalesOrder                 InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "sales_order"
	InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeBlock                      InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "block"
)

func (r InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType) IsKnown() bool {
	switch r {
	case InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomer, InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeInvoice, InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeTransaction, InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction, InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCreditNote, InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSubscription, InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSalesOrder, InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeBlock:
		return true
	}
	return false
}

type InvoiceAccountingSyncFailedWebhookEventProperties struct {
	ConnectionType string                                                `json:"connection_type" api:"required"`
	FailureReason  string                                                `json:"failure_reason" api:"required"`
	JSON           invoiceAccountingSyncFailedWebhookEventPropertiesJSON `json:"-"`
}

// invoiceAccountingSyncFailedWebhookEventPropertiesJSON contains the JSON metadata
// for the struct [InvoiceAccountingSyncFailedWebhookEventProperties]
type invoiceAccountingSyncFailedWebhookEventPropertiesJSON struct {
	ConnectionType apijson.Field
	FailureReason  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InvoiceAccountingSyncFailedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceAccountingSyncFailedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoiceAccountingSyncFailedWebhookEventType string

const (
	InvoiceAccountingSyncFailedWebhookEventTypeInvoiceAccountingSyncFailed InvoiceAccountingSyncFailedWebhookEventType = "invoice.accounting_sync_failed"
)

func (r InvoiceAccountingSyncFailedWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceAccountingSyncFailedWebhookEventTypeInvoiceAccountingSyncFailed:
		return true
	}
	return false
}

// Issued when an invoice accounting sync succeeds.
type InvoiceAccountingSyncSucceededWebhookEvent struct {
	// The ID of this webhook event.
	ID                   string                                                         `json:"id" api:"required"`
	AccountingSyncRecord InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecord `json:"accounting_sync_record" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice                                       `json:"invoice" api:"required"`
	Properties InvoiceAccountingSyncSucceededWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceAccountingSyncSucceededWebhookEventType `json:"type" api:"required"`
	JSON invoiceAccountingSyncSucceededWebhookEventJSON `json:"-"`
}

// invoiceAccountingSyncSucceededWebhookEventJSON contains the JSON metadata for
// the struct [InvoiceAccountingSyncSucceededWebhookEvent]
type invoiceAccountingSyncSucceededWebhookEventJSON struct {
	ID                   apijson.Field
	AccountingSyncRecord apijson.Field
	CreatedAt            apijson.Field
	Invoice              apijson.Field
	Properties           apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InvoiceAccountingSyncSucceededWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceAccountingSyncSucceededWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceAccountingSyncSucceededWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecord struct {
	ID                 string                                                                   `json:"id" api:"required"`
	CustomerID         string                                                                   `json:"customer_id" api:"required"`
	RecordType         InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType `json:"record_type" api:"required"`
	ErrorDetails       map[string]interface{}                                                   `json:"error_details" api:"nullable"`
	InvoiceID          string                                                                   `json:"invoice_id" api:"nullable"`
	ProviderCustomerID string                                                                   `json:"provider_customer_id" api:"nullable"`
	Status             string                                                                   `json:"status" api:"nullable"`
	SyncAction         string                                                                   `json:"sync_action" api:"nullable"`
	JSON               invoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON       `json:"-"`
}

// invoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON contains the
// JSON metadata for the struct
// [InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecord]
type invoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON struct {
	ID                 apijson.Field
	CustomerID         apijson.Field
	RecordType         apijson.Field
	ErrorDetails       apijson.Field
	InvoiceID          apijson.Field
	ProviderCustomerID apijson.Field
	Status             apijson.Field
	SyncAction         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON) RawJSON() string {
	return r.raw
}

type InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType string

const (
	InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomer                   InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "customer"
	InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeInvoice                    InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "invoice"
	InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeTransaction                InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "transaction"
	InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "customer_balance_transaction"
	InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCreditNote                 InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "credit_note"
	InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSubscription               InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "subscription"
	InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSalesOrder                 InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "sales_order"
	InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeBlock                      InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "block"
)

func (r InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType) IsKnown() bool {
	switch r {
	case InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomer, InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeInvoice, InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeTransaction, InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction, InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCreditNote, InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSubscription, InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSalesOrder, InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeBlock:
		return true
	}
	return false
}

type InvoiceAccountingSyncSucceededWebhookEventProperties struct {
	ConnectionType string                                                   `json:"connection_type" api:"required"`
	JSON           invoiceAccountingSyncSucceededWebhookEventPropertiesJSON `json:"-"`
}

// invoiceAccountingSyncSucceededWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [InvoiceAccountingSyncSucceededWebhookEventProperties]
type invoiceAccountingSyncSucceededWebhookEventPropertiesJSON struct {
	ConnectionType apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InvoiceAccountingSyncSucceededWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceAccountingSyncSucceededWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoiceAccountingSyncSucceededWebhookEventType string

const (
	InvoiceAccountingSyncSucceededWebhookEventTypeInvoiceAccountingSyncSucceeded InvoiceAccountingSyncSucceededWebhookEventType = "invoice.accounting_sync_succeeded"
)

func (r InvoiceAccountingSyncSucceededWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceAccountingSyncSucceededWebhookEventTypeInvoiceAccountingSyncSucceeded:
		return true
	}
	return false
}

// Issued when a collections-automation schedule step is executed for an invoice.
type InvoiceAutomationScheduleStepExecutedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The invoice fields a consumer needs to run their own notification flows without
	// a follow-up API call, mirroring the variables Orb's own automation emails render
	// against.
	Invoice    InvoiceAutomationScheduleStepExecutedWebhookEventInvoice    `json:"invoice" api:"required"`
	Properties InvoiceAutomationScheduleStepExecutedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceAutomationScheduleStepExecutedWebhookEventType `json:"type" api:"required"`
	JSON invoiceAutomationScheduleStepExecutedWebhookEventJSON `json:"-"`
}

// invoiceAutomationScheduleStepExecutedWebhookEventJSON contains the JSON metadata
// for the struct [InvoiceAutomationScheduleStepExecutedWebhookEvent]
type invoiceAutomationScheduleStepExecutedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceAutomationScheduleStepExecutedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceAutomationScheduleStepExecutedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceAutomationScheduleStepExecutedWebhookEvent) implementsUnwrapWebhookEvent() {}

// The invoice fields a consumer needs to run their own notification flows without
// a follow-up API call, mirroring the variables Orb's own automation emails render
// against.
type InvoiceAutomationScheduleStepExecutedWebhookEventInvoice struct {
	ID                          string                                                       `json:"id" api:"required"`
	AmountDue                   string                                                       `json:"amount_due" api:"required"`
	Currency                    string                                                       `json:"currency" api:"required"`
	CustomerID                  string                                                       `json:"customer_id" api:"required"`
	CustomerName                string                                                       `json:"customer_name" api:"required"`
	DueDate                     time.Time                                                    `json:"due_date" api:"required,nullable" format:"date"`
	ExternalCustomerID          string                                                       `json:"external_customer_id" api:"required,nullable"`
	HostedInvoiceURL            string                                                       `json:"hosted_invoice_url" api:"required,nullable"`
	InvoiceDate                 time.Time                                                    `json:"invoice_date" api:"required" format:"date-time"`
	InvoiceNumber               string                                                       `json:"invoice_number" api:"required"`
	IssuedAt                    time.Time                                                    `json:"issued_at" api:"required,nullable" format:"date-time"`
	Memo                        string                                                       `json:"memo" api:"required,nullable"`
	PaymentMethodLastFourDigits string                                                       `json:"payment_method_last_four_digits" api:"required,nullable"`
	Status                      string                                                       `json:"status" api:"required"`
	SubscriptionID              string                                                       `json:"subscription_id" api:"required,nullable"`
	JSON                        invoiceAutomationScheduleStepExecutedWebhookEventInvoiceJSON `json:"-"`
}

// invoiceAutomationScheduleStepExecutedWebhookEventInvoiceJSON contains the JSON
// metadata for the struct
// [InvoiceAutomationScheduleStepExecutedWebhookEventInvoice]
type invoiceAutomationScheduleStepExecutedWebhookEventInvoiceJSON struct {
	ID                          apijson.Field
	AmountDue                   apijson.Field
	Currency                    apijson.Field
	CustomerID                  apijson.Field
	CustomerName                apijson.Field
	DueDate                     apijson.Field
	ExternalCustomerID          apijson.Field
	HostedInvoiceURL            apijson.Field
	InvoiceDate                 apijson.Field
	InvoiceNumber               apijson.Field
	IssuedAt                    apijson.Field
	Memo                        apijson.Field
	PaymentMethodLastFourDigits apijson.Field
	Status                      apijson.Field
	SubscriptionID              apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *InvoiceAutomationScheduleStepExecutedWebhookEventInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceAutomationScheduleStepExecutedWebhookEventInvoiceJSON) RawJSON() string {
	return r.raw
}

type InvoiceAutomationScheduleStepExecutedWebhookEventProperties struct {
	Actions                        []InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesAction `json:"actions" api:"required"`
	AutomationScheduleTemplateID   string                                                              `json:"automation_schedule_template_id" api:"required,nullable"`
	AutomationScheduleTemplateName string                                                              `json:"automation_schedule_template_name" api:"required,nullable"`
	ExecutedAt                     time.Time                                                           `json:"executed_at" api:"required" format:"date-time"`
	Label                          string                                                              `json:"label" api:"required"`
	ScheduledAt                    time.Time                                                           `json:"scheduled_at" api:"required" format:"date-time"`
	StepID                         string                                                              `json:"step_id" api:"required"`
	JSON                           invoiceAutomationScheduleStepExecutedWebhookEventPropertiesJSON     `json:"-"`
}

// invoiceAutomationScheduleStepExecutedWebhookEventPropertiesJSON contains the
// JSON metadata for the struct
// [InvoiceAutomationScheduleStepExecutedWebhookEventProperties]
type invoiceAutomationScheduleStepExecutedWebhookEventPropertiesJSON struct {
	Actions                        apijson.Field
	AutomationScheduleTemplateID   apijson.Field
	AutomationScheduleTemplateName apijson.Field
	ExecutedAt                     apijson.Field
	Label                          apijson.Field
	ScheduledAt                    apijson.Field
	StepID                         apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *InvoiceAutomationScheduleStepExecutedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceAutomationScheduleStepExecutedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

type InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesAction struct {
	ActionType                   InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsActionType `json:"action_type"`
	AmountAttempted              string                                                                       `json:"amount_attempted" api:"nullable"`
	Currency                     string                                                                       `json:"currency" api:"nullable"`
	FailureReason                string                                                                       `json:"failure_reason" api:"nullable"`
	Outcome                      InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsOutcome    `json:"outcome"`
	PaymentProvider              string                                                                       `json:"payment_provider" api:"nullable"`
	PaymentProviderTransactionID string                                                                       `json:"payment_provider_transaction_id" api:"nullable"`
	PaymentTransactionRecordID   string                                                                       `json:"payment_transaction_record_id" api:"nullable"`
	Recipient                    string                                                                       `json:"recipient" api:"nullable"`
	Sent                         bool                                                                         `json:"sent"`
	JSON                         invoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionJSON        `json:"-"`
	union                        InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsUnion
}

// invoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionJSON contains
// the JSON metadata for the struct
// [InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesAction]
type invoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionJSON struct {
	ActionType                   apijson.Field
	AmountAttempted              apijson.Field
	Currency                     apijson.Field
	FailureReason                apijson.Field
	Outcome                      apijson.Field
	PaymentProvider              apijson.Field
	PaymentProviderTransactionID apijson.Field
	PaymentTransactionRecordID   apijson.Field
	Recipient                    apijson.Field
	Sent                         apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r invoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionJSON) RawJSON() string {
	return r.raw
}

func (r *InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesAction) UnmarshalJSON(data []byte) (err error) {
	*r = InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesAction{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsSendEmailActionEntry],
// [InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntry].
func (r InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesAction) AsUnion() InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsUnion {
	return r.union
}

// Union satisfied by
// [InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsSendEmailActionEntry]
// or
// [InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntry].
type InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsUnion interface {
	implementsInvoiceAutomationScheduleStepExecutedWebhookEventPropertiesAction()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsUnion)(nil)).Elem(),
		"action_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsSendEmailActionEntry{}),
			DiscriminatorValue: "send_email",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntry{}),
			DiscriminatorValue: "retry_payment",
		},
	)
}

type InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsSendEmailActionEntry struct {
	Recipient  string                                                                                           `json:"recipient" api:"required,nullable"`
	Sent       bool                                                                                             `json:"sent" api:"required"`
	ActionType InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsSendEmailActionEntryActionType `json:"action_type"`
	JSON       invoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsSendEmailActionEntryJSON       `json:"-"`
}

// invoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsSendEmailActionEntryJSON
// contains the JSON metadata for the struct
// [InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsSendEmailActionEntry]
type invoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsSendEmailActionEntryJSON struct {
	Recipient   apijson.Field
	Sent        apijson.Field
	ActionType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsSendEmailActionEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsSendEmailActionEntryJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsSendEmailActionEntry) implementsInvoiceAutomationScheduleStepExecutedWebhookEventPropertiesAction() {
}

type InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsSendEmailActionEntryActionType string

const (
	InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsSendEmailActionEntryActionTypeSendEmail InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsSendEmailActionEntryActionType = "send_email"
)

func (r InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsSendEmailActionEntryActionType) IsKnown() bool {
	switch r {
	case InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsSendEmailActionEntryActionTypeSendEmail:
		return true
	}
	return false
}

type InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntry struct {
	AmountAttempted              string                                                                                              `json:"amount_attempted" api:"required,nullable"`
	Currency                     string                                                                                              `json:"currency" api:"required,nullable"`
	FailureReason                string                                                                                              `json:"failure_reason" api:"required,nullable"`
	Outcome                      InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryOutcome    `json:"outcome" api:"required"`
	PaymentProvider              string                                                                                              `json:"payment_provider" api:"required,nullable"`
	PaymentProviderTransactionID string                                                                                              `json:"payment_provider_transaction_id" api:"required,nullable"`
	PaymentTransactionRecordID   string                                                                                              `json:"payment_transaction_record_id" api:"required,nullable"`
	ActionType                   InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryActionType `json:"action_type"`
	JSON                         invoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryJSON       `json:"-"`
}

// invoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryJSON
// contains the JSON metadata for the struct
// [InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntry]
type invoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryJSON struct {
	AmountAttempted              apijson.Field
	Currency                     apijson.Field
	FailureReason                apijson.Field
	Outcome                      apijson.Field
	PaymentProvider              apijson.Field
	PaymentProviderTransactionID apijson.Field
	PaymentTransactionRecordID   apijson.Field
	ActionType                   apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntry) implementsInvoiceAutomationScheduleStepExecutedWebhookEventPropertiesAction() {
}

type InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryOutcome string

const (
	InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryOutcomeSucceeded InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryOutcome = "succeeded"
	InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryOutcomeFailed    InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryOutcome = "failed"
	InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryOutcomeSkipped   InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryOutcome = "skipped"
)

func (r InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryOutcome) IsKnown() bool {
	switch r {
	case InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryOutcomeSucceeded, InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryOutcomeFailed, InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryOutcomeSkipped:
		return true
	}
	return false
}

type InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryActionType string

const (
	InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryActionTypeRetryPayment InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryActionType = "retry_payment"
)

func (r InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryActionType) IsKnown() bool {
	switch r {
	case InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsRetryPaymentActionEntryActionTypeRetryPayment:
		return true
	}
	return false
}

type InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsActionType string

const (
	InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsActionTypeSendEmail    InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsActionType = "send_email"
	InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsActionTypeRetryPayment InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsActionType = "retry_payment"
)

func (r InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsActionType) IsKnown() bool {
	switch r {
	case InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsActionTypeSendEmail, InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsActionTypeRetryPayment:
		return true
	}
	return false
}

type InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsOutcome string

const (
	InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsOutcomeSucceeded InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsOutcome = "succeeded"
	InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsOutcomeFailed    InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsOutcome = "failed"
	InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsOutcomeSkipped   InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsOutcome = "skipped"
)

func (r InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsOutcome) IsKnown() bool {
	switch r {
	case InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsOutcomeSucceeded, InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsOutcomeFailed, InvoiceAutomationScheduleStepExecutedWebhookEventPropertiesActionsOutcomeSkipped:
		return true
	}
	return false
}

// The event this payload describes.
type InvoiceAutomationScheduleStepExecutedWebhookEventType string

const (
	InvoiceAutomationScheduleStepExecutedWebhookEventTypeInvoiceAutomationScheduleStepExecuted InvoiceAutomationScheduleStepExecutedWebhookEventType = "invoice.automation_schedule_step_executed"
)

func (r InvoiceAutomationScheduleStepExecutedWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceAutomationScheduleStepExecutedWebhookEventTypeInvoiceAutomationScheduleStepExecuted:
		return true
	}
	return false
}

// Issued when invoice cost data is exported.
type InvoiceCostDataExportedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                     `json:"created_at" api:"required" format:"date-time"`
	Invoice    string                                        `json:"invoice" api:"required"`
	Properties InvoiceCostDataExportedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceCostDataExportedWebhookEventType `json:"type" api:"required"`
	JSON invoiceCostDataExportedWebhookEventJSON `json:"-"`
}

// invoiceCostDataExportedWebhookEventJSON contains the JSON metadata for the
// struct [InvoiceCostDataExportedWebhookEvent]
type invoiceCostDataExportedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceCostDataExportedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceCostDataExportedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceCostDataExportedWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoiceCostDataExportedWebhookEventProperties struct {
	ExportedDate time.Time                                         `json:"exported_date" api:"required" format:"date-time"`
	S3Bucket     string                                            `json:"s3_bucket" api:"required"`
	S3Key        string                                            `json:"s3_key" api:"required"`
	JSON         invoiceCostDataExportedWebhookEventPropertiesJSON `json:"-"`
}

// invoiceCostDataExportedWebhookEventPropertiesJSON contains the JSON metadata for
// the struct [InvoiceCostDataExportedWebhookEventProperties]
type invoiceCostDataExportedWebhookEventPropertiesJSON struct {
	ExportedDate apijson.Field
	S3Bucket     apijson.Field
	S3Key        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InvoiceCostDataExportedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceCostDataExportedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoiceCostDataExportedWebhookEventType string

const (
	InvoiceCostDataExportedWebhookEventTypeInvoiceCostDataExported InvoiceCostDataExportedWebhookEventType = "invoice.cost_data_exported"
)

func (r InvoiceCostDataExportedWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceCostDataExportedWebhookEventTypeInvoiceCostDataExported:
		return true
	}
	return false
}

// Issued when a dunning schedule is created for an invoice.
type InvoiceDunningScheduleCreatedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice                                      `json:"invoice" api:"required"`
	Properties InvoiceDunningScheduleCreatedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceDunningScheduleCreatedWebhookEventType `json:"type" api:"required"`
	JSON invoiceDunningScheduleCreatedWebhookEventJSON `json:"-"`
}

// invoiceDunningScheduleCreatedWebhookEventJSON contains the JSON metadata for the
// struct [InvoiceDunningScheduleCreatedWebhookEvent]
type invoiceDunningScheduleCreatedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceDunningScheduleCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDunningScheduleCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceDunningScheduleCreatedWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoiceDunningScheduleCreatedWebhookEventProperties struct {
	DunningSchedule InvoiceDunningScheduleCreatedWebhookEventPropertiesDunningSchedule `json:"dunning_schedule" api:"required"`
	DunningSteps    []InvoiceDunningScheduleCreatedWebhookEventPropertiesDunningStep   `json:"dunning_steps" api:"required"`
	Source          string                                                             `json:"source" api:"required"`
	JSON            invoiceDunningScheduleCreatedWebhookEventPropertiesJSON            `json:"-"`
}

// invoiceDunningScheduleCreatedWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [InvoiceDunningScheduleCreatedWebhookEventProperties]
type invoiceDunningScheduleCreatedWebhookEventPropertiesJSON struct {
	DunningSchedule apijson.Field
	DunningSteps    apijson.Field
	Source          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InvoiceDunningScheduleCreatedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDunningScheduleCreatedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

type InvoiceDunningScheduleCreatedWebhookEventPropertiesDunningSchedule struct {
	CompletionTime time.Time                                                              `json:"completion_time" api:"required,nullable" format:"date-time"`
	CreatedAt      time.Time                                                              `json:"created_at" api:"required,nullable" format:"date-time"`
	InvoiceID      string                                                                 `json:"invoice_id" api:"required"`
	ModifiedAt     time.Time                                                              `json:"modified_at" api:"required,nullable" format:"date-time"`
	StartTime      time.Time                                                              `json:"start_time" api:"required,nullable" format:"date-time"`
	Status         string                                                                 `json:"status" api:"required"`
	JSON           invoiceDunningScheduleCreatedWebhookEventPropertiesDunningScheduleJSON `json:"-"`
}

// invoiceDunningScheduleCreatedWebhookEventPropertiesDunningScheduleJSON contains
// the JSON metadata for the struct
// [InvoiceDunningScheduleCreatedWebhookEventPropertiesDunningSchedule]
type invoiceDunningScheduleCreatedWebhookEventPropertiesDunningScheduleJSON struct {
	CompletionTime apijson.Field
	CreatedAt      apijson.Field
	InvoiceID      apijson.Field
	ModifiedAt     apijson.Field
	StartTime      apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InvoiceDunningScheduleCreatedWebhookEventPropertiesDunningSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDunningScheduleCreatedWebhookEventPropertiesDunningScheduleJSON) RawJSON() string {
	return r.raw
}

type InvoiceDunningScheduleCreatedWebhookEventPropertiesDunningStep struct {
	Actions             []string                                                           `json:"actions" api:"required"`
	CreatedAt           time.Time                                                          `json:"created_at" api:"required,nullable" format:"date-time"`
	ExecutionTime       time.Time                                                          `json:"execution_time" api:"required,nullable" format:"date-time"`
	ManuallyTriggeredAt time.Time                                                          `json:"manually_triggered_at" api:"required,nullable" format:"date-time"`
	ModifiedAt          time.Time                                                          `json:"modified_at" api:"required,nullable" format:"date-time"`
	Status              string                                                             `json:"status" api:"required"`
	StepNumber          int64                                                              `json:"step_number" api:"required,nullable"`
	Timestamp           time.Time                                                          `json:"timestamp" api:"required,nullable" format:"date-time"`
	JSON                invoiceDunningScheduleCreatedWebhookEventPropertiesDunningStepJSON `json:"-"`
}

// invoiceDunningScheduleCreatedWebhookEventPropertiesDunningStepJSON contains the
// JSON metadata for the struct
// [InvoiceDunningScheduleCreatedWebhookEventPropertiesDunningStep]
type invoiceDunningScheduleCreatedWebhookEventPropertiesDunningStepJSON struct {
	Actions             apijson.Field
	CreatedAt           apijson.Field
	ExecutionTime       apijson.Field
	ManuallyTriggeredAt apijson.Field
	ModifiedAt          apijson.Field
	Status              apijson.Field
	StepNumber          apijson.Field
	Timestamp           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InvoiceDunningScheduleCreatedWebhookEventPropertiesDunningStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDunningScheduleCreatedWebhookEventPropertiesDunningStepJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoiceDunningScheduleCreatedWebhookEventType string

const (
	InvoiceDunningScheduleCreatedWebhookEventTypeInvoiceDunningScheduleCreated InvoiceDunningScheduleCreatedWebhookEventType = "invoice.dunning_schedule_created"
)

func (r InvoiceDunningScheduleCreatedWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceDunningScheduleCreatedWebhookEventTypeInvoiceDunningScheduleCreated:
		return true
	}
	return false
}

// Issued when a dunning schedule ends.
type InvoiceDunningScheduleEndedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice                                    `json:"invoice" api:"required"`
	Properties InvoiceDunningScheduleEndedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceDunningScheduleEndedWebhookEventType `json:"type" api:"required"`
	JSON invoiceDunningScheduleEndedWebhookEventJSON `json:"-"`
}

// invoiceDunningScheduleEndedWebhookEventJSON contains the JSON metadata for the
// struct [InvoiceDunningScheduleEndedWebhookEvent]
type invoiceDunningScheduleEndedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceDunningScheduleEndedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDunningScheduleEndedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceDunningScheduleEndedWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoiceDunningScheduleEndedWebhookEventProperties struct {
	DunningSchedule InvoiceDunningScheduleEndedWebhookEventPropertiesDunningSchedule `json:"dunning_schedule" api:"required"`
	DunningSteps    []InvoiceDunningScheduleEndedWebhookEventPropertiesDunningStep   `json:"dunning_steps" api:"required"`
	JSON            invoiceDunningScheduleEndedWebhookEventPropertiesJSON            `json:"-"`
}

// invoiceDunningScheduleEndedWebhookEventPropertiesJSON contains the JSON metadata
// for the struct [InvoiceDunningScheduleEndedWebhookEventProperties]
type invoiceDunningScheduleEndedWebhookEventPropertiesJSON struct {
	DunningSchedule apijson.Field
	DunningSteps    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InvoiceDunningScheduleEndedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDunningScheduleEndedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

type InvoiceDunningScheduleEndedWebhookEventPropertiesDunningSchedule struct {
	CompletionTime time.Time                                                            `json:"completion_time" api:"required,nullable" format:"date-time"`
	CreatedAt      time.Time                                                            `json:"created_at" api:"required,nullable" format:"date-time"`
	InvoiceID      string                                                               `json:"invoice_id" api:"required"`
	ModifiedAt     time.Time                                                            `json:"modified_at" api:"required,nullable" format:"date-time"`
	StartTime      time.Time                                                            `json:"start_time" api:"required,nullable" format:"date-time"`
	Status         string                                                               `json:"status" api:"required"`
	JSON           invoiceDunningScheduleEndedWebhookEventPropertiesDunningScheduleJSON `json:"-"`
}

// invoiceDunningScheduleEndedWebhookEventPropertiesDunningScheduleJSON contains
// the JSON metadata for the struct
// [InvoiceDunningScheduleEndedWebhookEventPropertiesDunningSchedule]
type invoiceDunningScheduleEndedWebhookEventPropertiesDunningScheduleJSON struct {
	CompletionTime apijson.Field
	CreatedAt      apijson.Field
	InvoiceID      apijson.Field
	ModifiedAt     apijson.Field
	StartTime      apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InvoiceDunningScheduleEndedWebhookEventPropertiesDunningSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDunningScheduleEndedWebhookEventPropertiesDunningScheduleJSON) RawJSON() string {
	return r.raw
}

type InvoiceDunningScheduleEndedWebhookEventPropertiesDunningStep struct {
	Actions             []string                                                         `json:"actions" api:"required"`
	CreatedAt           time.Time                                                        `json:"created_at" api:"required,nullable" format:"date-time"`
	ExecutionTime       time.Time                                                        `json:"execution_time" api:"required,nullable" format:"date-time"`
	ManuallyTriggeredAt time.Time                                                        `json:"manually_triggered_at" api:"required,nullable" format:"date-time"`
	ModifiedAt          time.Time                                                        `json:"modified_at" api:"required,nullable" format:"date-time"`
	Status              string                                                           `json:"status" api:"required"`
	StepNumber          int64                                                            `json:"step_number" api:"required,nullable"`
	Timestamp           time.Time                                                        `json:"timestamp" api:"required,nullable" format:"date-time"`
	JSON                invoiceDunningScheduleEndedWebhookEventPropertiesDunningStepJSON `json:"-"`
}

// invoiceDunningScheduleEndedWebhookEventPropertiesDunningStepJSON contains the
// JSON metadata for the struct
// [InvoiceDunningScheduleEndedWebhookEventPropertiesDunningStep]
type invoiceDunningScheduleEndedWebhookEventPropertiesDunningStepJSON struct {
	Actions             apijson.Field
	CreatedAt           apijson.Field
	ExecutionTime       apijson.Field
	ManuallyTriggeredAt apijson.Field
	ModifiedAt          apijson.Field
	Status              apijson.Field
	StepNumber          apijson.Field
	Timestamp           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InvoiceDunningScheduleEndedWebhookEventPropertiesDunningStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDunningScheduleEndedWebhookEventPropertiesDunningStepJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoiceDunningScheduleEndedWebhookEventType string

const (
	InvoiceDunningScheduleEndedWebhookEventTypeInvoiceDunningScheduleEnded InvoiceDunningScheduleEndedWebhookEventType = "invoice.dunning_schedule_ended"
)

func (r InvoiceDunningScheduleEndedWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceDunningScheduleEndedWebhookEventTypeInvoiceDunningScheduleEnded:
		return true
	}
	return false
}

// Issued when a dunning schedule is reset.
type InvoiceDunningScheduleResetWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice                                    `json:"invoice" api:"required"`
	Properties InvoiceDunningScheduleResetWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceDunningScheduleResetWebhookEventType `json:"type" api:"required"`
	JSON invoiceDunningScheduleResetWebhookEventJSON `json:"-"`
}

// invoiceDunningScheduleResetWebhookEventJSON contains the JSON metadata for the
// struct [InvoiceDunningScheduleResetWebhookEvent]
type invoiceDunningScheduleResetWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceDunningScheduleResetWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDunningScheduleResetWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceDunningScheduleResetWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoiceDunningScheduleResetWebhookEventProperties struct {
	// The schedule shape `DunningScheduleV2.external_serialization()` produces, which
	// is narrower than the schedule's api resource.
	DunningSchedule InvoiceDunningScheduleResetWebhookEventPropertiesDunningSchedule `json:"dunning_schedule" api:"required"`
	// The schedule shape `DunningScheduleV2.external_serialization()` produces, which
	// is narrower than the schedule's api resource.
	PreviousSchedule InvoiceDunningScheduleResetWebhookEventPropertiesPreviousSchedule `json:"previous_schedule" api:"required"`
	JSON             invoiceDunningScheduleResetWebhookEventPropertiesJSON             `json:"-"`
}

// invoiceDunningScheduleResetWebhookEventPropertiesJSON contains the JSON metadata
// for the struct [InvoiceDunningScheduleResetWebhookEventProperties]
type invoiceDunningScheduleResetWebhookEventPropertiesJSON struct {
	DunningSchedule  apijson.Field
	PreviousSchedule apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvoiceDunningScheduleResetWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDunningScheduleResetWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The schedule shape `DunningScheduleV2.external_serialization()` produces, which
// is narrower than the schedule's api resource.
type InvoiceDunningScheduleResetWebhookEventPropertiesDunningSchedule struct {
	CompletionTime time.Time                                                            `json:"completion_time" api:"required,nullable" format:"date-time"`
	CreatedAt      time.Time                                                            `json:"created_at" api:"required,nullable" format:"date-time"`
	InvoiceID      string                                                               `json:"invoice_id" api:"required"`
	ModifiedAt     time.Time                                                            `json:"modified_at" api:"required,nullable" format:"date-time"`
	StartTime      time.Time                                                            `json:"start_time" api:"required,nullable" format:"date-time"`
	Status         string                                                               `json:"status" api:"required,nullable"`
	JSON           invoiceDunningScheduleResetWebhookEventPropertiesDunningScheduleJSON `json:"-"`
}

// invoiceDunningScheduleResetWebhookEventPropertiesDunningScheduleJSON contains
// the JSON metadata for the struct
// [InvoiceDunningScheduleResetWebhookEventPropertiesDunningSchedule]
type invoiceDunningScheduleResetWebhookEventPropertiesDunningScheduleJSON struct {
	CompletionTime apijson.Field
	CreatedAt      apijson.Field
	InvoiceID      apijson.Field
	ModifiedAt     apijson.Field
	StartTime      apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InvoiceDunningScheduleResetWebhookEventPropertiesDunningSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDunningScheduleResetWebhookEventPropertiesDunningScheduleJSON) RawJSON() string {
	return r.raw
}

// The schedule shape `DunningScheduleV2.external_serialization()` produces, which
// is narrower than the schedule's api resource.
type InvoiceDunningScheduleResetWebhookEventPropertiesPreviousSchedule struct {
	CompletionTime time.Time                                                             `json:"completion_time" api:"required,nullable" format:"date-time"`
	CreatedAt      time.Time                                                             `json:"created_at" api:"required,nullable" format:"date-time"`
	InvoiceID      string                                                                `json:"invoice_id" api:"required"`
	ModifiedAt     time.Time                                                             `json:"modified_at" api:"required,nullable" format:"date-time"`
	StartTime      time.Time                                                             `json:"start_time" api:"required,nullable" format:"date-time"`
	Status         string                                                                `json:"status" api:"required,nullable"`
	JSON           invoiceDunningScheduleResetWebhookEventPropertiesPreviousScheduleJSON `json:"-"`
}

// invoiceDunningScheduleResetWebhookEventPropertiesPreviousScheduleJSON contains
// the JSON metadata for the struct
// [InvoiceDunningScheduleResetWebhookEventPropertiesPreviousSchedule]
type invoiceDunningScheduleResetWebhookEventPropertiesPreviousScheduleJSON struct {
	CompletionTime apijson.Field
	CreatedAt      apijson.Field
	InvoiceID      apijson.Field
	ModifiedAt     apijson.Field
	StartTime      apijson.Field
	Status         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InvoiceDunningScheduleResetWebhookEventPropertiesPreviousSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDunningScheduleResetWebhookEventPropertiesPreviousScheduleJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoiceDunningScheduleResetWebhookEventType string

const (
	InvoiceDunningScheduleResetWebhookEventTypeInvoiceDunningScheduleReset InvoiceDunningScheduleResetWebhookEventType = "invoice.dunning_schedule_reset"
)

func (r InvoiceDunningScheduleResetWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceDunningScheduleResetWebhookEventTypeInvoiceDunningScheduleReset:
		return true
	}
	return false
}

// Issued when a dunning schedule step is executed.
type InvoiceDunningScheduleStepExecutedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice                                           `json:"invoice" api:"required"`
	Properties InvoiceDunningScheduleStepExecutedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceDunningScheduleStepExecutedWebhookEventType `json:"type" api:"required"`
	JSON invoiceDunningScheduleStepExecutedWebhookEventJSON `json:"-"`
}

// invoiceDunningScheduleStepExecutedWebhookEventJSON contains the JSON metadata
// for the struct [InvoiceDunningScheduleStepExecutedWebhookEvent]
type invoiceDunningScheduleStepExecutedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceDunningScheduleStepExecutedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDunningScheduleStepExecutedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceDunningScheduleStepExecutedWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoiceDunningScheduleStepExecutedWebhookEventProperties struct {
	DunningStep InvoiceDunningScheduleStepExecutedWebhookEventPropertiesDunningStep `json:"dunning_step" api:"required"`
	JSON        invoiceDunningScheduleStepExecutedWebhookEventPropertiesJSON        `json:"-"`
}

// invoiceDunningScheduleStepExecutedWebhookEventPropertiesJSON contains the JSON
// metadata for the struct
// [InvoiceDunningScheduleStepExecutedWebhookEventProperties]
type invoiceDunningScheduleStepExecutedWebhookEventPropertiesJSON struct {
	DunningStep apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceDunningScheduleStepExecutedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDunningScheduleStepExecutedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

type InvoiceDunningScheduleStepExecutedWebhookEventPropertiesDunningStep struct {
	Actions             []string                                                                `json:"actions" api:"required"`
	CreatedAt           time.Time                                                               `json:"created_at" api:"required,nullable" format:"date-time"`
	ExecutionTime       time.Time                                                               `json:"execution_time" api:"required,nullable" format:"date-time"`
	ManuallyTriggeredAt time.Time                                                               `json:"manually_triggered_at" api:"required,nullable" format:"date-time"`
	ModifiedAt          time.Time                                                               `json:"modified_at" api:"required,nullable" format:"date-time"`
	Status              string                                                                  `json:"status" api:"required"`
	StepNumber          int64                                                                   `json:"step_number" api:"required,nullable"`
	Timestamp           time.Time                                                               `json:"timestamp" api:"required,nullable" format:"date-time"`
	JSON                invoiceDunningScheduleStepExecutedWebhookEventPropertiesDunningStepJSON `json:"-"`
}

// invoiceDunningScheduleStepExecutedWebhookEventPropertiesDunningStepJSON contains
// the JSON metadata for the struct
// [InvoiceDunningScheduleStepExecutedWebhookEventPropertiesDunningStep]
type invoiceDunningScheduleStepExecutedWebhookEventPropertiesDunningStepJSON struct {
	Actions             apijson.Field
	CreatedAt           apijson.Field
	ExecutionTime       apijson.Field
	ManuallyTriggeredAt apijson.Field
	ModifiedAt          apijson.Field
	Status              apijson.Field
	StepNumber          apijson.Field
	Timestamp           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InvoiceDunningScheduleStepExecutedWebhookEventPropertiesDunningStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDunningScheduleStepExecutedWebhookEventPropertiesDunningStepJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoiceDunningScheduleStepExecutedWebhookEventType string

const (
	InvoiceDunningScheduleStepExecutedWebhookEventTypeInvoiceDunningScheduleStepExecuted InvoiceDunningScheduleStepExecutedWebhookEventType = "invoice.dunning_schedule_step_executed"
)

func (r InvoiceDunningScheduleStepExecutedWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceDunningScheduleStepExecutedWebhookEventTypeInvoiceDunningScheduleStepExecuted:
		return true
	}
	return false
}

// Issued when a draft invoice has been edited.
type InvoiceEditedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice                      `json:"invoice" api:"required"`
	Properties InvoiceEditedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceEditedWebhookEventType `json:"type" api:"required"`
	JSON invoiceEditedWebhookEventJSON `json:"-"`
}

// invoiceEditedWebhookEventJSON contains the JSON metadata for the struct
// [InvoiceEditedWebhookEvent]
type invoiceEditedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceEditedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceEditedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceEditedWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoiceEditedWebhookEventProperties struct {
	// `due_date` is an ISO-8601 string rather than a datetime: the untyped message
	// called `.isoformat()` on it, so it keeps microseconds where the webhook JSON
	// provider would have truncated them.
	PreviousAttributes InvoiceEditedWebhookEventPropertiesPreviousAttributes `json:"previous_attributes" api:"required"`
	JSON               invoiceEditedWebhookEventPropertiesJSON               `json:"-"`
}

// invoiceEditedWebhookEventPropertiesJSON contains the JSON metadata for the
// struct [InvoiceEditedWebhookEventProperties]
type invoiceEditedWebhookEventPropertiesJSON struct {
	PreviousAttributes apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceEditedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceEditedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// `due_date` is an ISO-8601 string rather than a datetime: the untyped message
// called `.isoformat()` on it, so it keeps microseconds where the webhook JSON
// provider would have truncated them.
type InvoiceEditedWebhookEventPropertiesPreviousAttributes struct {
	AmountDue string                                                          `json:"amount_due" api:"nullable"`
	Discounts []shared.InvoiceLevelDiscount                                   `json:"discounts" api:"nullable"`
	DueDate   string                                                          `json:"due_date" api:"nullable"`
	LineItems []InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItem `json:"line_items" api:"nullable"`
	Maximum   shared.Maximum                                                  `json:"maximum" api:"nullable"`
	Metadata  map[string]string                                               `json:"metadata" api:"nullable"`
	Minimum   shared.Minimum                                                  `json:"minimum" api:"nullable"`
	NetTerms  int64                                                           `json:"net_terms" api:"nullable"`
	Subtotal  string                                                          `json:"subtotal" api:"nullable"`
	Total     string                                                          `json:"total" api:"nullable"`
	JSON      invoiceEditedWebhookEventPropertiesPreviousAttributesJSON       `json:"-"`
}

// invoiceEditedWebhookEventPropertiesPreviousAttributesJSON contains the JSON
// metadata for the struct [InvoiceEditedWebhookEventPropertiesPreviousAttributes]
type invoiceEditedWebhookEventPropertiesPreviousAttributesJSON struct {
	AmountDue   apijson.Field
	Discounts   apijson.Field
	DueDate     apijson.Field
	LineItems   apijson.Field
	Maximum     apijson.Field
	Metadata    apijson.Field
	Minimum     apijson.Field
	NetTerms    apijson.Field
	Subtotal    apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceEditedWebhookEventPropertiesPreviousAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceEditedWebhookEventPropertiesPreviousAttributesJSON) RawJSON() string {
	return r.raw
}

type InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItem struct {
	// A unique ID for this line item.
	ID string `json:"id" api:"required"`
	// The line amount after any adjustments and before overage conversion, credits and
	// partial invoicing.
	AdjustedSubtotal string `json:"adjusted_subtotal" api:"required"`
	// All adjustments applied to the line item in the order they were applied based on
	// invoice calculations (ie. usage discounts -> amount discounts -> percentage
	// discounts -> minimums -> maximums).
	Adjustments []InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustment `json:"adjustments" api:"required"`
	// The final amount for a line item after all adjustments and pre paid credits have
	// been applied.
	Amount string `json:"amount" api:"required"`
	// The number of prepaid credits applied.
	CreditsApplied string `json:"credits_applied" api:"required"`
	// The end date of the range of time applied for this line item's price.
	EndDate time.Time `json:"end_date" api:"required" format:"date-time"`
	// An additional filter that was used to calculate the usage for this line item.
	Filter string `json:"filter" api:"required,nullable"`
	// [DEPRECATED] For configured prices that are split by a grouping key, this will
	// be populated with the key and a value. The `amount` and `subtotal` will be the
	// values for this particular grouping.
	Grouping string `json:"grouping" api:"required,nullable"`
	// The name of the price associated with this line item.
	Name string `json:"name" api:"required"`
	// Any amount applied from a partial invoice
	PartiallyInvoicedAmount string `json:"partially_invoiced_amount" api:"required"`
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
	Price shared.Price `json:"price" api:"required"`
	// Either the fixed fee quantity or the usage during the service period.
	Quantity float64 `json:"quantity" api:"required"`
	// The start date of the range of time applied for this line item's price.
	StartDate time.Time `json:"start_date" api:"required" format:"date-time"`
	// For complex pricing structures, the line item can be broken down further in
	// `sub_line_items`.
	SubLineItems []InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItem `json:"sub_line_items" api:"required"`
	// The line amount before any adjustments.
	Subtotal string `json:"subtotal" api:"required"`
	// An array of tax rates and their incurred tax amounts. Empty if no tax
	// integration is configured.
	TaxAmounts []shared.TaxAmount `json:"tax_amounts" api:"required"`
	// A list of customer ids that were used to calculate the usage for this line item.
	UsageCustomerIDs []string                                                          `json:"usage_customer_ids" api:"required,nullable"`
	JSON             invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemJSON `json:"-"`
}

// invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemJSON contains the
// JSON metadata for the struct
// [InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItem]
type invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemJSON struct {
	ID                      apijson.Field
	AdjustedSubtotal        apijson.Field
	Adjustments             apijson.Field
	Amount                  apijson.Field
	CreditsApplied          apijson.Field
	EndDate                 apijson.Field
	Filter                  apijson.Field
	Grouping                apijson.Field
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

func (r *InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemJSON) RawJSON() string {
	return r.raw
}

type InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustment struct {
	ID             string                                                                                  `json:"id" api:"required"`
	AdjustmentType InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentType `json:"adjustment_type" api:"required"`
	// The value applied by an adjustment.
	Amount string `json:"amount" api:"required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids" api:"required"`
	// This field can have the runtime type of
	// [[]shared.MonetaryUsageDiscountAdjustmentFilter],
	// [[]shared.MonetaryAmountDiscountAdjustmentFilter],
	// [[]shared.MonetaryPercentageDiscountAdjustmentFilter],
	// [[]InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFilter],
	// [[]shared.MonetaryMinimumAdjustmentFilter],
	// [[]shared.MonetaryMaximumAdjustmentFilter].
	Filters interface{} `json:"filters" api:"required"`
	// True for adjustments that apply to an entire invoice, false for adjustments that
	// apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level" api:"required"`
	// The reason for the adjustment.
	Reason string `json:"reason" api:"required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string `json:"replaces_adjustment_id" api:"required,nullable"`
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
	// This field can have the runtime type of
	// [[]InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentTier].
	Tiers interface{} `json:"tiers"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                                                      `json:"usage_discount"`
	JSON          invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentJSON `json:"-"`
	union         InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsUnion
}

// invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentJSON
// contains the JSON metadata for the struct
// [InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustment]
type invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	Amount               apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	AmountDiscount       apijson.Field
	ItemID               apijson.Field
	MaximumAmount        apijson.Field
	MinimumAmount        apijson.Field
	PercentageDiscount   apijson.Field
	Tiers                apijson.Field
	UsageDiscount        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.MonetaryUsageDiscountAdjustment],
// [shared.MonetaryAmountDiscountAdjustment],
// [shared.MonetaryPercentageDiscountAdjustment],
// [InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustment],
// [shared.MonetaryMinimumAdjustment], [shared.MonetaryMaximumAdjustment].
func (r InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustment) AsUnion() InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsUnion {
	return r.union
}

// Union satisfied by [shared.MonetaryUsageDiscountAdjustment],
// [shared.MonetaryAmountDiscountAdjustment],
// [shared.MonetaryPercentageDiscountAdjustment],
// [InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustment],
// [shared.MonetaryMinimumAdjustment] or [shared.MonetaryMaximumAdjustment].
type InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsUnion interface {
	ImplementsInvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.MonetaryUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.MonetaryAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.MonetaryPercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustment{}),
			DiscriminatorValue: "tiered_percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.MonetaryMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.MonetaryMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustment struct {
	ID             string                                                                                                                            `json:"id" api:"required"`
	AdjustmentType InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type" api:"required"`
	// The value applied by an adjustment.
	Amount string `json:"amount" api:"required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids" api:"required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFilter `json:"filters" api:"required"`
	// True for adjustments that apply to an entire invoice, false for adjustments that
	// apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level" api:"required"`
	// The reason for the adjustment.
	Reason string `json:"reason" api:"required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string `json:"replaces_adjustment_id" api:"required,nullable"`
	// The ordered, contiguous bands of cumulative eligible spend, each discounted at
	// its own percentage (progressive fill-a-tier), applied to the prices this
	// adjustment covers in a given billing period.
	Tiers []InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentTier `json:"tiers" api:"required"`
	JSON  invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentJSON   `json:"-"`
}

// invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustment]
type invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	Amount               apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	Tiers                apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustment) ImplementsInvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustment() {
}

type InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentAdjustmentType string

const (
	InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentAdjustmentTypeTieredPercentageDiscount InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentAdjustmentType = "tiered_percentage_discount"
)

func (r InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

type InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                                                                                      `json:"values" api:"required"`
	JSON   invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFilterJSON `json:"-"`
}

// invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFilter]
type invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersField string

const (
	InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersFieldPriceID       InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersField = "price_id"
	InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersFieldItemID        InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersField = "item_id"
	InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersFieldPriceType     InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersField = "price_type"
	InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersFieldCurrency      InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersField = "currency"
	InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersFieldPricingUnitID InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersFieldPriceID, InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersFieldItemID, InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersFieldPriceType, InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersFieldCurrency, InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersOperator string

const (
	InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersOperatorIncludes InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersOperator = "includes"
	InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersOperatorExcludes InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersOperatorIncludes, InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

// One band of a tiered percentage discount. Bounds are denominated in the
// discount's currency. `lower_bound` is the exclusive start of the band and
// `upper_bound` is the inclusive end; `upper_bound` is null only for the
// open-ended final tier.
type InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentTier struct {
	// Exclusive lower bound of cumulative spend for this tier.
	LowerBound float64 `json:"lower_bound" api:"required"`
	// The percentage (between 0 and 1) discounted from spend that falls within this
	// tier.
	Percentage float64 `json:"percentage" api:"required"`
	// Inclusive upper bound of cumulative spend for this tier; null for the final
	// open-ended tier.
	UpperBound float64                                                                                                                     `json:"upper_bound" api:"nullable"`
	JSON       invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentTierJSON `json:"-"`
}

// invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentTierJSON
// contains the JSON metadata for the struct
// [InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentTier]
type invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentTierJSON struct {
	LowerBound  apijson.Field
	Percentage  apijson.Field
	UpperBound  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsMonetaryTieredPercentageDiscountAdjustmentTierJSON) RawJSON() string {
	return r.raw
}

type InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentType string

const (
	InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentTypeUsageDiscount            InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentType = "usage_discount"
	InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentTypeAmountDiscount           InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentType = "amount_discount"
	InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentTypePercentageDiscount       InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentType = "percentage_discount"
	InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentTypeTieredPercentageDiscount InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentType = "tiered_percentage_discount"
	InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentTypeMinimum                  InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentType = "minimum"
	InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentTypeMaximum                  InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentType = "maximum"
)

func (r InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentTypeUsageDiscount, InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentTypeAmountDiscount, InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentTypePercentageDiscount, InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentTypeTieredPercentageDiscount, InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentTypeMinimum, InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsAdjustmentsAdjustmentTypeMaximum:
		return true
	}
	return false
}

type InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItem struct {
	// The total amount for this sub line item.
	Amount       string                                                                         `json:"amount" api:"required"`
	Grouping     shared.SubLineItemGrouping                                                     `json:"grouping" api:"required,nullable"`
	Name         string                                                                         `json:"name" api:"required"`
	Quantity     float64                                                                        `json:"quantity" api:"required"`
	Type         InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemsType `json:"type" api:"required"`
	MatrixConfig shared.SubLineItemMatrixConfig                                                 `json:"matrix_config"`
	// The scaled quantity for this line item for specific pricing structures
	ScaledQuantity float64 `json:"scaled_quantity" api:"nullable"`
	// This field can have the runtime type of [shared.TierSubLineItemTierConfig].
	TierConfig interface{}                                                                   `json:"tier_config"`
	JSON       invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemJSON `json:"-"`
	union      InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemsUnion
}

// invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemJSON
// contains the JSON metadata for the struct
// [InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItem]
type invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemJSON struct {
	Amount         apijson.Field
	Grouping       apijson.Field
	Name           apijson.Field
	Quantity       apijson.Field
	Type           apijson.Field
	MatrixConfig   apijson.Field
	ScaledQuantity apijson.Field
	TierConfig     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r invoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemJSON) RawJSON() string {
	return r.raw
}

func (r *InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItem) UnmarshalJSON(data []byte) (err error) {
	*r = InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [shared.MatrixSubLineItem],
// [shared.TierSubLineItem], [shared.OtherSubLineItem].
func (r InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItem) AsUnion() InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemsUnion {
	return r.union
}

// Union satisfied by [shared.MatrixSubLineItem], [shared.TierSubLineItem] or
// [shared.OtherSubLineItem].
type InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemsUnion interface {
	ImplementsInvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.MatrixSubLineItem{}),
			DiscriminatorValue: "matrix",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.TierSubLineItem{}),
			DiscriminatorValue: "tier",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.OtherSubLineItem{}),
			DiscriminatorValue: "'null'",
		},
	)
}

type InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemsType string

const (
	InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemsTypeMatrix InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemsType = "matrix"
	InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemsTypeTier   InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemsType = "tier"
	InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemsTypeNull   InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemsType = "'null'"
)

func (r InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemsType) IsKnown() bool {
	switch r {
	case InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemsTypeMatrix, InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemsTypeTier, InvoiceEditedWebhookEventPropertiesPreviousAttributesLineItemsSubLineItemsTypeNull:
		return true
	}
	return false
}

// The event this payload describes.
type InvoiceEditedWebhookEventType string

const (
	InvoiceEditedWebhookEventTypeInvoiceEdited InvoiceEditedWebhookEventType = "invoice.edited"
)

func (r InvoiceEditedWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceEditedWebhookEventTypeInvoiceEdited:
		return true
	}
	return false
}

// Issued when an invoice's invoice date has elapsed.
type InvoiceInvoiceDateElapsedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                       `json:"created_at" api:"required" format:"date-time"`
	Invoice    InvoiceInvoiceDateElapsedWebhookEventInvoice    `json:"invoice" api:"required"`
	Properties InvoiceInvoiceDateElapsedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceInvoiceDateElapsedWebhookEventType `json:"type" api:"required"`
	JSON invoiceInvoiceDateElapsedWebhookEventJSON `json:"-"`
}

// invoiceInvoiceDateElapsedWebhookEventJSON contains the JSON metadata for the
// struct [InvoiceInvoiceDateElapsedWebhookEvent]
type invoiceInvoiceDateElapsedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceInvoiceDateElapsedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceInvoiceDateElapsedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceInvoiceDateElapsedWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoiceInvoiceDateElapsedWebhookEventInvoice struct {
	ID            string                                             `json:"id" api:"required"`
	Customer      shared.CustomerMinified                            `json:"customer" api:"required"`
	InvoiceNumber string                                             `json:"invoice_number" api:"required"`
	Status        InvoiceInvoiceDateElapsedWebhookEventInvoiceStatus `json:"status" api:"required"`
	Subscription  shared.SubscriptionMinified                        `json:"subscription" api:"required,nullable"`
	JSON          invoiceInvoiceDateElapsedWebhookEventInvoiceJSON   `json:"-"`
}

// invoiceInvoiceDateElapsedWebhookEventInvoiceJSON contains the JSON metadata for
// the struct [InvoiceInvoiceDateElapsedWebhookEventInvoice]
type invoiceInvoiceDateElapsedWebhookEventInvoiceJSON struct {
	ID            apijson.Field
	Customer      apijson.Field
	InvoiceNumber apijson.Field
	Status        apijson.Field
	Subscription  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InvoiceInvoiceDateElapsedWebhookEventInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceInvoiceDateElapsedWebhookEventInvoiceJSON) RawJSON() string {
	return r.raw
}

type InvoiceInvoiceDateElapsedWebhookEventInvoiceStatus string

const (
	InvoiceInvoiceDateElapsedWebhookEventInvoiceStatusIssued InvoiceInvoiceDateElapsedWebhookEventInvoiceStatus = "issued"
	InvoiceInvoiceDateElapsedWebhookEventInvoiceStatusPaid   InvoiceInvoiceDateElapsedWebhookEventInvoiceStatus = "paid"
	InvoiceInvoiceDateElapsedWebhookEventInvoiceStatusSynced InvoiceInvoiceDateElapsedWebhookEventInvoiceStatus = "synced"
	InvoiceInvoiceDateElapsedWebhookEventInvoiceStatusVoid   InvoiceInvoiceDateElapsedWebhookEventInvoiceStatus = "void"
	InvoiceInvoiceDateElapsedWebhookEventInvoiceStatusDraft  InvoiceInvoiceDateElapsedWebhookEventInvoiceStatus = "draft"
)

func (r InvoiceInvoiceDateElapsedWebhookEventInvoiceStatus) IsKnown() bool {
	switch r {
	case InvoiceInvoiceDateElapsedWebhookEventInvoiceStatusIssued, InvoiceInvoiceDateElapsedWebhookEventInvoiceStatusPaid, InvoiceInvoiceDateElapsedWebhookEventInvoiceStatusSynced, InvoiceInvoiceDateElapsedWebhookEventInvoiceStatusVoid, InvoiceInvoiceDateElapsedWebhookEventInvoiceStatusDraft:
		return true
	}
	return false
}

type InvoiceInvoiceDateElapsedWebhookEventProperties struct {
	InvoiceDate time.Time                                           `json:"invoice_date" api:"required" format:"date-time"`
	JSON        invoiceInvoiceDateElapsedWebhookEventPropertiesJSON `json:"-"`
}

// invoiceInvoiceDateElapsedWebhookEventPropertiesJSON contains the JSON metadata
// for the struct [InvoiceInvoiceDateElapsedWebhookEventProperties]
type invoiceInvoiceDateElapsedWebhookEventPropertiesJSON struct {
	InvoiceDate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceInvoiceDateElapsedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceInvoiceDateElapsedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoiceInvoiceDateElapsedWebhookEventType string

const (
	InvoiceInvoiceDateElapsedWebhookEventTypeInvoiceInvoiceDateElapsed InvoiceInvoiceDateElapsedWebhookEventType = "invoice.invoice_date_elapsed"
)

func (r InvoiceInvoiceDateElapsedWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceInvoiceDateElapsedWebhookEventTypeInvoiceInvoiceDateElapsed:
		return true
	}
	return false
}

// Issued when an invoice issue attempt fails.
type InvoiceIssueFailedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice                           `json:"invoice" api:"required"`
	Properties InvoiceIssueFailedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceIssueFailedWebhookEventType `json:"type" api:"required"`
	JSON invoiceIssueFailedWebhookEventJSON `json:"-"`
}

// invoiceIssueFailedWebhookEventJSON contains the JSON metadata for the struct
// [InvoiceIssueFailedWebhookEvent]
type invoiceIssueFailedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceIssueFailedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceIssueFailedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceIssueFailedWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoiceIssueFailedWebhookEventProperties struct {
	Reason string                                       `json:"reason" api:"required,nullable"`
	JSON   invoiceIssueFailedWebhookEventPropertiesJSON `json:"-"`
}

// invoiceIssueFailedWebhookEventPropertiesJSON contains the JSON metadata for the
// struct [InvoiceIssueFailedWebhookEventProperties]
type invoiceIssueFailedWebhookEventPropertiesJSON struct {
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceIssueFailedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceIssueFailedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoiceIssueFailedWebhookEventType string

const (
	InvoiceIssueFailedWebhookEventTypeInvoiceIssueFailed InvoiceIssueFailedWebhookEventType = "invoice.issue_failed"
)

func (r InvoiceIssueFailedWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceIssueFailedWebhookEventTypeInvoiceIssueFailed:
		return true
	}
	return false
}

// Issued when an invoice transitions to the "issued" state.
type InvoiceIssuedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice                      `json:"invoice" api:"required"`
	Properties InvoiceIssuedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceIssuedWebhookEventType `json:"type" api:"required"`
	JSON invoiceIssuedWebhookEventJSON `json:"-"`
}

// invoiceIssuedWebhookEventJSON contains the JSON metadata for the struct
// [InvoiceIssuedWebhookEvent]
type invoiceIssuedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceIssuedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceIssuedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceIssuedWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoiceIssuedWebhookEventProperties struct {
	AutomaticallyMarkedAsPaid bool                                    `json:"automatically_marked_as_paid" api:"required"`
	JSON                      invoiceIssuedWebhookEventPropertiesJSON `json:"-"`
}

// invoiceIssuedWebhookEventPropertiesJSON contains the JSON metadata for the
// struct [InvoiceIssuedWebhookEventProperties]
type invoiceIssuedWebhookEventPropertiesJSON struct {
	AutomaticallyMarkedAsPaid apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *InvoiceIssuedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceIssuedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoiceIssuedWebhookEventType string

const (
	InvoiceIssuedWebhookEventTypeInvoiceIssued InvoiceIssuedWebhookEventType = "invoice.issued"
)

func (r InvoiceIssuedWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceIssuedWebhookEventTypeInvoiceIssued:
		return true
	}
	return false
}

// A lightweight variant of invoice.issued for accounts configured to receive a
// summarized invoice payload.
type InvoiceIssuedSummaryWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// #InvoiceApiResourceWithoutLineItems
	Invoice    InvoiceIssuedSummaryWebhookEventInvoice    `json:"invoice" api:"required"`
	Properties InvoiceIssuedSummaryWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceIssuedSummaryWebhookEventType `json:"type" api:"required"`
	JSON invoiceIssuedSummaryWebhookEventJSON `json:"-"`
}

// invoiceIssuedSummaryWebhookEventJSON contains the JSON metadata for the struct
// [InvoiceIssuedSummaryWebhookEvent]
type invoiceIssuedSummaryWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceIssuedSummaryWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceIssuedSummaryWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceIssuedSummaryWebhookEvent) implementsUnwrapWebhookEvent() {}

// #InvoiceApiResourceWithoutLineItems
type InvoiceIssuedSummaryWebhookEventInvoice struct {
	ID string `json:"id" api:"required"`
	// This is the final amount required to be charged to the customer and reflects the
	// application of the customer balance to the `total` of the invoice.
	AmountDue      string                                                `json:"amount_due" api:"required"`
	AutoCollection InvoiceIssuedSummaryWebhookEventInvoiceAutoCollection `json:"auto_collection" api:"required"`
	BillingAddress shared.Address                                        `json:"billing_address" api:"required,nullable"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A list of credit notes associated with the invoice
	CreditNotes []InvoiceIssuedSummaryWebhookEventInvoiceCreditNote `json:"credit_notes" api:"required"`
	// An ISO 4217 currency string or `credits`
	Currency                    string                                                              `json:"currency" api:"required"`
	Customer                    shared.CustomerMinified                                             `json:"customer" api:"required"`
	CustomerBalanceTransactions []InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransaction `json:"customer_balance_transactions" api:"required"`
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
	// | Faroe Islands          | `fo_vat`     | Faroe Islands VAT Number                                                                                |
	// | Finland                | `eu_vat`     | European VAT Number                                                                                     |
	// | France                 | `eu_vat`     | European VAT Number                                                                                     |
	// | Georgia                | `ge_vat`     | Georgian VAT                                                                                            |
	// | Germany                | `de_stn`     | German Tax Number (Steuernummer)                                                                        |
	// | Germany                | `eu_vat`     | European VAT Number                                                                                     |
	// | Gibraltar              | `gi_tin`     | Gibraltar Tax Identification Number                                                                     |
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
	// | Italy                  | `it_cf`      | Italian Codice Fiscale Number                                                                           |
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
	// | Paraguay               | `py_ruc`     | Paraguayan RUC Number                                                                                   |
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
	// | Sri Lanka              | `lk_vat`     | Sri Lanka VAT Number                                                                                    |
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
	CustomerTaxID shared.CustomerTaxID `json:"customer_tax_id" api:"required,nullable"`
	// When the invoice payment is due. The due date is null if the invoice is not yet
	// finalized.
	DueDate time.Time `json:"due_date" api:"required,nullable" format:"date-time"`
	// If the invoice has a status of `draft`, this will be the time that the invoice
	// will be eligible to be issued, otherwise it will be `null`. If `auto-issue` is
	// true, the invoice will automatically begin issuing at this time.
	EligibleToIssueAt time.Time `json:"eligible_to_issue_at" api:"required,nullable" format:"date-time"`
	// A URL for the customer-facing invoice portal. This URL expires 60 days after the
	// link is generated, or 30 days after the invoice's due date — whichever is later.
	HostedInvoiceURL string `json:"hosted_invoice_url" api:"required,nullable"`
	// The scheduled date of the invoice
	InvoiceDate time.Time `json:"invoice_date" api:"required" format:"date-time"`
	// Automatically generated invoice number to help track and reconcile invoices.
	// Invoice numbers have a prefix such as `RFOBWG`. These can be sequential per
	// account or customer.
	InvoiceNumber string `json:"invoice_number" api:"required"`
	// The link to download the PDF representation of the `Invoice`.
	InvoicePdf    string                                               `json:"invoice_pdf" api:"required,nullable"`
	InvoiceSource InvoiceIssuedSummaryWebhookEventInvoiceInvoiceSource `json:"invoice_source" api:"required"`
	// If the invoice failed to issue, this will be the last time it failed to issue
	// (even if it is now in a different state.)
	IssueFailedAt time.Time `json:"issue_failed_at" api:"required,nullable" format:"date-time"`
	// If the invoice has been issued, this will be the time it transitioned to
	// `issued` (even if it is now in a different state.)
	IssuedAt time.Time `json:"issued_at" api:"required,nullable" format:"date-time"`
	// Free-form text which is available on the invoice PDF and the Orb invoice portal.
	Memo string `json:"memo" api:"required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata" api:"required"`
	// If the invoice has a status of `paid`, this gives a timestamp when the invoice
	// was paid.
	PaidAt time.Time `json:"paid_at" api:"required,nullable" format:"date-time"`
	// A list of payment attempts associated with the invoice
	PaymentAttempts []InvoiceIssuedSummaryWebhookEventInvoicePaymentAttempt `json:"payment_attempts" api:"required"`
	// If payment was attempted on this invoice but failed, this will be the time of
	// the most recent attempt.
	PaymentFailedAt time.Time `json:"payment_failed_at" api:"required,nullable" format:"date-time"`
	// If payment was attempted on this invoice, this will be the start time of the
	// most recent attempt. This field is especially useful for delayed-notification
	// payment mechanisms (like bank transfers), where payment can take 3 days or more.
	PaymentStartedAt time.Time `json:"payment_started_at" api:"required,nullable" format:"date-time"`
	// If the invoice is in draft, this timestamp will reflect when the invoice is
	// scheduled to be issued.
	ScheduledIssueAt time.Time                                     `json:"scheduled_issue_at" api:"required,nullable" format:"date-time"`
	ShippingAddress  shared.Address                                `json:"shipping_address" api:"required,nullable"`
	Status           InvoiceIssuedSummaryWebhookEventInvoiceStatus `json:"status" api:"required"`
	Subscription     shared.SubscriptionMinified                   `json:"subscription" api:"required,nullable"`
	// If the invoice failed to sync, this will be the last time an external invoicing
	// provider sync was attempted. This field will always be `null` for invoices using
	// Orb Invoicing.
	SyncFailedAt time.Time `json:"sync_failed_at" api:"required,nullable" format:"date-time"`
	// The total after any minimums and discounts have been applied.
	Total string `json:"total" api:"required"`
	// If the invoice has a status of `void`, this gives a timestamp when the invoice
	// was voided.
	VoidedAt time.Time `json:"voided_at" api:"required,nullable" format:"date-time"`
	// This is true if the invoice will be automatically issued in the future, and
	// false otherwise.
	WillAutoIssue bool                                        `json:"will_auto_issue" api:"required"`
	JSON          invoiceIssuedSummaryWebhookEventInvoiceJSON `json:"-"`
}

// invoiceIssuedSummaryWebhookEventInvoiceJSON contains the JSON metadata for the
// struct [InvoiceIssuedSummaryWebhookEventInvoice]
type invoiceIssuedSummaryWebhookEventInvoiceJSON struct {
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
	DueDate                     apijson.Field
	EligibleToIssueAt           apijson.Field
	HostedInvoiceURL            apijson.Field
	InvoiceDate                 apijson.Field
	InvoiceNumber               apijson.Field
	InvoicePdf                  apijson.Field
	InvoiceSource               apijson.Field
	IssueFailedAt               apijson.Field
	IssuedAt                    apijson.Field
	Memo                        apijson.Field
	Metadata                    apijson.Field
	PaidAt                      apijson.Field
	PaymentAttempts             apijson.Field
	PaymentFailedAt             apijson.Field
	PaymentStartedAt            apijson.Field
	ScheduledIssueAt            apijson.Field
	ShippingAddress             apijson.Field
	Status                      apijson.Field
	Subscription                apijson.Field
	SyncFailedAt                apijson.Field
	Total                       apijson.Field
	VoidedAt                    apijson.Field
	WillAutoIssue               apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *InvoiceIssuedSummaryWebhookEventInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceIssuedSummaryWebhookEventInvoiceJSON) RawJSON() string {
	return r.raw
}

type InvoiceIssuedSummaryWebhookEventInvoiceAutoCollection struct {
	// True only if auto-collection is enabled for this invoice.
	Enabled bool `json:"enabled" api:"required,nullable"`
	// If the invoice is scheduled for auto-collection, this field will reflect when
	// the next attempt will occur. If dunning has been exhausted, or auto-collection
	// is not enabled for this invoice, this field will be `null`.
	NextAttemptAt time.Time `json:"next_attempt_at" api:"required,nullable" format:"date-time"`
	// Number of auto-collection payment attempts.
	NumAttempts int64 `json:"num_attempts" api:"required,nullable"`
	// If Orb has ever attempted payment auto-collection for this invoice, this field
	// will reflect when that attempt occurred. In conjunction with `next_attempt_at`,
	// this can be used to tell whether the invoice is currently in dunning (that is,
	// `previously_attempted_at` is non-null, and `next_attempt_at` is non-null), or if
	// dunning has been exhausted (`previously_attempted_at` is non-null, but
	// `next_attempt_at` is null).
	PreviouslyAttemptedAt time.Time                                                 `json:"previously_attempted_at" api:"required,nullable" format:"date-time"`
	JSON                  invoiceIssuedSummaryWebhookEventInvoiceAutoCollectionJSON `json:"-"`
}

// invoiceIssuedSummaryWebhookEventInvoiceAutoCollectionJSON contains the JSON
// metadata for the struct [InvoiceIssuedSummaryWebhookEventInvoiceAutoCollection]
type invoiceIssuedSummaryWebhookEventInvoiceAutoCollectionJSON struct {
	Enabled               apijson.Field
	NextAttemptAt         apijson.Field
	NumAttempts           apijson.Field
	PreviouslyAttemptedAt apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InvoiceIssuedSummaryWebhookEventInvoiceAutoCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceIssuedSummaryWebhookEventInvoiceAutoCollectionJSON) RawJSON() string {
	return r.raw
}

type InvoiceIssuedSummaryWebhookEventInvoiceCreditNote struct {
	ID               string `json:"id" api:"required"`
	CreditNoteNumber string `json:"credit_note_number" api:"required"`
	// An optional memo supplied on the credit note.
	Memo   string `json:"memo" api:"required,nullable"`
	Reason string `json:"reason" api:"required"`
	Total  string `json:"total" api:"required"`
	Type   string `json:"type" api:"required"`
	// If the credit note has a status of `void`, this gives a timestamp when the
	// credit note was voided.
	VoidedAt time.Time                                             `json:"voided_at" api:"required,nullable" format:"date-time"`
	JSON     invoiceIssuedSummaryWebhookEventInvoiceCreditNoteJSON `json:"-"`
}

// invoiceIssuedSummaryWebhookEventInvoiceCreditNoteJSON contains the JSON metadata
// for the struct [InvoiceIssuedSummaryWebhookEventInvoiceCreditNote]
type invoiceIssuedSummaryWebhookEventInvoiceCreditNoteJSON struct {
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

func (r *InvoiceIssuedSummaryWebhookEventInvoiceCreditNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceIssuedSummaryWebhookEventInvoiceCreditNoteJSON) RawJSON() string {
	return r.raw
}

type InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransaction struct {
	// A unique id for this transaction.
	ID     string                                                                   `json:"id" api:"required"`
	Action InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsAction `json:"action" api:"required"`
	// The value of the amount changed in the transaction.
	Amount string `json:"amount" api:"required"`
	// The creation time of this transaction.
	CreatedAt  time.Time             `json:"created_at" api:"required" format:"date-time"`
	CreditNote shared.CreditNoteTiny `json:"credit_note" api:"required,nullable"`
	// An optional description provided for manual customer balance adjustments.
	Description string `json:"description" api:"required,nullable"`
	// The new value of the customer's balance prior to the transaction, in the
	// customer's currency.
	EndingBalance string             `json:"ending_balance" api:"required"`
	Invoice       shared.InvoiceTiny `json:"invoice" api:"required,nullable"`
	// The original value of the customer's balance prior to the transaction, in the
	// customer's currency.
	StartingBalance string                                                                 `json:"starting_balance" api:"required"`
	Type            InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsType `json:"type" api:"required"`
	JSON            invoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionJSON  `json:"-"`
}

// invoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionJSON contains
// the JSON metadata for the struct
// [InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransaction]
type invoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionJSON struct {
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

func (r *InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionJSON) RawJSON() string {
	return r.raw
}

type InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsAction string

const (
	InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionAppliedToInvoice      InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsAction = "applied_to_invoice"
	InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionManualAdjustment      InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsAction = "manual_adjustment"
	InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionProratedRefund        InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsAction = "prorated_refund"
	InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionRevertProratedRefund  InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsAction = "revert_prorated_refund"
	InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionReturnFromVoiding     InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsAction = "return_from_voiding"
	InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionCreditNoteApplied     InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsAction = "credit_note_applied"
	InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionCreditNoteVoided      InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsAction = "credit_note_voided"
	InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionOverpaymentRefund     InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsAction = "overpayment_refund"
	InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionExternalPayment       InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsAction = "external_payment"
	InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionSmallInvoiceCarryover InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsAction = "small_invoice_carryover"
	InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionPrepaidCommitCancel   InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsAction = "prepaid_commit_cancel"
)

func (r InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsAction) IsKnown() bool {
	switch r {
	case InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionAppliedToInvoice, InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionManualAdjustment, InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionProratedRefund, InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionRevertProratedRefund, InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionReturnFromVoiding, InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionCreditNoteApplied, InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionCreditNoteVoided, InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionOverpaymentRefund, InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionExternalPayment, InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionSmallInvoiceCarryover, InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsActionPrepaidCommitCancel:
		return true
	}
	return false
}

type InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsType string

const (
	InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsTypeIncrement InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsType = "increment"
	InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsTypeDecrement InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsType = "decrement"
)

func (r InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsType) IsKnown() bool {
	switch r {
	case InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsTypeIncrement, InvoiceIssuedSummaryWebhookEventInvoiceCustomerBalanceTransactionsTypeDecrement:
		return true
	}
	return false
}

type InvoiceIssuedSummaryWebhookEventInvoiceInvoiceSource string

const (
	InvoiceIssuedSummaryWebhookEventInvoiceInvoiceSourceSubscription InvoiceIssuedSummaryWebhookEventInvoiceInvoiceSource = "subscription"
	InvoiceIssuedSummaryWebhookEventInvoiceInvoiceSourcePartial      InvoiceIssuedSummaryWebhookEventInvoiceInvoiceSource = "partial"
	InvoiceIssuedSummaryWebhookEventInvoiceInvoiceSourceOneOff       InvoiceIssuedSummaryWebhookEventInvoiceInvoiceSource = "one_off"
)

func (r InvoiceIssuedSummaryWebhookEventInvoiceInvoiceSource) IsKnown() bool {
	switch r {
	case InvoiceIssuedSummaryWebhookEventInvoiceInvoiceSourceSubscription, InvoiceIssuedSummaryWebhookEventInvoiceInvoiceSourcePartial, InvoiceIssuedSummaryWebhookEventInvoiceInvoiceSourceOneOff:
		return true
	}
	return false
}

type InvoiceIssuedSummaryWebhookEventInvoicePaymentAttempt struct {
	// The ID of the payment attempt.
	ID string `json:"id" api:"required"`
	// The amount of the payment attempt.
	Amount string `json:"amount" api:"required"`
	// The time at which the payment attempt was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The payment provider that attempted to collect the payment.
	PaymentProvider InvoiceIssuedSummaryWebhookEventInvoicePaymentAttemptsPaymentProvider `json:"payment_provider" api:"required,nullable"`
	// The ID of the payment attempt in the payment provider.
	PaymentProviderID string `json:"payment_provider_id" api:"required,nullable"`
	// URL to the downloadable PDF version of the receipt. This field will be `null`
	// for payment attempts that did not succeed.
	ReceiptPdf string `json:"receipt_pdf" api:"required,nullable"`
	// Whether the payment attempt succeeded.
	Succeeded bool                                                      `json:"succeeded" api:"required"`
	JSON      invoiceIssuedSummaryWebhookEventInvoicePaymentAttemptJSON `json:"-"`
}

// invoiceIssuedSummaryWebhookEventInvoicePaymentAttemptJSON contains the JSON
// metadata for the struct [InvoiceIssuedSummaryWebhookEventInvoicePaymentAttempt]
type invoiceIssuedSummaryWebhookEventInvoicePaymentAttemptJSON struct {
	ID                apijson.Field
	Amount            apijson.Field
	CreatedAt         apijson.Field
	PaymentProvider   apijson.Field
	PaymentProviderID apijson.Field
	ReceiptPdf        apijson.Field
	Succeeded         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceIssuedSummaryWebhookEventInvoicePaymentAttempt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceIssuedSummaryWebhookEventInvoicePaymentAttemptJSON) RawJSON() string {
	return r.raw
}

// The payment provider that attempted to collect the payment.
type InvoiceIssuedSummaryWebhookEventInvoicePaymentAttemptsPaymentProvider string

const (
	InvoiceIssuedSummaryWebhookEventInvoicePaymentAttemptsPaymentProviderStripe InvoiceIssuedSummaryWebhookEventInvoicePaymentAttemptsPaymentProvider = "stripe"
	InvoiceIssuedSummaryWebhookEventInvoicePaymentAttemptsPaymentProviderAdyen  InvoiceIssuedSummaryWebhookEventInvoicePaymentAttemptsPaymentProvider = "adyen"
)

func (r InvoiceIssuedSummaryWebhookEventInvoicePaymentAttemptsPaymentProvider) IsKnown() bool {
	switch r {
	case InvoiceIssuedSummaryWebhookEventInvoicePaymentAttemptsPaymentProviderStripe, InvoiceIssuedSummaryWebhookEventInvoicePaymentAttemptsPaymentProviderAdyen:
		return true
	}
	return false
}

type InvoiceIssuedSummaryWebhookEventInvoiceStatus string

const (
	InvoiceIssuedSummaryWebhookEventInvoiceStatusIssued InvoiceIssuedSummaryWebhookEventInvoiceStatus = "issued"
	InvoiceIssuedSummaryWebhookEventInvoiceStatusPaid   InvoiceIssuedSummaryWebhookEventInvoiceStatus = "paid"
	InvoiceIssuedSummaryWebhookEventInvoiceStatusSynced InvoiceIssuedSummaryWebhookEventInvoiceStatus = "synced"
	InvoiceIssuedSummaryWebhookEventInvoiceStatusVoid   InvoiceIssuedSummaryWebhookEventInvoiceStatus = "void"
	InvoiceIssuedSummaryWebhookEventInvoiceStatusDraft  InvoiceIssuedSummaryWebhookEventInvoiceStatus = "draft"
)

func (r InvoiceIssuedSummaryWebhookEventInvoiceStatus) IsKnown() bool {
	switch r {
	case InvoiceIssuedSummaryWebhookEventInvoiceStatusIssued, InvoiceIssuedSummaryWebhookEventInvoiceStatusPaid, InvoiceIssuedSummaryWebhookEventInvoiceStatusSynced, InvoiceIssuedSummaryWebhookEventInvoiceStatusVoid, InvoiceIssuedSummaryWebhookEventInvoiceStatusDraft:
		return true
	}
	return false
}

type InvoiceIssuedSummaryWebhookEventProperties struct {
	AutomaticallyMarkedAsPaid bool                                           `json:"automatically_marked_as_paid" api:"required"`
	JSON                      invoiceIssuedSummaryWebhookEventPropertiesJSON `json:"-"`
}

// invoiceIssuedSummaryWebhookEventPropertiesJSON contains the JSON metadata for
// the struct [InvoiceIssuedSummaryWebhookEventProperties]
type invoiceIssuedSummaryWebhookEventPropertiesJSON struct {
	AutomaticallyMarkedAsPaid apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *InvoiceIssuedSummaryWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceIssuedSummaryWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoiceIssuedSummaryWebhookEventType string

const (
	InvoiceIssuedSummaryWebhookEventTypeInvoiceIssuedSummary InvoiceIssuedSummaryWebhookEventType = "invoice.issued_summary"
)

func (r InvoiceIssuedSummaryWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceIssuedSummaryWebhookEventTypeInvoiceIssuedSummary:
		return true
	}
	return false
}

// Issued when an invoice is manually marked as paid.
type InvoiceManuallyMarkedAsPaidWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice                                    `json:"invoice" api:"required"`
	Properties InvoiceManuallyMarkedAsPaidWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceManuallyMarkedAsPaidWebhookEventType `json:"type" api:"required"`
	JSON invoiceManuallyMarkedAsPaidWebhookEventJSON `json:"-"`
}

// invoiceManuallyMarkedAsPaidWebhookEventJSON contains the JSON metadata for the
// struct [InvoiceManuallyMarkedAsPaidWebhookEvent]
type invoiceManuallyMarkedAsPaidWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceManuallyMarkedAsPaidWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceManuallyMarkedAsPaidWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceManuallyMarkedAsPaidWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoiceManuallyMarkedAsPaidWebhookEventProperties struct {
	ExternalID          string                                                `json:"external_id" api:"required,nullable"`
	Notes               string                                                `json:"notes" api:"required,nullable"`
	PaymentReceivedDate time.Time                                             `json:"payment_received_date" api:"required,nullable" format:"date-time"`
	JSON                invoiceManuallyMarkedAsPaidWebhookEventPropertiesJSON `json:"-"`
}

// invoiceManuallyMarkedAsPaidWebhookEventPropertiesJSON contains the JSON metadata
// for the struct [InvoiceManuallyMarkedAsPaidWebhookEventProperties]
type invoiceManuallyMarkedAsPaidWebhookEventPropertiesJSON struct {
	ExternalID          apijson.Field
	Notes               apijson.Field
	PaymentReceivedDate apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InvoiceManuallyMarkedAsPaidWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceManuallyMarkedAsPaidWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoiceManuallyMarkedAsPaidWebhookEventType string

const (
	InvoiceManuallyMarkedAsPaidWebhookEventTypeInvoiceManuallyMarkedAsPaid InvoiceManuallyMarkedAsPaidWebhookEventType = "invoice.manually_marked_as_paid"
)

func (r InvoiceManuallyMarkedAsPaidWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceManuallyMarkedAsPaidWebhookEventTypeInvoiceManuallyMarkedAsPaid:
		return true
	}
	return false
}

// Issued when an invoice is marked as void.
type InvoiceManuallyMarkedAsVoidWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice `json:"invoice" api:"required"`
	Properties interface{}    `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceManuallyMarkedAsVoidWebhookEventType `json:"type" api:"required"`
	JSON invoiceManuallyMarkedAsVoidWebhookEventJSON `json:"-"`
}

// invoiceManuallyMarkedAsVoidWebhookEventJSON contains the JSON metadata for the
// struct [InvoiceManuallyMarkedAsVoidWebhookEvent]
type invoiceManuallyMarkedAsVoidWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceManuallyMarkedAsVoidWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceManuallyMarkedAsVoidWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceManuallyMarkedAsVoidWebhookEvent) implementsUnwrapWebhookEvent() {}

// The event this payload describes.
type InvoiceManuallyMarkedAsVoidWebhookEventType string

const (
	InvoiceManuallyMarkedAsVoidWebhookEventTypeInvoiceManuallyMarkedAsVoid InvoiceManuallyMarkedAsVoidWebhookEventType = "invoice.manually_marked_as_void"
)

func (r InvoiceManuallyMarkedAsVoidWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceManuallyMarkedAsVoidWebhookEventTypeInvoiceManuallyMarkedAsVoid:
		return true
	}
	return false
}

// Issued when automated payment collection for an invoice fails for a configured
// payment gateway.
type InvoicePaymentFailedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice                             `json:"invoice" api:"required"`
	Properties InvoicePaymentFailedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoicePaymentFailedWebhookEventType `json:"type" api:"required"`
	JSON invoicePaymentFailedWebhookEventJSON `json:"-"`
}

// invoicePaymentFailedWebhookEventJSON contains the JSON metadata for the struct
// [InvoicePaymentFailedWebhookEvent]
type invoicePaymentFailedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoicePaymentFailedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoicePaymentFailedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoicePaymentFailedWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoicePaymentFailedWebhookEventProperties struct {
	PaymentProvider              string                                         `json:"payment_provider" api:"required,nullable"`
	PaymentProviderID            string                                         `json:"payment_provider_id" api:"required,nullable"`
	PaymentProviderTransactionID string                                         `json:"payment_provider_transaction_id" api:"required,nullable"`
	JSON                         invoicePaymentFailedWebhookEventPropertiesJSON `json:"-"`
}

// invoicePaymentFailedWebhookEventPropertiesJSON contains the JSON metadata for
// the struct [InvoicePaymentFailedWebhookEventProperties]
type invoicePaymentFailedWebhookEventPropertiesJSON struct {
	PaymentProvider              apijson.Field
	PaymentProviderID            apijson.Field
	PaymentProviderTransactionID apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *InvoicePaymentFailedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoicePaymentFailedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoicePaymentFailedWebhookEventType string

const (
	InvoicePaymentFailedWebhookEventTypeInvoicePaymentFailed InvoicePaymentFailedWebhookEventType = "invoice.payment_failed"
)

func (r InvoicePaymentFailedWebhookEventType) IsKnown() bool {
	switch r {
	case InvoicePaymentFailedWebhookEventTypeInvoicePaymentFailed:
		return true
	}
	return false
}

// Issued when an invoice payment is being processed.
type InvoicePaymentProcessingWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice                                 `json:"invoice" api:"required"`
	Properties InvoicePaymentProcessingWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoicePaymentProcessingWebhookEventType `json:"type" api:"required"`
	JSON invoicePaymentProcessingWebhookEventJSON `json:"-"`
}

// invoicePaymentProcessingWebhookEventJSON contains the JSON metadata for the
// struct [InvoicePaymentProcessingWebhookEvent]
type invoicePaymentProcessingWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoicePaymentProcessingWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoicePaymentProcessingWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoicePaymentProcessingWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoicePaymentProcessingWebhookEventProperties struct {
	PaymentProviderID string                                             `json:"payment_provider_id" api:"required,nullable"`
	PaymentProvider   string                                             `json:"payment_provider"`
	JSON              invoicePaymentProcessingWebhookEventPropertiesJSON `json:"-"`
}

// invoicePaymentProcessingWebhookEventPropertiesJSON contains the JSON metadata
// for the struct [InvoicePaymentProcessingWebhookEventProperties]
type invoicePaymentProcessingWebhookEventPropertiesJSON struct {
	PaymentProviderID apijson.Field
	PaymentProvider   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoicePaymentProcessingWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoicePaymentProcessingWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoicePaymentProcessingWebhookEventType string

const (
	InvoicePaymentProcessingWebhookEventTypeInvoicePaymentProcessing InvoicePaymentProcessingWebhookEventType = "invoice.payment_processing"
)

func (r InvoicePaymentProcessingWebhookEventType) IsKnown() bool {
	switch r {
	case InvoicePaymentProcessingWebhookEventTypeInvoicePaymentProcessing:
		return true
	}
	return false
}

// Issued when automated payment collection for an invoice succeeds for a
// configured payment gateway.
type InvoicePaymentSucceededWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice shared.Invoice `json:"invoice" api:"required"`
	// `shared_payment_token_id` is only on the wire when the payment used one.
	Properties InvoicePaymentSucceededWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoicePaymentSucceededWebhookEventType `json:"type" api:"required"`
	JSON invoicePaymentSucceededWebhookEventJSON `json:"-"`
}

// invoicePaymentSucceededWebhookEventJSON contains the JSON metadata for the
// struct [InvoicePaymentSucceededWebhookEvent]
type invoicePaymentSucceededWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoicePaymentSucceededWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoicePaymentSucceededWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoicePaymentSucceededWebhookEvent) implementsUnwrapWebhookEvent() {}

// `shared_payment_token_id` is only on the wire when the payment used one.
type InvoicePaymentSucceededWebhookEventProperties struct {
	PaymentProvider              string                                            `json:"payment_provider" api:"required,nullable"`
	PaymentProviderID            string                                            `json:"payment_provider_id" api:"required,nullable"`
	PaymentProviderTransactionID string                                            `json:"payment_provider_transaction_id" api:"required,nullable"`
	SharedPaymentTokenID         string                                            `json:"shared_payment_token_id" api:"nullable"`
	JSON                         invoicePaymentSucceededWebhookEventPropertiesJSON `json:"-"`
}

// invoicePaymentSucceededWebhookEventPropertiesJSON contains the JSON metadata for
// the struct [InvoicePaymentSucceededWebhookEventProperties]
type invoicePaymentSucceededWebhookEventPropertiesJSON struct {
	PaymentProvider              apijson.Field
	PaymentProviderID            apijson.Field
	PaymentProviderTransactionID apijson.Field
	SharedPaymentTokenID         apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *InvoicePaymentSucceededWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoicePaymentSucceededWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoicePaymentSucceededWebhookEventType string

const (
	InvoicePaymentSucceededWebhookEventTypeInvoicePaymentSucceeded InvoicePaymentSucceededWebhookEventType = "invoice.payment_succeeded"
)

func (r InvoicePaymentSucceededWebhookEventType) IsKnown() bool {
	switch r {
	case InvoicePaymentSucceededWebhookEventTypeInvoicePaymentSucceeded:
		return true
	}
	return false
}

// Issued when an invoice sync fails.
type InvoiceSyncFailedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice                          `json:"invoice" api:"required"`
	Properties InvoiceSyncFailedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceSyncFailedWebhookEventType `json:"type" api:"required"`
	JSON invoiceSyncFailedWebhookEventJSON `json:"-"`
}

// invoiceSyncFailedWebhookEventJSON contains the JSON metadata for the struct
// [InvoiceSyncFailedWebhookEvent]
type invoiceSyncFailedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceSyncFailedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceSyncFailedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceSyncFailedWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoiceSyncFailedWebhookEventProperties struct {
	PaymentProvider   string                                      `json:"payment_provider" api:"required,nullable"`
	PaymentProviderID string                                      `json:"payment_provider_id" api:"required,nullable"`
	JSON              invoiceSyncFailedWebhookEventPropertiesJSON `json:"-"`
}

// invoiceSyncFailedWebhookEventPropertiesJSON contains the JSON metadata for the
// struct [InvoiceSyncFailedWebhookEventProperties]
type invoiceSyncFailedWebhookEventPropertiesJSON struct {
	PaymentProvider   apijson.Field
	PaymentProviderID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceSyncFailedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceSyncFailedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoiceSyncFailedWebhookEventType string

const (
	InvoiceSyncFailedWebhookEventTypeInvoiceSyncFailed InvoiceSyncFailedWebhookEventType = "invoice.sync_failed"
)

func (r InvoiceSyncFailedWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceSyncFailedWebhookEventTypeInvoiceSyncFailed:
		return true
	}
	return false
}

// Issued when an invoice sync succeeds.
type InvoiceSyncSucceededWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice                             `json:"invoice" api:"required"`
	Properties InvoiceSyncSucceededWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceSyncSucceededWebhookEventType `json:"type" api:"required"`
	JSON invoiceSyncSucceededWebhookEventJSON `json:"-"`
}

// invoiceSyncSucceededWebhookEventJSON contains the JSON metadata for the struct
// [InvoiceSyncSucceededWebhookEvent]
type invoiceSyncSucceededWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceSyncSucceededWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceSyncSucceededWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceSyncSucceededWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoiceSyncSucceededWebhookEventProperties struct {
	PaymentProvider   string                                         `json:"payment_provider" api:"required,nullable"`
	PaymentProviderID string                                         `json:"payment_provider_id" api:"required,nullable"`
	JSON              invoiceSyncSucceededWebhookEventPropertiesJSON `json:"-"`
}

// invoiceSyncSucceededWebhookEventPropertiesJSON contains the JSON metadata for
// the struct [InvoiceSyncSucceededWebhookEventProperties]
type invoiceSyncSucceededWebhookEventPropertiesJSON struct {
	PaymentProvider   apijson.Field
	PaymentProviderID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceSyncSucceededWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceSyncSucceededWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoiceSyncSucceededWebhookEventType string

const (
	InvoiceSyncSucceededWebhookEventTypeInvoiceSyncSucceeded InvoiceSyncSucceededWebhookEventType = "invoice.sync_succeeded"
)

func (r InvoiceSyncSucceededWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceSyncSucceededWebhookEventTypeInvoiceSyncSucceeded:
		return true
	}
	return false
}

// Issued when an invoice is undone from marked as paid.
type InvoiceUndoMarkAsPaidWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice `json:"invoice" api:"required"`
	Properties interface{}    `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceUndoMarkAsPaidWebhookEventType `json:"type" api:"required"`
	JSON invoiceUndoMarkAsPaidWebhookEventJSON `json:"-"`
}

// invoiceUndoMarkAsPaidWebhookEventJSON contains the JSON metadata for the struct
// [InvoiceUndoMarkAsPaidWebhookEvent]
type invoiceUndoMarkAsPaidWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Invoice     apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceUndoMarkAsPaidWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceUndoMarkAsPaidWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceUndoMarkAsPaidWebhookEvent) implementsUnwrapWebhookEvent() {}

// The event this payload describes.
type InvoiceUndoMarkAsPaidWebhookEventType string

const (
	InvoiceUndoMarkAsPaidWebhookEventTypeInvoiceUndoMarkAsPaid InvoiceUndoMarkAsPaidWebhookEventType = "invoice.undo_mark_as_paid"
)

func (r InvoiceUndoMarkAsPaidWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceUndoMarkAsPaidWebhookEventTypeInvoiceUndoMarkAsPaid:
		return true
	}
	return false
}

// Issued when an invoice due date recalculation is canceled.
type InvoiceDueDateRecalculationCanceledWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                                 `json:"created_at" api:"required" format:"date-time"`
	Properties InvoiceDueDateRecalculationCanceledWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceDueDateRecalculationCanceledWebhookEventType `json:"type" api:"required"`
	JSON invoiceDueDateRecalculationCanceledWebhookEventJSON `json:"-"`
}

// invoiceDueDateRecalculationCanceledWebhookEventJSON contains the JSON metadata
// for the struct [InvoiceDueDateRecalculationCanceledWebhookEvent]
type invoiceDueDateRecalculationCanceledWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceDueDateRecalculationCanceledWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDueDateRecalculationCanceledWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceDueDateRecalculationCanceledWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoiceDueDateRecalculationCanceledWebhookEventProperties struct {
	CanceledAt time.Time                                                     `json:"canceled_at" api:"required" format:"date-time"`
	StartedAt  time.Time                                                     `json:"started_at" api:"required" format:"date-time"`
	JSON       invoiceDueDateRecalculationCanceledWebhookEventPropertiesJSON `json:"-"`
}

// invoiceDueDateRecalculationCanceledWebhookEventPropertiesJSON contains the JSON
// metadata for the struct
// [InvoiceDueDateRecalculationCanceledWebhookEventProperties]
type invoiceDueDateRecalculationCanceledWebhookEventPropertiesJSON struct {
	CanceledAt  apijson.Field
	StartedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceDueDateRecalculationCanceledWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDueDateRecalculationCanceledWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoiceDueDateRecalculationCanceledWebhookEventType string

const (
	InvoiceDueDateRecalculationCanceledWebhookEventTypeInvoiceDueDateRecalculationCanceled InvoiceDueDateRecalculationCanceledWebhookEventType = "invoice_due_date_recalculation.canceled"
)

func (r InvoiceDueDateRecalculationCanceledWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceDueDateRecalculationCanceledWebhookEventTypeInvoiceDueDateRecalculationCanceled:
		return true
	}
	return false
}

// Issued when an invoice due date recalculation is completed.
type InvoiceDueDateRecalculationCompletedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                                  `json:"created_at" api:"required" format:"date-time"`
	Properties InvoiceDueDateRecalculationCompletedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceDueDateRecalculationCompletedWebhookEventType `json:"type" api:"required"`
	JSON invoiceDueDateRecalculationCompletedWebhookEventJSON `json:"-"`
}

// invoiceDueDateRecalculationCompletedWebhookEventJSON contains the JSON metadata
// for the struct [InvoiceDueDateRecalculationCompletedWebhookEvent]
type invoiceDueDateRecalculationCompletedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceDueDateRecalculationCompletedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDueDateRecalculationCompletedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceDueDateRecalculationCompletedWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoiceDueDateRecalculationCompletedWebhookEventProperties struct {
	CompletedAt time.Time                                                      `json:"completed_at" api:"required" format:"date-time"`
	StartedAt   time.Time                                                      `json:"started_at" api:"required" format:"date-time"`
	JSON        invoiceDueDateRecalculationCompletedWebhookEventPropertiesJSON `json:"-"`
}

// invoiceDueDateRecalculationCompletedWebhookEventPropertiesJSON contains the JSON
// metadata for the struct
// [InvoiceDueDateRecalculationCompletedWebhookEventProperties]
type invoiceDueDateRecalculationCompletedWebhookEventPropertiesJSON struct {
	CompletedAt apijson.Field
	StartedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceDueDateRecalculationCompletedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDueDateRecalculationCompletedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoiceDueDateRecalculationCompletedWebhookEventType string

const (
	InvoiceDueDateRecalculationCompletedWebhookEventTypeInvoiceDueDateRecalculationCompleted InvoiceDueDateRecalculationCompletedWebhookEventType = "invoice_due_date_recalculation.completed"
)

func (r InvoiceDueDateRecalculationCompletedWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceDueDateRecalculationCompletedWebhookEventTypeInvoiceDueDateRecalculationCompleted:
		return true
	}
	return false
}

// Issued when an invoice due date recalculation is started.
type InvoiceDueDateRecalculationStartedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                                `json:"created_at" api:"required" format:"date-time"`
	Properties InvoiceDueDateRecalculationStartedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type InvoiceDueDateRecalculationStartedWebhookEventType `json:"type" api:"required"`
	JSON invoiceDueDateRecalculationStartedWebhookEventJSON `json:"-"`
}

// invoiceDueDateRecalculationStartedWebhookEventJSON contains the JSON metadata
// for the struct [InvoiceDueDateRecalculationStartedWebhookEvent]
type invoiceDueDateRecalculationStartedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceDueDateRecalculationStartedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDueDateRecalculationStartedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r InvoiceDueDateRecalculationStartedWebhookEvent) implementsUnwrapWebhookEvent() {}

type InvoiceDueDateRecalculationStartedWebhookEventProperties struct {
	StartedAt time.Time                                                    `json:"started_at" api:"required" format:"date-time"`
	JSON      invoiceDueDateRecalculationStartedWebhookEventPropertiesJSON `json:"-"`
}

// invoiceDueDateRecalculationStartedWebhookEventPropertiesJSON contains the JSON
// metadata for the struct
// [InvoiceDueDateRecalculationStartedWebhookEventProperties]
type invoiceDueDateRecalculationStartedWebhookEventPropertiesJSON struct {
	StartedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceDueDateRecalculationStartedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceDueDateRecalculationStartedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type InvoiceDueDateRecalculationStartedWebhookEventType string

const (
	InvoiceDueDateRecalculationStartedWebhookEventTypeInvoiceDueDateRecalculationStarted InvoiceDueDateRecalculationStartedWebhookEventType = "invoice_due_date_recalculation.started"
)

func (r InvoiceDueDateRecalculationStartedWebhookEventType) IsKnown() bool {
	switch r {
	case InvoiceDueDateRecalculationStartedWebhookEventTypeInvoiceDueDateRecalculationStarted:
		return true
	}
	return false
}

// Issued when metric events are dropped by watermark threshold.
type MetricEventsDroppedByWatermarkWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// `window_start` and `window_end` are ISO-8601 strings rather than datetimes: the
	// untyped message called `.isoformat()` on them, so they keep microseconds where
	// the webhook JSON provider would have truncated them.
	Properties MetricEventsDroppedByWatermarkWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type MetricEventsDroppedByWatermarkWebhookEventType `json:"type" api:"required"`
	JSON metricEventsDroppedByWatermarkWebhookEventJSON `json:"-"`
}

// metricEventsDroppedByWatermarkWebhookEventJSON contains the JSON metadata for
// the struct [MetricEventsDroppedByWatermarkWebhookEvent]
type metricEventsDroppedByWatermarkWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricEventsDroppedByWatermarkWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricEventsDroppedByWatermarkWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r MetricEventsDroppedByWatermarkWebhookEvent) implementsUnwrapWebhookEvent() {}

// `window_start` and `window_end` are ISO-8601 strings rather than datetimes: the
// untyped message called `.isoformat()` on them, so they keep microseconds where
// the webhook JSON provider would have truncated them.
type MetricEventsDroppedByWatermarkWebhookEventProperties struct {
	Dropped     int64                                                    `json:"dropped" api:"required"`
	EventName   string                                                   `json:"event_name" api:"required"`
	Total       int64                                                    `json:"total" api:"required"`
	WindowEnd   string                                                   `json:"window_end" api:"required"`
	WindowStart string                                                   `json:"window_start" api:"required"`
	JSON        metricEventsDroppedByWatermarkWebhookEventPropertiesJSON `json:"-"`
}

// metricEventsDroppedByWatermarkWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [MetricEventsDroppedByWatermarkWebhookEventProperties]
type metricEventsDroppedByWatermarkWebhookEventPropertiesJSON struct {
	Dropped     apijson.Field
	EventName   apijson.Field
	Total       apijson.Field
	WindowEnd   apijson.Field
	WindowStart apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MetricEventsDroppedByWatermarkWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r metricEventsDroppedByWatermarkWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type MetricEventsDroppedByWatermarkWebhookEventType string

const (
	MetricEventsDroppedByWatermarkWebhookEventTypeMetricEventsDroppedByWatermark MetricEventsDroppedByWatermarkWebhookEventType = "metric.events_dropped_by_watermark"
)

func (r MetricEventsDroppedByWatermarkWebhookEventType) IsKnown() bool {
	switch r {
	case MetricEventsDroppedByWatermarkWebhookEventTypeMetricEventsDroppedByWatermark:
		return true
	}
	return false
}

// Issued when a payment method is created.
type PaymentMethodCreatedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A payment method represents a customer's stored payment instrument held with an
	// external payment provider (such as Adyen or Stripe).
	//
	// The serialization is intentionally minimal for now; provider-pulled details
	// (e.g. card display metadata) will be added over time.
	PaymentMethod PaymentMethodCreatedWebhookEventPaymentMethod `json:"payment_method" api:"required"`
	Properties    interface{}                                   `json:"properties" api:"required"`
	// The event this payload describes.
	Type PaymentMethodCreatedWebhookEventType `json:"type" api:"required"`
	JSON paymentMethodCreatedWebhookEventJSON `json:"-"`
}

// paymentMethodCreatedWebhookEventJSON contains the JSON metadata for the struct
// [PaymentMethodCreatedWebhookEvent]
type paymentMethodCreatedWebhookEventJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	PaymentMethod apijson.Field
	Properties    apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PaymentMethodCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentMethodCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r PaymentMethodCreatedWebhookEvent) implementsUnwrapWebhookEvent() {}

// A payment method represents a customer's stored payment instrument held with an
// external payment provider (such as Adyen or Stripe).
//
// The serialization is intentionally minimal for now; provider-pulled details
// (e.g. card display metadata) will be added over time.
type PaymentMethodCreatedWebhookEventPaymentMethod struct {
	// The Orb-assigned unique identifier for the payment method.
	ID string `json:"id" api:"required"`
	// The time at which the payment method was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The ID of the Orb customer this payment method is attached to.
	CustomerID string `json:"customer_id" api:"required"`
	// Whether this is the customer's default payment method.
	Default bool `json:"default" api:"required"`
	// The identifier of this payment method in the external payment provider.
	ExternalPaymentMethodID string `json:"external_payment_method_id" api:"required"`
	// The type of the underlying payment instrument, e.g. `card` or `us_bank_account`.
	PaymentMethodType PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodType `json:"payment_method_type" api:"required"`
	// The external payment provider this method belongs to, derived from the linked
	// payment gateway connection (e.g. `adyen` or `stripe`). Null if the connection
	// has been removed.
	ProviderType string                                            `json:"provider_type" api:"required,nullable"`
	JSON         paymentMethodCreatedWebhookEventPaymentMethodJSON `json:"-"`
}

// paymentMethodCreatedWebhookEventPaymentMethodJSON contains the JSON metadata for
// the struct [PaymentMethodCreatedWebhookEventPaymentMethod]
type paymentMethodCreatedWebhookEventPaymentMethodJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	CustomerID              apijson.Field
	Default                 apijson.Field
	ExternalPaymentMethodID apijson.Field
	PaymentMethodType       apijson.Field
	ProviderType            apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PaymentMethodCreatedWebhookEventPaymentMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentMethodCreatedWebhookEventPaymentMethodJSON) RawJSON() string {
	return r.raw
}

// The type of the underlying payment instrument, e.g. `card` or `us_bank_account`.
type PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodType string

const (
	PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodTypeCard          PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodType = "card"
	PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodTypeUsBankAccount PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodType = "us_bank_account"
	PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodTypeLink          PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodType = "link"
	PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodTypeAmazonPay     PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodType = "amazon_pay"
	PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodTypeCrypto        PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodType = "crypto"
)

func (r PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodType) IsKnown() bool {
	switch r {
	case PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodTypeCard, PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodTypeUsBankAccount, PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodTypeLink, PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodTypeAmazonPay, PaymentMethodCreatedWebhookEventPaymentMethodPaymentMethodTypeCrypto:
		return true
	}
	return false
}

// The event this payload describes.
type PaymentMethodCreatedWebhookEventType string

const (
	PaymentMethodCreatedWebhookEventTypePaymentMethodCreated PaymentMethodCreatedWebhookEventType = "payment_method.created"
)

func (r PaymentMethodCreatedWebhookEventType) IsKnown() bool {
	switch r {
	case PaymentMethodCreatedWebhookEventTypePaymentMethodCreated:
		return true
	}
	return false
}

// Issued when a payment method is deleted.
type PaymentMethodDeletedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A payment method represents a customer's stored payment instrument held with an
	// external payment provider (such as Adyen or Stripe).
	//
	// The serialization is intentionally minimal for now; provider-pulled details
	// (e.g. card display metadata) will be added over time.
	PaymentMethod PaymentMethodDeletedWebhookEventPaymentMethod `json:"payment_method" api:"required"`
	Properties    interface{}                                   `json:"properties" api:"required"`
	// The event this payload describes.
	Type PaymentMethodDeletedWebhookEventType `json:"type" api:"required"`
	JSON paymentMethodDeletedWebhookEventJSON `json:"-"`
}

// paymentMethodDeletedWebhookEventJSON contains the JSON metadata for the struct
// [PaymentMethodDeletedWebhookEvent]
type paymentMethodDeletedWebhookEventJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	PaymentMethod apijson.Field
	Properties    apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PaymentMethodDeletedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentMethodDeletedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r PaymentMethodDeletedWebhookEvent) implementsUnwrapWebhookEvent() {}

// A payment method represents a customer's stored payment instrument held with an
// external payment provider (such as Adyen or Stripe).
//
// The serialization is intentionally minimal for now; provider-pulled details
// (e.g. card display metadata) will be added over time.
type PaymentMethodDeletedWebhookEventPaymentMethod struct {
	// The Orb-assigned unique identifier for the payment method.
	ID string `json:"id" api:"required"`
	// The time at which the payment method was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The ID of the Orb customer this payment method is attached to.
	CustomerID string `json:"customer_id" api:"required"`
	// Whether this is the customer's default payment method.
	Default bool `json:"default" api:"required"`
	// The identifier of this payment method in the external payment provider.
	ExternalPaymentMethodID string `json:"external_payment_method_id" api:"required"`
	// The type of the underlying payment instrument, e.g. `card` or `us_bank_account`.
	PaymentMethodType PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodType `json:"payment_method_type" api:"required"`
	// The external payment provider this method belongs to, derived from the linked
	// payment gateway connection (e.g. `adyen` or `stripe`). Null if the connection
	// has been removed.
	ProviderType string                                            `json:"provider_type" api:"required,nullable"`
	JSON         paymentMethodDeletedWebhookEventPaymentMethodJSON `json:"-"`
}

// paymentMethodDeletedWebhookEventPaymentMethodJSON contains the JSON metadata for
// the struct [PaymentMethodDeletedWebhookEventPaymentMethod]
type paymentMethodDeletedWebhookEventPaymentMethodJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	CustomerID              apijson.Field
	Default                 apijson.Field
	ExternalPaymentMethodID apijson.Field
	PaymentMethodType       apijson.Field
	ProviderType            apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PaymentMethodDeletedWebhookEventPaymentMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentMethodDeletedWebhookEventPaymentMethodJSON) RawJSON() string {
	return r.raw
}

// The type of the underlying payment instrument, e.g. `card` or `us_bank_account`.
type PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodType string

const (
	PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodTypeCard          PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodType = "card"
	PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodTypeUsBankAccount PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodType = "us_bank_account"
	PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodTypeLink          PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodType = "link"
	PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodTypeAmazonPay     PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodType = "amazon_pay"
	PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodTypeCrypto        PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodType = "crypto"
)

func (r PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodType) IsKnown() bool {
	switch r {
	case PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodTypeCard, PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodTypeUsBankAccount, PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodTypeLink, PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodTypeAmazonPay, PaymentMethodDeletedWebhookEventPaymentMethodPaymentMethodTypeCrypto:
		return true
	}
	return false
}

// The event this payload describes.
type PaymentMethodDeletedWebhookEventType string

const (
	PaymentMethodDeletedWebhookEventTypePaymentMethodDeleted PaymentMethodDeletedWebhookEventType = "payment_method.deleted"
)

func (r PaymentMethodDeletedWebhookEventType) IsKnown() bool {
	switch r {
	case PaymentMethodDeletedWebhookEventTypePaymentMethodDeleted:
		return true
	}
	return false
}

// Issued when a plan's default version is set.
type PlanDefaultVersionSetWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                   `json:"created_at" api:"required" format:"date-time"`
	Plan       PlanDefaultVersionSetWebhookEventPlan       `json:"plan" api:"required"`
	Properties PlanDefaultVersionSetWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type PlanDefaultVersionSetWebhookEventType `json:"type" api:"required"`
	JSON planDefaultVersionSetWebhookEventJSON `json:"-"`
}

// planDefaultVersionSetWebhookEventJSON contains the JSON metadata for the struct
// [PlanDefaultVersionSetWebhookEvent]
type planDefaultVersionSetWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Plan        apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanDefaultVersionSetWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planDefaultVersionSetWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r PlanDefaultVersionSetWebhookEvent) implementsUnwrapWebhookEvent() {}

type PlanDefaultVersionSetWebhookEventPlan struct {
	ID string `json:"id" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                                    `json:"external_plan_id" api:"required,nullable"`
	Name           string                                    `json:"name" api:"required,nullable"`
	JSON           planDefaultVersionSetWebhookEventPlanJSON `json:"-"`
}

// planDefaultVersionSetWebhookEventPlanJSON contains the JSON metadata for the
// struct [PlanDefaultVersionSetWebhookEventPlan]
type planDefaultVersionSetWebhookEventPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PlanDefaultVersionSetWebhookEventPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planDefaultVersionSetWebhookEventPlanJSON) RawJSON() string {
	return r.raw
}

type PlanDefaultVersionSetWebhookEventProperties struct {
	NewDefaultVersionNumber      int64                                           `json:"new_default_version_number" api:"required"`
	PreviousDefaultVersionNumber int64                                           `json:"previous_default_version_number" api:"required"`
	JSON                         planDefaultVersionSetWebhookEventPropertiesJSON `json:"-"`
}

// planDefaultVersionSetWebhookEventPropertiesJSON contains the JSON metadata for
// the struct [PlanDefaultVersionSetWebhookEventProperties]
type planDefaultVersionSetWebhookEventPropertiesJSON struct {
	NewDefaultVersionNumber      apijson.Field
	PreviousDefaultVersionNumber apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *PlanDefaultVersionSetWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planDefaultVersionSetWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type PlanDefaultVersionSetWebhookEventType string

const (
	PlanDefaultVersionSetWebhookEventTypePlanDefaultVersionSet PlanDefaultVersionSetWebhookEventType = "plan.default_version_set"
)

func (r PlanDefaultVersionSetWebhookEventType) IsKnown() bool {
	switch r {
	case PlanDefaultVersionSetWebhookEventTypePlanDefaultVersionSet:
		return true
	}
	return false
}

// Issued when a new plan version is created.
type PlanVersionCreatedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                `json:"created_at" api:"required" format:"date-time"`
	Plan       PlanVersionCreatedWebhookEventPlan       `json:"plan" api:"required"`
	Properties PlanVersionCreatedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type PlanVersionCreatedWebhookEventType `json:"type" api:"required"`
	JSON planVersionCreatedWebhookEventJSON `json:"-"`
}

// planVersionCreatedWebhookEventJSON contains the JSON metadata for the struct
// [PlanVersionCreatedWebhookEvent]
type planVersionCreatedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Plan        apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionCreatedWebhookEvent) implementsUnwrapWebhookEvent() {}

type PlanVersionCreatedWebhookEventPlan struct {
	ID string `json:"id" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                                 `json:"external_plan_id" api:"required,nullable"`
	Name           string                                 `json:"name" api:"required,nullable"`
	JSON           planVersionCreatedWebhookEventPlanJSON `json:"-"`
}

// planVersionCreatedWebhookEventPlanJSON contains the JSON metadata for the struct
// [PlanVersionCreatedWebhookEventPlan]
type planVersionCreatedWebhookEventPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PlanVersionCreatedWebhookEventPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionCreatedWebhookEventPlanJSON) RawJSON() string {
	return r.raw
}

type PlanVersionCreatedWebhookEventProperties struct {
	PlanVersionDescription string                                       `json:"plan_version_description" api:"required,nullable"`
	PlanVersionNumber      int64                                        `json:"plan_version_number" api:"required"`
	JSON                   planVersionCreatedWebhookEventPropertiesJSON `json:"-"`
}

// planVersionCreatedWebhookEventPropertiesJSON contains the JSON metadata for the
// struct [PlanVersionCreatedWebhookEventProperties]
type planVersionCreatedWebhookEventPropertiesJSON struct {
	PlanVersionDescription apijson.Field
	PlanVersionNumber      apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PlanVersionCreatedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionCreatedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type PlanVersionCreatedWebhookEventType string

const (
	PlanVersionCreatedWebhookEventTypePlanVersionCreated PlanVersionCreatedWebhookEventType = "plan.version_created"
)

func (r PlanVersionCreatedWebhookEventType) IsKnown() bool {
	switch r {
	case PlanVersionCreatedWebhookEventTypePlanVersionCreated:
		return true
	}
	return false
}

// Issued when a price is edited.
type PriceEditedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
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
	Price      shared.Price                      `json:"price" api:"required"`
	Properties PriceEditedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type PriceEditedWebhookEventType `json:"type" api:"required"`
	JSON priceEditedWebhookEventJSON `json:"-"`
}

// priceEditedWebhookEventJSON contains the JSON metadata for the struct
// [PriceEditedWebhookEvent]
type priceEditedWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Price       apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceEditedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceEditedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r PriceEditedWebhookEvent) implementsUnwrapWebhookEvent() {}

type PriceEditedWebhookEventProperties struct {
	// metadata values are non-null on the wire (deleting a key removes it from
	// storage); the Optional[str] values only exist on the new half of a metadata
	// FieldChange.
	PreviousAttributes PriceEditedWebhookEventPropertiesPreviousAttributes `json:"previous_attributes" api:"required"`
	JSON               priceEditedWebhookEventPropertiesJSON               `json:"-"`
}

// priceEditedWebhookEventPropertiesJSON contains the JSON metadata for the struct
// [PriceEditedWebhookEventProperties]
type priceEditedWebhookEventPropertiesJSON struct {
	PreviousAttributes apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PriceEditedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceEditedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// metadata values are non-null on the wire (deleting a key removes it from
// storage); the Optional[str] values only exist on the new half of a metadata
// FieldChange.
type PriceEditedWebhookEventPropertiesPreviousAttributes struct {
	Metadata map[string]string                                       `json:"metadata" api:"nullable"`
	JSON     priceEditedWebhookEventPropertiesPreviousAttributesJSON `json:"-"`
}

// priceEditedWebhookEventPropertiesPreviousAttributesJSON contains the JSON
// metadata for the struct [PriceEditedWebhookEventPropertiesPreviousAttributes]
type priceEditedWebhookEventPropertiesPreviousAttributesJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceEditedWebhookEventPropertiesPreviousAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceEditedWebhookEventPropertiesPreviousAttributesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type PriceEditedWebhookEventType string

const (
	PriceEditedWebhookEventTypePriceEdited PriceEditedWebhookEventType = "price.edited"
)

func (r PriceEditedWebhookEventType) IsKnown() bool {
	switch r {
	case PriceEditedWebhookEventTypePriceEdited:
		return true
	}
	return false
}

// Issued when a test webhook is sent.
type ResourceEventTestWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                               `json:"created_at" api:"required" format:"date-time"`
	Properties ResourceEventTestWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type ResourceEventTestWebhookEventType `json:"type" api:"required"`
	JSON resourceEventTestWebhookEventJSON `json:"-"`
}

// resourceEventTestWebhookEventJSON contains the JSON metadata for the struct
// [ResourceEventTestWebhookEvent]
type resourceEventTestWebhookEventJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceEventTestWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceEventTestWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r ResourceEventTestWebhookEvent) implementsUnwrapWebhookEvent() {}

type ResourceEventTestWebhookEventProperties struct {
	Message string                                      `json:"message"`
	JSON    resourceEventTestWebhookEventPropertiesJSON `json:"-"`
}

// resourceEventTestWebhookEventPropertiesJSON contains the JSON metadata for the
// struct [ResourceEventTestWebhookEventProperties]
type resourceEventTestWebhookEventPropertiesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceEventTestWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceEventTestWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type ResourceEventTestWebhookEventType string

const (
	ResourceEventTestWebhookEventTypeResourceEventTest ResourceEventTestWebhookEventType = "resource_event.test"
)

func (r ResourceEventTestWebhookEventType) IsKnown() bool {
	switch r {
	case ResourceEventTestWebhookEventTypeResourceEventTest:
		return true
	}
	return false
}

// Issued when a sales order accounting sync fails.
type SalesOrderAccountingSyncFailedWebhookEvent struct {
	// The ID of this webhook event.
	ID                   string                                                         `json:"id" api:"required"`
	AccountingSyncRecord SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecord `json:"accounting_sync_record" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice                                       `json:"invoice" api:"required"`
	Properties SalesOrderAccountingSyncFailedWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type SalesOrderAccountingSyncFailedWebhookEventType `json:"type" api:"required"`
	JSON salesOrderAccountingSyncFailedWebhookEventJSON `json:"-"`
}

// salesOrderAccountingSyncFailedWebhookEventJSON contains the JSON metadata for
// the struct [SalesOrderAccountingSyncFailedWebhookEvent]
type salesOrderAccountingSyncFailedWebhookEventJSON struct {
	ID                   apijson.Field
	AccountingSyncRecord apijson.Field
	CreatedAt            apijson.Field
	Invoice              apijson.Field
	Properties           apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SalesOrderAccountingSyncFailedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r salesOrderAccountingSyncFailedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SalesOrderAccountingSyncFailedWebhookEvent) implementsUnwrapWebhookEvent() {}

type SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecord struct {
	ID                 string                                                                   `json:"id" api:"required"`
	CustomerID         string                                                                   `json:"customer_id" api:"required"`
	RecordType         SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType `json:"record_type" api:"required"`
	ErrorDetails       map[string]interface{}                                                   `json:"error_details" api:"nullable"`
	InvoiceID          string                                                                   `json:"invoice_id" api:"nullable"`
	ProviderCustomerID string                                                                   `json:"provider_customer_id" api:"nullable"`
	Status             string                                                                   `json:"status" api:"nullable"`
	SyncAction         string                                                                   `json:"sync_action" api:"nullable"`
	JSON               salesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordJSON       `json:"-"`
}

// salesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordJSON contains the
// JSON metadata for the struct
// [SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecord]
type salesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordJSON struct {
	ID                 apijson.Field
	CustomerID         apijson.Field
	RecordType         apijson.Field
	ErrorDetails       apijson.Field
	InvoiceID          apijson.Field
	ProviderCustomerID apijson.Field
	Status             apijson.Field
	SyncAction         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r salesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordJSON) RawJSON() string {
	return r.raw
}

type SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType string

const (
	SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomer                   SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "customer"
	SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeInvoice                    SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "invoice"
	SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeTransaction                SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "transaction"
	SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "customer_balance_transaction"
	SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCreditNote                 SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "credit_note"
	SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSubscription               SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "subscription"
	SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSalesOrder                 SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "sales_order"
	SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeBlock                      SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "block"
)

func (r SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType) IsKnown() bool {
	switch r {
	case SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomer, SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeInvoice, SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeTransaction, SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction, SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCreditNote, SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSubscription, SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSalesOrder, SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeBlock:
		return true
	}
	return false
}

type SalesOrderAccountingSyncFailedWebhookEventProperties struct {
	ConnectionType string                                                   `json:"connection_type" api:"required"`
	FailureReason  string                                                   `json:"failure_reason" api:"required"`
	JSON           salesOrderAccountingSyncFailedWebhookEventPropertiesJSON `json:"-"`
}

// salesOrderAccountingSyncFailedWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [SalesOrderAccountingSyncFailedWebhookEventProperties]
type salesOrderAccountingSyncFailedWebhookEventPropertiesJSON struct {
	ConnectionType apijson.Field
	FailureReason  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SalesOrderAccountingSyncFailedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r salesOrderAccountingSyncFailedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type SalesOrderAccountingSyncFailedWebhookEventType string

const (
	SalesOrderAccountingSyncFailedWebhookEventTypeSalesOrderAccountingSyncFailed SalesOrderAccountingSyncFailedWebhookEventType = "sales_order.accounting_sync_failed"
)

func (r SalesOrderAccountingSyncFailedWebhookEventType) IsKnown() bool {
	switch r {
	case SalesOrderAccountingSyncFailedWebhookEventTypeSalesOrderAccountingSyncFailed:
		return true
	}
	return false
}

// Issued when a sales order accounting sync succeeds.
type SalesOrderAccountingSyncSucceededWebhookEvent struct {
	// The ID of this webhook event.
	ID                   string                                                            `json:"id" api:"required"`
	AccountingSyncRecord SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecord `json:"accounting_sync_record" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoice    shared.Invoice                                          `json:"invoice" api:"required"`
	Properties SalesOrderAccountingSyncSucceededWebhookEventProperties `json:"properties" api:"required"`
	// The event this payload describes.
	Type SalesOrderAccountingSyncSucceededWebhookEventType `json:"type" api:"required"`
	JSON salesOrderAccountingSyncSucceededWebhookEventJSON `json:"-"`
}

// salesOrderAccountingSyncSucceededWebhookEventJSON contains the JSON metadata for
// the struct [SalesOrderAccountingSyncSucceededWebhookEvent]
type salesOrderAccountingSyncSucceededWebhookEventJSON struct {
	ID                   apijson.Field
	AccountingSyncRecord apijson.Field
	CreatedAt            apijson.Field
	Invoice              apijson.Field
	Properties           apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SalesOrderAccountingSyncSucceededWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r salesOrderAccountingSyncSucceededWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SalesOrderAccountingSyncSucceededWebhookEvent) implementsUnwrapWebhookEvent() {}

type SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecord struct {
	ID                 string                                                                      `json:"id" api:"required"`
	CustomerID         string                                                                      `json:"customer_id" api:"required"`
	RecordType         SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType `json:"record_type" api:"required"`
	ErrorDetails       map[string]interface{}                                                      `json:"error_details" api:"nullable"`
	InvoiceID          string                                                                      `json:"invoice_id" api:"nullable"`
	ProviderCustomerID string                                                                      `json:"provider_customer_id" api:"nullable"`
	Status             string                                                                      `json:"status" api:"nullable"`
	SyncAction         string                                                                      `json:"sync_action" api:"nullable"`
	JSON               salesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON       `json:"-"`
}

// salesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON contains
// the JSON metadata for the struct
// [SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecord]
type salesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON struct {
	ID                 apijson.Field
	CustomerID         apijson.Field
	RecordType         apijson.Field
	ErrorDetails       apijson.Field
	InvoiceID          apijson.Field
	ProviderCustomerID apijson.Field
	Status             apijson.Field
	SyncAction         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r salesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON) RawJSON() string {
	return r.raw
}

type SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType string

const (
	SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomer                   SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "customer"
	SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeInvoice                    SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "invoice"
	SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeTransaction                SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "transaction"
	SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "customer_balance_transaction"
	SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCreditNote                 SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "credit_note"
	SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSubscription               SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "subscription"
	SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSalesOrder                 SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "sales_order"
	SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeBlock                      SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "block"
)

func (r SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType) IsKnown() bool {
	switch r {
	case SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomer, SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeInvoice, SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeTransaction, SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction, SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCreditNote, SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSubscription, SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSalesOrder, SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeBlock:
		return true
	}
	return false
}

type SalesOrderAccountingSyncSucceededWebhookEventProperties struct {
	ConnectionType string                                                      `json:"connection_type" api:"required"`
	JSON           salesOrderAccountingSyncSucceededWebhookEventPropertiesJSON `json:"-"`
}

// salesOrderAccountingSyncSucceededWebhookEventPropertiesJSON contains the JSON
// metadata for the struct
// [SalesOrderAccountingSyncSucceededWebhookEventProperties]
type salesOrderAccountingSyncSucceededWebhookEventPropertiesJSON struct {
	ConnectionType apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SalesOrderAccountingSyncSucceededWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r salesOrderAccountingSyncSucceededWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type SalesOrderAccountingSyncSucceededWebhookEventType string

const (
	SalesOrderAccountingSyncSucceededWebhookEventTypeSalesOrderAccountingSyncSucceeded SalesOrderAccountingSyncSucceededWebhookEventType = "sales_order.accounting_sync_succeeded"
)

func (r SalesOrderAccountingSyncSucceededWebhookEventType) IsKnown() bool {
	switch r {
	case SalesOrderAccountingSyncSucceededWebhookEventTypeSalesOrderAccountingSyncSucceeded:
		return true
	}
	return false
}

// Issued when a subscription accounting sync fails.
type SubscriptionAccountingSyncFailedWebhookEvent struct {
	// The ID of this webhook event.
	ID                   string                                                           `json:"id" api:"required"`
	AccountingSyncRecord SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecord `json:"accounting_sync_record" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                              `json:"created_at" api:"required" format:"date-time"`
	Properties SubscriptionAccountingSyncFailedWebhookEventProperties `json:"properties" api:"required"`
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
	Subscription Subscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionAccountingSyncFailedWebhookEventType `json:"type" api:"required"`
	JSON subscriptionAccountingSyncFailedWebhookEventJSON `json:"-"`
}

// subscriptionAccountingSyncFailedWebhookEventJSON contains the JSON metadata for
// the struct [SubscriptionAccountingSyncFailedWebhookEvent]
type subscriptionAccountingSyncFailedWebhookEventJSON struct {
	ID                   apijson.Field
	AccountingSyncRecord apijson.Field
	CreatedAt            apijson.Field
	Properties           apijson.Field
	Subscription         apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SubscriptionAccountingSyncFailedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAccountingSyncFailedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionAccountingSyncFailedWebhookEvent) implementsUnwrapWebhookEvent() {}

type SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecord struct {
	ID                 string                                                                     `json:"id" api:"required"`
	CustomerID         string                                                                     `json:"customer_id" api:"required"`
	RecordType         SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType `json:"record_type" api:"required"`
	ErrorDetails       map[string]interface{}                                                     `json:"error_details" api:"nullable"`
	ProviderCustomerID string                                                                     `json:"provider_customer_id" api:"nullable"`
	Status             string                                                                     `json:"status" api:"nullable"`
	SubscriptionID     string                                                                     `json:"subscription_id" api:"nullable"`
	SyncAction         string                                                                     `json:"sync_action" api:"nullable"`
	JSON               subscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordJSON       `json:"-"`
}

// subscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordJSON contains
// the JSON metadata for the struct
// [SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecord]
type subscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordJSON struct {
	ID                 apijson.Field
	CustomerID         apijson.Field
	RecordType         apijson.Field
	ErrorDetails       apijson.Field
	ProviderCustomerID apijson.Field
	Status             apijson.Field
	SubscriptionID     apijson.Field
	SyncAction         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordJSON) RawJSON() string {
	return r.raw
}

type SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType string

const (
	SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomer                   SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "customer"
	SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeInvoice                    SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "invoice"
	SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeTransaction                SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "transaction"
	SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "customer_balance_transaction"
	SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCreditNote                 SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "credit_note"
	SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSubscription               SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "subscription"
	SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSalesOrder                 SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "sales_order"
	SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeBlock                      SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "block"
)

func (r SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType) IsKnown() bool {
	switch r {
	case SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomer, SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeInvoice, SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeTransaction, SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction, SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCreditNote, SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSubscription, SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSalesOrder, SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeBlock:
		return true
	}
	return false
}

type SubscriptionAccountingSyncFailedWebhookEventProperties struct {
	ConnectionType string                                                     `json:"connection_type" api:"required"`
	FailureReason  string                                                     `json:"failure_reason" api:"required"`
	JSON           subscriptionAccountingSyncFailedWebhookEventPropertiesJSON `json:"-"`
}

// subscriptionAccountingSyncFailedWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [SubscriptionAccountingSyncFailedWebhookEventProperties]
type subscriptionAccountingSyncFailedWebhookEventPropertiesJSON struct {
	ConnectionType apijson.Field
	FailureReason  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionAccountingSyncFailedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAccountingSyncFailedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type SubscriptionAccountingSyncFailedWebhookEventType string

const (
	SubscriptionAccountingSyncFailedWebhookEventTypeSubscriptionAccountingSyncFailed SubscriptionAccountingSyncFailedWebhookEventType = "subscription.accounting_sync_failed"
)

func (r SubscriptionAccountingSyncFailedWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionAccountingSyncFailedWebhookEventTypeSubscriptionAccountingSyncFailed:
		return true
	}
	return false
}

// Issued when a subscription accounting sync succeeds.
type SubscriptionAccountingSyncSucceededWebhookEvent struct {
	// The ID of this webhook event.
	ID                   string                                                              `json:"id" api:"required"`
	AccountingSyncRecord SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecord `json:"accounting_sync_record" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                                 `json:"created_at" api:"required" format:"date-time"`
	Properties SubscriptionAccountingSyncSucceededWebhookEventProperties `json:"properties" api:"required"`
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
	Subscription Subscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionAccountingSyncSucceededWebhookEventType `json:"type" api:"required"`
	JSON subscriptionAccountingSyncSucceededWebhookEventJSON `json:"-"`
}

// subscriptionAccountingSyncSucceededWebhookEventJSON contains the JSON metadata
// for the struct [SubscriptionAccountingSyncSucceededWebhookEvent]
type subscriptionAccountingSyncSucceededWebhookEventJSON struct {
	ID                   apijson.Field
	AccountingSyncRecord apijson.Field
	CreatedAt            apijson.Field
	Properties           apijson.Field
	Subscription         apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SubscriptionAccountingSyncSucceededWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAccountingSyncSucceededWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionAccountingSyncSucceededWebhookEvent) implementsUnwrapWebhookEvent() {}

type SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecord struct {
	ID                 string                                                                        `json:"id" api:"required"`
	CustomerID         string                                                                        `json:"customer_id" api:"required"`
	RecordType         SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType `json:"record_type" api:"required"`
	ErrorDetails       map[string]interface{}                                                        `json:"error_details" api:"nullable"`
	ProviderCustomerID string                                                                        `json:"provider_customer_id" api:"nullable"`
	Status             string                                                                        `json:"status" api:"nullable"`
	SubscriptionID     string                                                                        `json:"subscription_id" api:"nullable"`
	SyncAction         string                                                                        `json:"sync_action" api:"nullable"`
	JSON               subscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON       `json:"-"`
}

// subscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON contains
// the JSON metadata for the struct
// [SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecord]
type subscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON struct {
	ID                 apijson.Field
	CustomerID         apijson.Field
	RecordType         apijson.Field
	ErrorDetails       apijson.Field
	ProviderCustomerID apijson.Field
	Status             apijson.Field
	SubscriptionID     apijson.Field
	SyncAction         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON) RawJSON() string {
	return r.raw
}

type SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType string

const (
	SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomer                   SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "customer"
	SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeInvoice                    SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "invoice"
	SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeTransaction                SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "transaction"
	SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "customer_balance_transaction"
	SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCreditNote                 SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "credit_note"
	SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSubscription               SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "subscription"
	SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSalesOrder                 SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "sales_order"
	SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeBlock                      SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "block"
)

func (r SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType) IsKnown() bool {
	switch r {
	case SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomer, SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeInvoice, SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeTransaction, SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction, SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCreditNote, SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSubscription, SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSalesOrder, SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeBlock:
		return true
	}
	return false
}

type SubscriptionAccountingSyncSucceededWebhookEventProperties struct {
	ConnectionType string                                                        `json:"connection_type" api:"required"`
	JSON           subscriptionAccountingSyncSucceededWebhookEventPropertiesJSON `json:"-"`
}

// subscriptionAccountingSyncSucceededWebhookEventPropertiesJSON contains the JSON
// metadata for the struct
// [SubscriptionAccountingSyncSucceededWebhookEventProperties]
type subscriptionAccountingSyncSucceededWebhookEventPropertiesJSON struct {
	ConnectionType apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionAccountingSyncSucceededWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAccountingSyncSucceededWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type SubscriptionAccountingSyncSucceededWebhookEventType string

const (
	SubscriptionAccountingSyncSucceededWebhookEventTypeSubscriptionAccountingSyncSucceeded SubscriptionAccountingSyncSucceededWebhookEventType = "subscription.accounting_sync_succeeded"
)

func (r SubscriptionAccountingSyncSucceededWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionAccountingSyncSucceededWebhookEventTypeSubscriptionAccountingSyncSucceeded:
		return true
	}
	return false
}

// Issued when an alert is automatically disabled by the system.
type SubscriptionAlertDisabledWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Sent when the system disables an alert on its own, currently for
	// cardinality-limit enforcement on grouped cost alerts.
	Properties SubscriptionAlertDisabledWebhookEventProperties `json:"properties" api:"required"`
	// A lightweight subscription representation for webhook payloads.
	//
	// This avoids the expensive to_subscription_params() call required for full
	// serialization.
	Subscription SubscriptionAlertDisabledWebhookEventSubscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionAlertDisabledWebhookEventType `json:"type" api:"required"`
	JSON subscriptionAlertDisabledWebhookEventJSON `json:"-"`
}

// subscriptionAlertDisabledWebhookEventJSON contains the JSON metadata for the
// struct [SubscriptionAlertDisabledWebhookEvent]
type subscriptionAlertDisabledWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionAlertDisabledWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAlertDisabledWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionAlertDisabledWebhookEvent) implementsUnwrapWebhookEvent() {}

// Sent when the system disables an alert on its own, currently for
// cardinality-limit enforcement on grouped cost alerts.
type SubscriptionAlertDisabledWebhookEventProperties struct {
	// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
	// usage, or credit balance and trigger webhooks when a threshold is exceeded.
	//
	// Alerts created through the API can be scoped to either customers or
	// subscriptions.
	AlertConfiguration SubscriptionAlertDisabledWebhookEventPropertiesAlertConfiguration `json:"alert_configuration" api:"required"`
	Cardinality        int64                                                             `json:"cardinality" api:"required,nullable"`
	CardinalityLimit   int64                                                             `json:"cardinality_limit" api:"required,nullable"`
	Reason             SubscriptionAlertDisabledWebhookEventPropertiesReason             `json:"reason" api:"required"`
	JSON               subscriptionAlertDisabledWebhookEventPropertiesJSON               `json:"-"`
}

// subscriptionAlertDisabledWebhookEventPropertiesJSON contains the JSON metadata
// for the struct [SubscriptionAlertDisabledWebhookEventProperties]
type subscriptionAlertDisabledWebhookEventPropertiesJSON struct {
	AlertConfiguration apijson.Field
	Cardinality        apijson.Field
	CardinalityLimit   apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionAlertDisabledWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAlertDisabledWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
// usage, or credit balance and trigger webhooks when a threshold is exceeded.
//
// Alerts created through the API can be scoped to either customers or
// subscriptions.
type SubscriptionAlertDisabledWebhookEventPropertiesAlertConfiguration struct {
	// Also referred to as alert_id in this documentation.
	ID string `json:"id" api:"required"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The name of the currency the credit balance or invoice cost is denominated in.
	Currency string `json:"currency" api:"required,nullable"`
	// The customer the alert applies to.
	Customer shared.CustomerMinified `json:"customer" api:"required,nullable"`
	// Whether the alert is enabled or disabled.
	Enabled bool `json:"enabled" api:"required"`
	// The metric the alert applies to.
	Metric SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationMetric `json:"metric" api:"required,nullable"`
	// The plan the alert applies to.
	Plan SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPlan `json:"plan" api:"required,nullable"`
	// The subscription the alert applies to.
	Subscription shared.SubscriptionMinified `json:"subscription" api:"required,nullable"`
	// The thresholds that define the conditions under which the alert will be
	// triggered.
	Thresholds []SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThreshold `json:"thresholds" api:"required,nullable"`
	// The type of alert. This must be a valid alert type.
	Type SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationType `json:"type" api:"required"`
	// The current status of the alert. This field is only present for credit balance
	// alerts.
	BalanceAlertStatus []SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationBalanceAlertStatus `json:"balance_alert_status" api:"nullable"`
	// The property keys to group cost alerts by. Only present for cost alerts with
	// grouping enabled.
	GroupingKeys []string `json:"grouping_keys" api:"nullable"`
	// Minified license type for alert serialization.
	LicenseType SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationLicenseType `json:"license_type" api:"nullable"`
	// Filters scoping which prices are included in spend and grouped cost alert
	// evaluation. Alerts use the price_id, item_id, and price_type fields only; the
	// alert's pricing unit is reported by currency.
	PriceFilters []SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFilter `json:"price_filters" api:"nullable"`
	// Per-group threshold overrides. Each override maps a specific combination of
	// grouping_keys values to a replacement threshold list. Only present for grouped
	// cost alerts that have at least one override.
	ThresholdOverrides []SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdOverride `json:"threshold_overrides" api:"nullable"`
	JSON               subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationJSON                `json:"-"`
}

// subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationJSON contains
// the JSON metadata for the struct
// [SubscriptionAlertDisabledWebhookEventPropertiesAlertConfiguration]
type subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Customer           apijson.Field
	Enabled            apijson.Field
	Metric             apijson.Field
	Plan               apijson.Field
	Subscription       apijson.Field
	Thresholds         apijson.Field
	Type               apijson.Field
	BalanceAlertStatus apijson.Field
	GroupingKeys       apijson.Field
	LicenseType        apijson.Field
	PriceFilters       apijson.Field
	ThresholdOverrides apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionAlertDisabledWebhookEventPropertiesAlertConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationJSON) RawJSON() string {
	return r.raw
}

// The metric the alert applies to.
type SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationMetric struct {
	ID   string                                                                      `json:"id" api:"required"`
	JSON subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationMetricJSON `json:"-"`
}

// subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationMetricJSON
// contains the JSON metadata for the struct
// [SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationMetric]
type subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationMetricJSON) RawJSON() string {
	return r.raw
}

// The plan the alert applies to.
type SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPlan struct {
	ID string `json:"id" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                                                                    `json:"external_plan_id" api:"required,nullable"`
	Name           string                                                                    `json:"name" api:"required,nullable"`
	PlanVersion    string                                                                    `json:"plan_version" api:"required"`
	JSON           subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPlanJSON `json:"-"`
}

// subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPlanJSON
// contains the JSON metadata for the struct
// [SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPlan]
type subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	PlanVersion    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPlanJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                         `json:"value" api:"required"`
	JSON  subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdJSON `json:"-"`
}

// subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdJSON
// contains the JSON metadata for the struct
// [SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThreshold]
type subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdJSON) RawJSON() string {
	return r.raw
}

// The type of alert. This must be a valid alert type.
type SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationType string

const (
	SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted          SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationType = "credit_balance_depleted"
	SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped           SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationType = "credit_balance_dropped"
	SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered         SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationType = "credit_balance_recovered"
	SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationTypeUsageExceeded                  SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationType = "usage_exceeded"
	SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationTypeCostExceeded                   SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationType = "cost_exceeded"
	SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationTypeSpendExceeded                  SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationType = "spend_exceeded"
	SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationType = "license_balance_threshold_reached"
)

func (r SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationType) IsKnown() bool {
	switch r {
	case SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted, SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped, SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered, SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationTypeUsageExceeded, SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationTypeCostExceeded, SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationTypeSpendExceeded, SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached:
		return true
	}
	return false
}

// Alert status is used to determine if an alert is currently in-alert or not.
type SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationBalanceAlertStatus struct {
	// Whether the alert is currently in-alert or not.
	InAlert bool `json:"in_alert" api:"required"`
	// The value of the threshold that defines the alert status.
	ThresholdValue string                                                                                  `json:"threshold_value" api:"required"`
	JSON           subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON `json:"-"`
}

// subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON
// contains the JSON metadata for the struct
// [SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationBalanceAlertStatus]
type subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON struct {
	InAlert        apijson.Field
	ThresholdValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationBalanceAlertStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON) RawJSON() string {
	return r.raw
}

// Minified license type for alert serialization.
type SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationLicenseType struct {
	ID   string                                                                           `json:"id" api:"required"`
	JSON subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationLicenseTypeJSON `json:"-"`
}

// subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationLicenseTypeJSON
// contains the JSON metadata for the struct
// [SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationLicenseType]
type subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationLicenseTypeJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationLicenseType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationLicenseTypeJSON) RawJSON() string {
	return r.raw
}

type SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFilter struct {
	// The property of the price to filter on.
	Field SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                                         `json:"values" api:"required"`
	JSON   subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFilterJSON `json:"-"`
}

// subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFilter]
type subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersField string

const (
	SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID       SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_id"
	SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID        SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersField = "item_id"
	SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType     SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_type"
	SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency      SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersField = "currency"
	SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersField = "pricing_unit_id"
)

func (r SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID, SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID, SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType, SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency, SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersOperator string

const (
	SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "includes"
	SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "excludes"
)

func (r SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes, SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes:
		return true
	}
	return false
}

// A per-group threshold override on a grouped cost alert.
//
// An empty `thresholds` list means the group is silenced (never fires). A
// non-empty list fully replaces the default thresholds for that group.
type SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdOverride struct {
	// The values identifying this group, ordered to match group_keys when set and the
	// alert's grouping_keys otherwise.
	GroupValues []string `json:"group_values" api:"required"`
	// The thresholds applied to this group. An empty list means the group is silenced.
	Thresholds []SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold `json:"thresholds" api:"required"`
	// The subset of the alert's grouping_keys this override binds. Null when the
	// override targets one exact group across every grouping key.
	GroupKeys []string                                                                               `json:"group_keys" api:"nullable"`
	JSON      subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON `json:"-"`
}

// subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON
// contains the JSON metadata for the struct
// [SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdOverride]
type subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON struct {
	GroupValues apijson.Field
	Thresholds  apijson.Field
	GroupKeys   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                                           `json:"value" api:"required"`
	JSON  subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON `json:"-"`
}

// subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON
// contains the JSON metadata for the struct
// [SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold]
type subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAlertDisabledWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON) RawJSON() string {
	return r.raw
}

type SubscriptionAlertDisabledWebhookEventPropertiesReason string

const (
	SubscriptionAlertDisabledWebhookEventPropertiesReasonUser                   SubscriptionAlertDisabledWebhookEventPropertiesReason = "user"
	SubscriptionAlertDisabledWebhookEventPropertiesReasonSystemCardinalityLimit SubscriptionAlertDisabledWebhookEventPropertiesReason = "system_cardinality_limit"
)

func (r SubscriptionAlertDisabledWebhookEventPropertiesReason) IsKnown() bool {
	switch r {
	case SubscriptionAlertDisabledWebhookEventPropertiesReasonUser, SubscriptionAlertDisabledWebhookEventPropertiesReasonSystemCardinalityLimit:
		return true
	}
	return false
}

// A lightweight subscription representation for webhook payloads.
//
// This avoids the expensive to_subscription_params() call required for full
// serialization.
type SubscriptionAlertDisabledWebhookEventSubscription struct {
	ID        string                                                  `json:"id" api:"required"`
	Customer  shared.CustomerMinified                                 `json:"customer" api:"required"`
	EndDate   time.Time                                               `json:"end_date" api:"required,nullable" format:"date-time"`
	Plan      SubscriptionAlertDisabledWebhookEventSubscriptionPlan   `json:"plan" api:"required,nullable"`
	StartDate time.Time                                               `json:"start_date" api:"required" format:"date-time"`
	Status    SubscriptionAlertDisabledWebhookEventSubscriptionStatus `json:"status" api:"required"`
	JSON      subscriptionAlertDisabledWebhookEventSubscriptionJSON   `json:"-"`
}

// subscriptionAlertDisabledWebhookEventSubscriptionJSON contains the JSON metadata
// for the struct [SubscriptionAlertDisabledWebhookEventSubscription]
type subscriptionAlertDisabledWebhookEventSubscriptionJSON struct {
	ID          apijson.Field
	Customer    apijson.Field
	EndDate     apijson.Field
	Plan        apijson.Field
	StartDate   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionAlertDisabledWebhookEventSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAlertDisabledWebhookEventSubscriptionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionAlertDisabledWebhookEventSubscriptionPlan struct {
	ID string `json:"id" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                                                    `json:"external_plan_id" api:"required,nullable"`
	Name           string                                                    `json:"name" api:"required,nullable"`
	JSON           subscriptionAlertDisabledWebhookEventSubscriptionPlanJSON `json:"-"`
}

// subscriptionAlertDisabledWebhookEventSubscriptionPlanJSON contains the JSON
// metadata for the struct [SubscriptionAlertDisabledWebhookEventSubscriptionPlan]
type subscriptionAlertDisabledWebhookEventSubscriptionPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionAlertDisabledWebhookEventSubscriptionPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAlertDisabledWebhookEventSubscriptionPlanJSON) RawJSON() string {
	return r.raw
}

type SubscriptionAlertDisabledWebhookEventSubscriptionStatus string

const (
	SubscriptionAlertDisabledWebhookEventSubscriptionStatusActive   SubscriptionAlertDisabledWebhookEventSubscriptionStatus = "active"
	SubscriptionAlertDisabledWebhookEventSubscriptionStatusEnded    SubscriptionAlertDisabledWebhookEventSubscriptionStatus = "ended"
	SubscriptionAlertDisabledWebhookEventSubscriptionStatusUpcoming SubscriptionAlertDisabledWebhookEventSubscriptionStatus = "upcoming"
)

func (r SubscriptionAlertDisabledWebhookEventSubscriptionStatus) IsKnown() bool {
	switch r {
	case SubscriptionAlertDisabledWebhookEventSubscriptionStatusActive, SubscriptionAlertDisabledWebhookEventSubscriptionStatusEnded, SubscriptionAlertDisabledWebhookEventSubscriptionStatusUpcoming:
		return true
	}
	return false
}

// The event this payload describes.
type SubscriptionAlertDisabledWebhookEventType string

const (
	SubscriptionAlertDisabledWebhookEventTypeSubscriptionAlertDisabled SubscriptionAlertDisabledWebhookEventType = "subscription.alert_disabled"
)

func (r SubscriptionAlertDisabledWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionAlertDisabledWebhookEventTypeSubscriptionAlertDisabled:
		return true
	}
	return false
}

// Issued when a subscription cancellation is scheduled.
type SubscriptionCancellationScheduledWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                               `json:"created_at" api:"required" format:"date-time"`
	Properties SubscriptionCancellationScheduledWebhookEventProperties `json:"properties" api:"required"`
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
	Subscription Subscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionCancellationScheduledWebhookEventType `json:"type" api:"required"`
	JSON subscriptionCancellationScheduledWebhookEventJSON `json:"-"`
}

// subscriptionCancellationScheduledWebhookEventJSON contains the JSON metadata for
// the struct [SubscriptionCancellationScheduledWebhookEvent]
type subscriptionCancellationScheduledWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionCancellationScheduledWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCancellationScheduledWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionCancellationScheduledWebhookEvent) implementsUnwrapWebhookEvent() {}

type SubscriptionCancellationScheduledWebhookEventProperties struct {
	CancellationDate string                                                      `json:"cancellation_date" api:"required"`
	JSON             subscriptionCancellationScheduledWebhookEventPropertiesJSON `json:"-"`
}

// subscriptionCancellationScheduledWebhookEventPropertiesJSON contains the JSON
// metadata for the struct
// [SubscriptionCancellationScheduledWebhookEventProperties]
type subscriptionCancellationScheduledWebhookEventPropertiesJSON struct {
	CancellationDate apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SubscriptionCancellationScheduledWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCancellationScheduledWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type SubscriptionCancellationScheduledWebhookEventType string

const (
	SubscriptionCancellationScheduledWebhookEventTypeSubscriptionCancellationScheduled SubscriptionCancellationScheduledWebhookEventType = "subscription.cancellation_scheduled"
)

func (r SubscriptionCancellationScheduledWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionCancellationScheduledWebhookEventTypeSubscriptionCancellationScheduled:
		return true
	}
	return false
}

// Issued when a scheduled subscription cancellation is unscheduled.
type SubscriptionCancellationUnscheduledWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                                 `json:"created_at" api:"required" format:"date-time"`
	Properties SubscriptionCancellationUnscheduledWebhookEventProperties `json:"properties" api:"required"`
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
	Subscription Subscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionCancellationUnscheduledWebhookEventType `json:"type" api:"required"`
	JSON subscriptionCancellationUnscheduledWebhookEventJSON `json:"-"`
}

// subscriptionCancellationUnscheduledWebhookEventJSON contains the JSON metadata
// for the struct [SubscriptionCancellationUnscheduledWebhookEvent]
type subscriptionCancellationUnscheduledWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionCancellationUnscheduledWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCancellationUnscheduledWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionCancellationUnscheduledWebhookEvent) implementsUnwrapWebhookEvent() {}

type SubscriptionCancellationUnscheduledWebhookEventProperties struct {
	OriginalCancellationDate string                                                        `json:"original_cancellation_date" api:"required"`
	JSON                     subscriptionCancellationUnscheduledWebhookEventPropertiesJSON `json:"-"`
}

// subscriptionCancellationUnscheduledWebhookEventPropertiesJSON contains the JSON
// metadata for the struct
// [SubscriptionCancellationUnscheduledWebhookEventProperties]
type subscriptionCancellationUnscheduledWebhookEventPropertiesJSON struct {
	OriginalCancellationDate apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *SubscriptionCancellationUnscheduledWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCancellationUnscheduledWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type SubscriptionCancellationUnscheduledWebhookEventType string

const (
	SubscriptionCancellationUnscheduledWebhookEventTypeSubscriptionCancellationUnscheduled SubscriptionCancellationUnscheduledWebhookEventType = "subscription.cancellation_unscheduled"
)

func (r SubscriptionCancellationUnscheduledWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionCancellationUnscheduledWebhookEventTypeSubscriptionCancellationUnscheduled:
		return true
	}
	return false
}

// Issued when a subscription's cost exceeds a pre-configured amount threshold.
type SubscriptionCostExceededWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                      `json:"created_at" api:"required" format:"date-time"`
	Properties SubscriptionCostExceededWebhookEventProperties `json:"properties" api:"required"`
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
	Subscription Subscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionCostExceededWebhookEventType `json:"type" api:"required"`
	JSON subscriptionCostExceededWebhookEventJSON `json:"-"`
}

// subscriptionCostExceededWebhookEventJSON contains the JSON metadata for the
// struct [SubscriptionCostExceededWebhookEvent]
type subscriptionCostExceededWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionCostExceededWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCostExceededWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionCostExceededWebhookEvent) implementsUnwrapWebhookEvent() {}

type SubscriptionCostExceededWebhookEventProperties struct {
	// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
	// usage, or credit balance and trigger webhooks when a threshold is exceeded.
	//
	// Alerts created through the API can be scoped to either customers or
	// subscriptions.
	AlertConfiguration SubscriptionCostExceededWebhookEventPropertiesAlertConfiguration `json:"alert_configuration" api:"required"`
	AmountThreshold    string                                                           `json:"amount_threshold" api:"required,nullable"`
	EvaluatedAmount    string                                                           `json:"evaluated_amount" api:"required,nullable"`
	TimeframeEnd       time.Time                                                        `json:"timeframe_end" api:"required" format:"date-time"`
	TimeframeStart     time.Time                                                        `json:"timeframe_start" api:"required" format:"date-time"`
	JSON               subscriptionCostExceededWebhookEventPropertiesJSON               `json:"-"`
}

// subscriptionCostExceededWebhookEventPropertiesJSON contains the JSON metadata
// for the struct [SubscriptionCostExceededWebhookEventProperties]
type subscriptionCostExceededWebhookEventPropertiesJSON struct {
	AlertConfiguration apijson.Field
	AmountThreshold    apijson.Field
	EvaluatedAmount    apijson.Field
	TimeframeEnd       apijson.Field
	TimeframeStart     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionCostExceededWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCostExceededWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
// usage, or credit balance and trigger webhooks when a threshold is exceeded.
//
// Alerts created through the API can be scoped to either customers or
// subscriptions.
type SubscriptionCostExceededWebhookEventPropertiesAlertConfiguration struct {
	// Also referred to as alert_id in this documentation.
	ID string `json:"id" api:"required"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The name of the currency the credit balance or invoice cost is denominated in.
	Currency string `json:"currency" api:"required,nullable"`
	// The customer the alert applies to.
	Customer shared.CustomerMinified `json:"customer" api:"required,nullable"`
	// Whether the alert is enabled or disabled.
	Enabled bool `json:"enabled" api:"required"`
	// The metric the alert applies to.
	Metric SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationMetric `json:"metric" api:"required,nullable"`
	// The plan the alert applies to.
	Plan SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPlan `json:"plan" api:"required,nullable"`
	// The subscription the alert applies to.
	Subscription shared.SubscriptionMinified `json:"subscription" api:"required,nullable"`
	// The thresholds that define the conditions under which the alert will be
	// triggered.
	Thresholds []SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationThreshold `json:"thresholds" api:"required,nullable"`
	// The type of alert. This must be a valid alert type.
	Type SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationType `json:"type" api:"required"`
	// The current status of the alert. This field is only present for credit balance
	// alerts.
	BalanceAlertStatus []SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatus `json:"balance_alert_status" api:"nullable"`
	// The property keys to group cost alerts by. Only present for cost alerts with
	// grouping enabled.
	GroupingKeys []string `json:"grouping_keys" api:"nullable"`
	// Minified license type for alert serialization.
	LicenseType SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationLicenseType `json:"license_type" api:"nullable"`
	// Filters scoping which prices are included in spend and grouped cost alert
	// evaluation. Alerts use the price_id, item_id, and price_type fields only; the
	// alert's pricing unit is reported by currency.
	PriceFilters []SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFilter `json:"price_filters" api:"nullable"`
	// Per-group threshold overrides. Each override maps a specific combination of
	// grouping_keys values to a replacement threshold list. Only present for grouped
	// cost alerts that have at least one override.
	ThresholdOverrides []SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverride `json:"threshold_overrides" api:"nullable"`
	JSON               subscriptionCostExceededWebhookEventPropertiesAlertConfigurationJSON                `json:"-"`
}

// subscriptionCostExceededWebhookEventPropertiesAlertConfigurationJSON contains
// the JSON metadata for the struct
// [SubscriptionCostExceededWebhookEventPropertiesAlertConfiguration]
type subscriptionCostExceededWebhookEventPropertiesAlertConfigurationJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Customer           apijson.Field
	Enabled            apijson.Field
	Metric             apijson.Field
	Plan               apijson.Field
	Subscription       apijson.Field
	Thresholds         apijson.Field
	Type               apijson.Field
	BalanceAlertStatus apijson.Field
	GroupingKeys       apijson.Field
	LicenseType        apijson.Field
	PriceFilters       apijson.Field
	ThresholdOverrides apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionCostExceededWebhookEventPropertiesAlertConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCostExceededWebhookEventPropertiesAlertConfigurationJSON) RawJSON() string {
	return r.raw
}

// The metric the alert applies to.
type SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationMetric struct {
	ID   string                                                                     `json:"id" api:"required"`
	JSON subscriptionCostExceededWebhookEventPropertiesAlertConfigurationMetricJSON `json:"-"`
}

// subscriptionCostExceededWebhookEventPropertiesAlertConfigurationMetricJSON
// contains the JSON metadata for the struct
// [SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationMetric]
type subscriptionCostExceededWebhookEventPropertiesAlertConfigurationMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCostExceededWebhookEventPropertiesAlertConfigurationMetricJSON) RawJSON() string {
	return r.raw
}

// The plan the alert applies to.
type SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPlan struct {
	ID string `json:"id" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                                                                   `json:"external_plan_id" api:"required,nullable"`
	Name           string                                                                   `json:"name" api:"required,nullable"`
	PlanVersion    string                                                                   `json:"plan_version" api:"required"`
	JSON           subscriptionCostExceededWebhookEventPropertiesAlertConfigurationPlanJSON `json:"-"`
}

// subscriptionCostExceededWebhookEventPropertiesAlertConfigurationPlanJSON
// contains the JSON metadata for the struct
// [SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPlan]
type subscriptionCostExceededWebhookEventPropertiesAlertConfigurationPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	PlanVersion    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCostExceededWebhookEventPropertiesAlertConfigurationPlanJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                        `json:"value" api:"required"`
	JSON  subscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdJSON `json:"-"`
}

// subscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdJSON
// contains the JSON metadata for the struct
// [SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationThreshold]
type subscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdJSON) RawJSON() string {
	return r.raw
}

// The type of alert. This must be a valid alert type.
type SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationType string

const (
	SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted          SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationType = "credit_balance_depleted"
	SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped           SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationType = "credit_balance_dropped"
	SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered         SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationType = "credit_balance_recovered"
	SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationTypeUsageExceeded                  SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationType = "usage_exceeded"
	SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationTypeCostExceeded                   SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationType = "cost_exceeded"
	SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationTypeSpendExceeded                  SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationType = "spend_exceeded"
	SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationType = "license_balance_threshold_reached"
)

func (r SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationType) IsKnown() bool {
	switch r {
	case SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted, SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped, SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered, SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationTypeUsageExceeded, SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationTypeCostExceeded, SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationTypeSpendExceeded, SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached:
		return true
	}
	return false
}

// Alert status is used to determine if an alert is currently in-alert or not.
type SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatus struct {
	// Whether the alert is currently in-alert or not.
	InAlert bool `json:"in_alert" api:"required"`
	// The value of the threshold that defines the alert status.
	ThresholdValue string                                                                                 `json:"threshold_value" api:"required"`
	JSON           subscriptionCostExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON `json:"-"`
}

// subscriptionCostExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON
// contains the JSON metadata for the struct
// [SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatus]
type subscriptionCostExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON struct {
	InAlert        apijson.Field
	ThresholdValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCostExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON) RawJSON() string {
	return r.raw
}

// Minified license type for alert serialization.
type SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationLicenseType struct {
	ID   string                                                                          `json:"id" api:"required"`
	JSON subscriptionCostExceededWebhookEventPropertiesAlertConfigurationLicenseTypeJSON `json:"-"`
}

// subscriptionCostExceededWebhookEventPropertiesAlertConfigurationLicenseTypeJSON
// contains the JSON metadata for the struct
// [SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationLicenseType]
type subscriptionCostExceededWebhookEventPropertiesAlertConfigurationLicenseTypeJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationLicenseType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCostExceededWebhookEventPropertiesAlertConfigurationLicenseTypeJSON) RawJSON() string {
	return r.raw
}

type SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFilter struct {
	// The property of the price to filter on.
	Field SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                                        `json:"values" api:"required"`
	JSON   subscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFilterJSON `json:"-"`
}

// subscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFilter]
type subscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField string

const (
	SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID       SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_id"
	SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID        SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "item_id"
	SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType     SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_type"
	SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency      SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "currency"
	SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "pricing_unit_id"
)

func (r SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID, SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID, SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType, SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency, SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator string

const (
	SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "includes"
	SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "excludes"
)

func (r SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes, SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes:
		return true
	}
	return false
}

// A per-group threshold override on a grouped cost alert.
//
// An empty `thresholds` list means the group is silenced (never fires). A
// non-empty list fully replaces the default thresholds for that group.
type SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverride struct {
	// The values identifying this group, ordered to match group_keys when set and the
	// alert's grouping_keys otherwise.
	GroupValues []string `json:"group_values" api:"required"`
	// The thresholds applied to this group. An empty list means the group is silenced.
	Thresholds []SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold `json:"thresholds" api:"required"`
	// The subset of the alert's grouping_keys this override binds. Null when the
	// override targets one exact group across every grouping key.
	GroupKeys []string                                                                              `json:"group_keys" api:"nullable"`
	JSON      subscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON `json:"-"`
}

// subscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON
// contains the JSON metadata for the struct
// [SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverride]
type subscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON struct {
	GroupValues apijson.Field
	Thresholds  apijson.Field
	GroupKeys   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                                          `json:"value" api:"required"`
	JSON  subscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON `json:"-"`
}

// subscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON
// contains the JSON metadata for the struct
// [SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold]
type subscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type SubscriptionCostExceededWebhookEventType string

const (
	SubscriptionCostExceededWebhookEventTypeSubscriptionCostExceeded SubscriptionCostExceededWebhookEventType = "subscription.cost_exceeded"
)

func (r SubscriptionCostExceededWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionCostExceededWebhookEventTypeSubscriptionCostExceeded:
		return true
	}
	return false
}

// Issued when a subscription resource is created.
type SubscriptionCreatedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time   `json:"created_at" api:"required" format:"date-time"`
	Properties interface{} `json:"properties" api:"required"`
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
	Subscription Subscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionCreatedWebhookEventType `json:"type" api:"required"`
	JSON subscriptionCreatedWebhookEventJSON `json:"-"`
}

// subscriptionCreatedWebhookEventJSON contains the JSON metadata for the struct
// [SubscriptionCreatedWebhookEvent]
type subscriptionCreatedWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionCreatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionCreatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionCreatedWebhookEvent) implementsUnwrapWebhookEvent() {}

// The event this payload describes.
type SubscriptionCreatedWebhookEventType string

const (
	SubscriptionCreatedWebhookEventTypeSubscriptionCreated SubscriptionCreatedWebhookEventType = "subscription.created"
)

func (r SubscriptionCreatedWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionCreatedWebhookEventTypeSubscriptionCreated:
		return true
	}
	return false
}

// Issued when a subscription is updated.
type SubscriptionEditedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                `json:"created_at" api:"required" format:"date-time"`
	Properties SubscriptionEditedWebhookEventProperties `json:"properties" api:"required"`
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
	Subscription Subscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionEditedWebhookEventType `json:"type" api:"required"`
	JSON subscriptionEditedWebhookEventJSON `json:"-"`
}

// subscriptionEditedWebhookEventJSON contains the JSON metadata for the struct
// [SubscriptionEditedWebhookEvent]
type subscriptionEditedWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionEditedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionEditedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionEditedWebhookEvent) implementsUnwrapWebhookEvent() {}

type SubscriptionEditedWebhookEventProperties struct {
	PreviousAttributes SubscriptionEditedWebhookEventPropertiesPreviousAttributes `json:"previous_attributes" api:"required"`
	JSON               subscriptionEditedWebhookEventPropertiesJSON               `json:"-"`
}

// subscriptionEditedWebhookEventPropertiesJSON contains the JSON metadata for the
// struct [SubscriptionEditedWebhookEventProperties]
type subscriptionEditedWebhookEventPropertiesJSON struct {
	PreviousAttributes apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionEditedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionEditedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

type SubscriptionEditedWebhookEventPropertiesPreviousAttributes struct {
	AutoCollection     bool                                                           `json:"auto_collection" api:"nullable"`
	AutoIssuance       bool                                                           `json:"auto_issuance" api:"nullable"`
	DefaultInvoiceMemo string                                                         `json:"default_invoice_memo" api:"nullable"`
	InvoicingThreshold string                                                         `json:"invoicing_threshold" api:"nullable"`
	Metadata           map[string]string                                              `json:"metadata" api:"nullable"`
	NetTerms           int64                                                          `json:"net_terms" api:"nullable"`
	JSON               subscriptionEditedWebhookEventPropertiesPreviousAttributesJSON `json:"-"`
}

// subscriptionEditedWebhookEventPropertiesPreviousAttributesJSON contains the JSON
// metadata for the struct
// [SubscriptionEditedWebhookEventPropertiesPreviousAttributes]
type subscriptionEditedWebhookEventPropertiesPreviousAttributesJSON struct {
	AutoCollection     apijson.Field
	AutoIssuance       apijson.Field
	DefaultInvoiceMemo apijson.Field
	InvoicingThreshold apijson.Field
	Metadata           apijson.Field
	NetTerms           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionEditedWebhookEventPropertiesPreviousAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionEditedWebhookEventPropertiesPreviousAttributesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type SubscriptionEditedWebhookEventType string

const (
	SubscriptionEditedWebhookEventTypeSubscriptionEdited SubscriptionEditedWebhookEventType = "subscription.edited"
)

func (r SubscriptionEditedWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionEditedWebhookEventTypeSubscriptionEdited:
		return true
	}
	return false
}

// Issued whenever a customer's subscription ends/lapses.
type SubscriptionEndedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time   `json:"created_at" api:"required" format:"date-time"`
	Properties interface{} `json:"properties" api:"required"`
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
	Subscription Subscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionEndedWebhookEventType `json:"type" api:"required"`
	JSON subscriptionEndedWebhookEventJSON `json:"-"`
}

// subscriptionEndedWebhookEventJSON contains the JSON metadata for the struct
// [SubscriptionEndedWebhookEvent]
type subscriptionEndedWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionEndedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionEndedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionEndedWebhookEvent) implementsUnwrapWebhookEvent() {}

// The event this payload describes.
type SubscriptionEndedWebhookEventType string

const (
	SubscriptionEndedWebhookEventTypeSubscriptionEnded SubscriptionEndedWebhookEventType = "subscription.ended"
)

func (r SubscriptionEndedWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionEndedWebhookEventTypeSubscriptionEnded:
		return true
	}
	return false
}

// Issued when a subscription's fixed fee quantity is updated.
type SubscriptionFixedFeeQuantityUpdatedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                                 `json:"created_at" api:"required" format:"date-time"`
	Properties SubscriptionFixedFeeQuantityUpdatedWebhookEventProperties `json:"properties" api:"required"`
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
	Subscription Subscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionFixedFeeQuantityUpdatedWebhookEventType `json:"type" api:"required"`
	JSON subscriptionFixedFeeQuantityUpdatedWebhookEventJSON `json:"-"`
}

// subscriptionFixedFeeQuantityUpdatedWebhookEventJSON contains the JSON metadata
// for the struct [SubscriptionFixedFeeQuantityUpdatedWebhookEvent]
type subscriptionFixedFeeQuantityUpdatedWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionFixedFeeQuantityUpdatedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionFixedFeeQuantityUpdatedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionFixedFeeQuantityUpdatedWebhookEvent) implementsUnwrapWebhookEvent() {}

type SubscriptionFixedFeeQuantityUpdatedWebhookEventProperties struct {
	EffectiveDate time.Time                                                     `json:"effective_date" api:"required" format:"date-time"`
	NewQuantity   float64                                                       `json:"new_quantity" api:"required"`
	OldQuantity   float64                                                       `json:"old_quantity" api:"required"`
	PriceID       string                                                        `json:"price_id" api:"required"`
	JSON          subscriptionFixedFeeQuantityUpdatedWebhookEventPropertiesJSON `json:"-"`
}

// subscriptionFixedFeeQuantityUpdatedWebhookEventPropertiesJSON contains the JSON
// metadata for the struct
// [SubscriptionFixedFeeQuantityUpdatedWebhookEventProperties]
type subscriptionFixedFeeQuantityUpdatedWebhookEventPropertiesJSON struct {
	EffectiveDate apijson.Field
	NewQuantity   apijson.Field
	OldQuantity   apijson.Field
	PriceID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SubscriptionFixedFeeQuantityUpdatedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionFixedFeeQuantityUpdatedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type SubscriptionFixedFeeQuantityUpdatedWebhookEventType string

const (
	SubscriptionFixedFeeQuantityUpdatedWebhookEventTypeSubscriptionFixedFeeQuantityUpdated SubscriptionFixedFeeQuantityUpdatedWebhookEventType = "subscription.fixed_fee_quantity_updated"
)

func (r SubscriptionFixedFeeQuantityUpdatedWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionFixedFeeQuantityUpdatedWebhookEventTypeSubscriptionFixedFeeQuantityUpdated:
		return true
	}
	return false
}

// Issued when grouped subscription costs exceed a pre-configured amount threshold.
type SubscriptionGroupedCostExceededWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Every group that crossed a threshold for one alert, batched into a single
	// message.
	Properties SubscriptionGroupedCostExceededWebhookEventProperties `json:"properties" api:"required"`
	// A lightweight subscription representation for webhook payloads.
	//
	// This avoids the expensive to_subscription_params() call required for full
	// serialization.
	Subscription SubscriptionGroupedCostExceededWebhookEventSubscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionGroupedCostExceededWebhookEventType `json:"type" api:"required"`
	JSON subscriptionGroupedCostExceededWebhookEventJSON `json:"-"`
}

// subscriptionGroupedCostExceededWebhookEventJSON contains the JSON metadata for
// the struct [SubscriptionGroupedCostExceededWebhookEvent]
type subscriptionGroupedCostExceededWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionGroupedCostExceededWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGroupedCostExceededWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionGroupedCostExceededWebhookEvent) implementsUnwrapWebhookEvent() {}

// Every group that crossed a threshold for one alert, batched into a single
// message.
type SubscriptionGroupedCostExceededWebhookEventProperties struct {
	// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
	// usage, or credit balance and trigger webhooks when a threshold is exceeded.
	//
	// Alerts created through the API can be scoped to either customers or
	// subscriptions.
	AlertConfiguration SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfiguration `json:"alert_configuration" api:"required"`
	GroupingKeys       []string                                                                `json:"grouping_keys" api:"required"`
	Groups             []SubscriptionGroupedCostExceededWebhookEventPropertiesGroup            `json:"groups" api:"required"`
	TimeframeEnd       time.Time                                                               `json:"timeframe_end" api:"required" format:"date-time"`
	TimeframeStart     time.Time                                                               `json:"timeframe_start" api:"required" format:"date-time"`
	JSON               subscriptionGroupedCostExceededWebhookEventPropertiesJSON               `json:"-"`
}

// subscriptionGroupedCostExceededWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [SubscriptionGroupedCostExceededWebhookEventProperties]
type subscriptionGroupedCostExceededWebhookEventPropertiesJSON struct {
	AlertConfiguration apijson.Field
	GroupingKeys       apijson.Field
	Groups             apijson.Field
	TimeframeEnd       apijson.Field
	TimeframeStart     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionGroupedCostExceededWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGroupedCostExceededWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
// usage, or credit balance and trigger webhooks when a threshold is exceeded.
//
// Alerts created through the API can be scoped to either customers or
// subscriptions.
type SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfiguration struct {
	// Also referred to as alert_id in this documentation.
	ID string `json:"id" api:"required"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The name of the currency the credit balance or invoice cost is denominated in.
	Currency string `json:"currency" api:"required,nullable"`
	// The customer the alert applies to.
	Customer shared.CustomerMinified `json:"customer" api:"required,nullable"`
	// Whether the alert is enabled or disabled.
	Enabled bool `json:"enabled" api:"required"`
	// The metric the alert applies to.
	Metric SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationMetric `json:"metric" api:"required,nullable"`
	// The plan the alert applies to.
	Plan SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPlan `json:"plan" api:"required,nullable"`
	// The subscription the alert applies to.
	Subscription shared.SubscriptionMinified `json:"subscription" api:"required,nullable"`
	// The thresholds that define the conditions under which the alert will be
	// triggered.
	Thresholds []SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThreshold `json:"thresholds" api:"required,nullable"`
	// The type of alert. This must be a valid alert type.
	Type SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationType `json:"type" api:"required"`
	// The current status of the alert. This field is only present for credit balance
	// alerts.
	BalanceAlertStatus []SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatus `json:"balance_alert_status" api:"nullable"`
	// The property keys to group cost alerts by. Only present for cost alerts with
	// grouping enabled.
	GroupingKeys []string `json:"grouping_keys" api:"nullable"`
	// Minified license type for alert serialization.
	LicenseType SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationLicenseType `json:"license_type" api:"nullable"`
	// Filters scoping which prices are included in spend and grouped cost alert
	// evaluation. Alerts use the price_id, item_id, and price_type fields only; the
	// alert's pricing unit is reported by currency.
	PriceFilters []SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFilter `json:"price_filters" api:"nullable"`
	// Per-group threshold overrides. Each override maps a specific combination of
	// grouping_keys values to a replacement threshold list. Only present for grouped
	// cost alerts that have at least one override.
	ThresholdOverrides []SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverride `json:"threshold_overrides" api:"nullable"`
	JSON               subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationJSON                `json:"-"`
}

// subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationJSON
// contains the JSON metadata for the struct
// [SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfiguration]
type subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Customer           apijson.Field
	Enabled            apijson.Field
	Metric             apijson.Field
	Plan               apijson.Field
	Subscription       apijson.Field
	Thresholds         apijson.Field
	Type               apijson.Field
	BalanceAlertStatus apijson.Field
	GroupingKeys       apijson.Field
	LicenseType        apijson.Field
	PriceFilters       apijson.Field
	ThresholdOverrides apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationJSON) RawJSON() string {
	return r.raw
}

// The metric the alert applies to.
type SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationMetric struct {
	ID   string                                                                            `json:"id" api:"required"`
	JSON subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationMetricJSON `json:"-"`
}

// subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationMetricJSON
// contains the JSON metadata for the struct
// [SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationMetric]
type subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationMetricJSON) RawJSON() string {
	return r.raw
}

// The plan the alert applies to.
type SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPlan struct {
	ID string `json:"id" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                                                                          `json:"external_plan_id" api:"required,nullable"`
	Name           string                                                                          `json:"name" api:"required,nullable"`
	PlanVersion    string                                                                          `json:"plan_version" api:"required"`
	JSON           subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPlanJSON `json:"-"`
}

// subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPlanJSON
// contains the JSON metadata for the struct
// [SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPlan]
type subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	PlanVersion    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPlanJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                               `json:"value" api:"required"`
	JSON  subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdJSON `json:"-"`
}

// subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdJSON
// contains the JSON metadata for the struct
// [SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThreshold]
type subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdJSON) RawJSON() string {
	return r.raw
}

// The type of alert. This must be a valid alert type.
type SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationType string

const (
	SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted          SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationType = "credit_balance_depleted"
	SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped           SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationType = "credit_balance_dropped"
	SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered         SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationType = "credit_balance_recovered"
	SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationTypeUsageExceeded                  SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationType = "usage_exceeded"
	SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationTypeCostExceeded                   SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationType = "cost_exceeded"
	SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationTypeSpendExceeded                  SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationType = "spend_exceeded"
	SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationType = "license_balance_threshold_reached"
)

func (r SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationType) IsKnown() bool {
	switch r {
	case SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted, SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped, SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered, SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationTypeUsageExceeded, SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationTypeCostExceeded, SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationTypeSpendExceeded, SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached:
		return true
	}
	return false
}

// Alert status is used to determine if an alert is currently in-alert or not.
type SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatus struct {
	// Whether the alert is currently in-alert or not.
	InAlert bool `json:"in_alert" api:"required"`
	// The value of the threshold that defines the alert status.
	ThresholdValue string                                                                                        `json:"threshold_value" api:"required"`
	JSON           subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON `json:"-"`
}

// subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON
// contains the JSON metadata for the struct
// [SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatus]
type subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON struct {
	InAlert        apijson.Field
	ThresholdValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON) RawJSON() string {
	return r.raw
}

// Minified license type for alert serialization.
type SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationLicenseType struct {
	ID   string                                                                                 `json:"id" api:"required"`
	JSON subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationLicenseTypeJSON `json:"-"`
}

// subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationLicenseTypeJSON
// contains the JSON metadata for the struct
// [SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationLicenseType]
type subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationLicenseTypeJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationLicenseType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationLicenseTypeJSON) RawJSON() string {
	return r.raw
}

type SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFilter struct {
	// The property of the price to filter on.
	Field SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                                               `json:"values" api:"required"`
	JSON   subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFilterJSON `json:"-"`
}

// subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFilter]
type subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField string

const (
	SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID       SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_id"
	SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID        SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "item_id"
	SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType     SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_type"
	SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency      SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "currency"
	SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "pricing_unit_id"
)

func (r SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID, SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID, SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType, SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency, SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator string

const (
	SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "includes"
	SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "excludes"
)

func (r SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes, SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes:
		return true
	}
	return false
}

// A per-group threshold override on a grouped cost alert.
//
// An empty `thresholds` list means the group is silenced (never fires). A
// non-empty list fully replaces the default thresholds for that group.
type SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverride struct {
	// The values identifying this group, ordered to match group_keys when set and the
	// alert's grouping_keys otherwise.
	GroupValues []string `json:"group_values" api:"required"`
	// The thresholds applied to this group. An empty list means the group is silenced.
	Thresholds []SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold `json:"thresholds" api:"required"`
	// The subset of the alert's grouping_keys this override binds. Null when the
	// override targets one exact group across every grouping key.
	GroupKeys []string                                                                                     `json:"group_keys" api:"nullable"`
	JSON      subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON `json:"-"`
}

// subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON
// contains the JSON metadata for the struct
// [SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverride]
type subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON struct {
	GroupValues apijson.Field
	Thresholds  apijson.Field
	GroupKeys   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                                                 `json:"value" api:"required"`
	JSON  subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON `json:"-"`
}

// subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON
// contains the JSON metadata for the struct
// [SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold]
type subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGroupedCostExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON) RawJSON() string {
	return r.raw
}

type SubscriptionGroupedCostExceededWebhookEventPropertiesGroup struct {
	AmountThreshold string                                                         `json:"amount_threshold" api:"required,nullable"`
	EvaluatedAmount string                                                         `json:"evaluated_amount" api:"required,nullable"`
	GroupValues     []string                                                       `json:"group_values" api:"required"`
	JSON            subscriptionGroupedCostExceededWebhookEventPropertiesGroupJSON `json:"-"`
}

// subscriptionGroupedCostExceededWebhookEventPropertiesGroupJSON contains the JSON
// metadata for the struct
// [SubscriptionGroupedCostExceededWebhookEventPropertiesGroup]
type subscriptionGroupedCostExceededWebhookEventPropertiesGroupJSON struct {
	AmountThreshold apijson.Field
	EvaluatedAmount apijson.Field
	GroupValues     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SubscriptionGroupedCostExceededWebhookEventPropertiesGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGroupedCostExceededWebhookEventPropertiesGroupJSON) RawJSON() string {
	return r.raw
}

// A lightweight subscription representation for webhook payloads.
//
// This avoids the expensive to_subscription_params() call required for full
// serialization.
type SubscriptionGroupedCostExceededWebhookEventSubscription struct {
	ID        string                                                        `json:"id" api:"required"`
	Customer  shared.CustomerMinified                                       `json:"customer" api:"required"`
	EndDate   time.Time                                                     `json:"end_date" api:"required,nullable" format:"date-time"`
	Plan      SubscriptionGroupedCostExceededWebhookEventSubscriptionPlan   `json:"plan" api:"required,nullable"`
	StartDate time.Time                                                     `json:"start_date" api:"required" format:"date-time"`
	Status    SubscriptionGroupedCostExceededWebhookEventSubscriptionStatus `json:"status" api:"required"`
	JSON      subscriptionGroupedCostExceededWebhookEventSubscriptionJSON   `json:"-"`
}

// subscriptionGroupedCostExceededWebhookEventSubscriptionJSON contains the JSON
// metadata for the struct
// [SubscriptionGroupedCostExceededWebhookEventSubscription]
type subscriptionGroupedCostExceededWebhookEventSubscriptionJSON struct {
	ID          apijson.Field
	Customer    apijson.Field
	EndDate     apijson.Field
	Plan        apijson.Field
	StartDate   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGroupedCostExceededWebhookEventSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGroupedCostExceededWebhookEventSubscriptionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionGroupedCostExceededWebhookEventSubscriptionPlan struct {
	ID string `json:"id" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                                                          `json:"external_plan_id" api:"required,nullable"`
	Name           string                                                          `json:"name" api:"required,nullable"`
	JSON           subscriptionGroupedCostExceededWebhookEventSubscriptionPlanJSON `json:"-"`
}

// subscriptionGroupedCostExceededWebhookEventSubscriptionPlanJSON contains the
// JSON metadata for the struct
// [SubscriptionGroupedCostExceededWebhookEventSubscriptionPlan]
type subscriptionGroupedCostExceededWebhookEventSubscriptionPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionGroupedCostExceededWebhookEventSubscriptionPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGroupedCostExceededWebhookEventSubscriptionPlanJSON) RawJSON() string {
	return r.raw
}

type SubscriptionGroupedCostExceededWebhookEventSubscriptionStatus string

const (
	SubscriptionGroupedCostExceededWebhookEventSubscriptionStatusActive   SubscriptionGroupedCostExceededWebhookEventSubscriptionStatus = "active"
	SubscriptionGroupedCostExceededWebhookEventSubscriptionStatusEnded    SubscriptionGroupedCostExceededWebhookEventSubscriptionStatus = "ended"
	SubscriptionGroupedCostExceededWebhookEventSubscriptionStatusUpcoming SubscriptionGroupedCostExceededWebhookEventSubscriptionStatus = "upcoming"
)

func (r SubscriptionGroupedCostExceededWebhookEventSubscriptionStatus) IsKnown() bool {
	switch r {
	case SubscriptionGroupedCostExceededWebhookEventSubscriptionStatusActive, SubscriptionGroupedCostExceededWebhookEventSubscriptionStatusEnded, SubscriptionGroupedCostExceededWebhookEventSubscriptionStatusUpcoming:
		return true
	}
	return false
}

// The event this payload describes.
type SubscriptionGroupedCostExceededWebhookEventType string

const (
	SubscriptionGroupedCostExceededWebhookEventTypeSubscriptionGroupedCostExceeded SubscriptionGroupedCostExceededWebhookEventType = "subscription.grouped_cost_exceeded"
)

func (r SubscriptionGroupedCostExceededWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionGroupedCostExceededWebhookEventTypeSubscriptionGroupedCostExceeded:
		return true
	}
	return false
}

// Issued when a subscription's invoicing threshold is exceeded and an evaluation
// is performed.
type SubscriptionInvoicingThresholdExceededWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                                    `json:"created_at" api:"required" format:"date-time"`
	Properties SubscriptionInvoicingThresholdExceededWebhookEventProperties `json:"properties" api:"required"`
	// A lightweight subscription representation for webhook payloads.
	//
	// This avoids the expensive to_subscription_params() call required for full
	// serialization.
	Subscription SubscriptionInvoicingThresholdExceededWebhookEventSubscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionInvoicingThresholdExceededWebhookEventType `json:"type" api:"required"`
	JSON subscriptionInvoicingThresholdExceededWebhookEventJSON `json:"-"`
}

// subscriptionInvoicingThresholdExceededWebhookEventJSON contains the JSON
// metadata for the struct [SubscriptionInvoicingThresholdExceededWebhookEvent]
type subscriptionInvoicingThresholdExceededWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionInvoicingThresholdExceededWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionInvoicingThresholdExceededWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionInvoicingThresholdExceededWebhookEvent) implementsUnwrapWebhookEvent() {}

type SubscriptionInvoicingThresholdExceededWebhookEventProperties struct {
	EvaluatedAmount         string                                                           `json:"evaluated_amount" api:"required"`
	InvoiceID               string                                                           `json:"invoice_id" api:"required"`
	InvoicingThreshold      string                                                           `json:"invoicing_threshold" api:"required"`
	ThresholdInvoiceCreated bool                                                             `json:"threshold_invoice_created" api:"required"`
	JSON                    subscriptionInvoicingThresholdExceededWebhookEventPropertiesJSON `json:"-"`
}

// subscriptionInvoicingThresholdExceededWebhookEventPropertiesJSON contains the
// JSON metadata for the struct
// [SubscriptionInvoicingThresholdExceededWebhookEventProperties]
type subscriptionInvoicingThresholdExceededWebhookEventPropertiesJSON struct {
	EvaluatedAmount         apijson.Field
	InvoiceID               apijson.Field
	InvoicingThreshold      apijson.Field
	ThresholdInvoiceCreated apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SubscriptionInvoicingThresholdExceededWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionInvoicingThresholdExceededWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// A lightweight subscription representation for webhook payloads.
//
// This avoids the expensive to_subscription_params() call required for full
// serialization.
type SubscriptionInvoicingThresholdExceededWebhookEventSubscription struct {
	ID        string                                                               `json:"id" api:"required"`
	Customer  shared.CustomerMinified                                              `json:"customer" api:"required"`
	EndDate   time.Time                                                            `json:"end_date" api:"required,nullable" format:"date-time"`
	Plan      SubscriptionInvoicingThresholdExceededWebhookEventSubscriptionPlan   `json:"plan" api:"required,nullable"`
	StartDate time.Time                                                            `json:"start_date" api:"required" format:"date-time"`
	Status    SubscriptionInvoicingThresholdExceededWebhookEventSubscriptionStatus `json:"status" api:"required"`
	JSON      subscriptionInvoicingThresholdExceededWebhookEventSubscriptionJSON   `json:"-"`
}

// subscriptionInvoicingThresholdExceededWebhookEventSubscriptionJSON contains the
// JSON metadata for the struct
// [SubscriptionInvoicingThresholdExceededWebhookEventSubscription]
type subscriptionInvoicingThresholdExceededWebhookEventSubscriptionJSON struct {
	ID          apijson.Field
	Customer    apijson.Field
	EndDate     apijson.Field
	Plan        apijson.Field
	StartDate   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionInvoicingThresholdExceededWebhookEventSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionInvoicingThresholdExceededWebhookEventSubscriptionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionInvoicingThresholdExceededWebhookEventSubscriptionPlan struct {
	ID string `json:"id" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                                                                 `json:"external_plan_id" api:"required,nullable"`
	Name           string                                                                 `json:"name" api:"required,nullable"`
	JSON           subscriptionInvoicingThresholdExceededWebhookEventSubscriptionPlanJSON `json:"-"`
}

// subscriptionInvoicingThresholdExceededWebhookEventSubscriptionPlanJSON contains
// the JSON metadata for the struct
// [SubscriptionInvoicingThresholdExceededWebhookEventSubscriptionPlan]
type subscriptionInvoicingThresholdExceededWebhookEventSubscriptionPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionInvoicingThresholdExceededWebhookEventSubscriptionPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionInvoicingThresholdExceededWebhookEventSubscriptionPlanJSON) RawJSON() string {
	return r.raw
}

type SubscriptionInvoicingThresholdExceededWebhookEventSubscriptionStatus string

const (
	SubscriptionInvoicingThresholdExceededWebhookEventSubscriptionStatusActive   SubscriptionInvoicingThresholdExceededWebhookEventSubscriptionStatus = "active"
	SubscriptionInvoicingThresholdExceededWebhookEventSubscriptionStatusEnded    SubscriptionInvoicingThresholdExceededWebhookEventSubscriptionStatus = "ended"
	SubscriptionInvoicingThresholdExceededWebhookEventSubscriptionStatusUpcoming SubscriptionInvoicingThresholdExceededWebhookEventSubscriptionStatus = "upcoming"
)

func (r SubscriptionInvoicingThresholdExceededWebhookEventSubscriptionStatus) IsKnown() bool {
	switch r {
	case SubscriptionInvoicingThresholdExceededWebhookEventSubscriptionStatusActive, SubscriptionInvoicingThresholdExceededWebhookEventSubscriptionStatusEnded, SubscriptionInvoicingThresholdExceededWebhookEventSubscriptionStatusUpcoming:
		return true
	}
	return false
}

// The event this payload describes.
type SubscriptionInvoicingThresholdExceededWebhookEventType string

const (
	SubscriptionInvoicingThresholdExceededWebhookEventTypeSubscriptionInvoicingThresholdExceeded SubscriptionInvoicingThresholdExceededWebhookEventType = "subscription.invoicing_threshold_exceeded"
)

func (r SubscriptionInvoicingThresholdExceededWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionInvoicingThresholdExceededWebhookEventTypeSubscriptionInvoicingThresholdExceeded:
		return true
	}
	return false
}

// Issued when a license allocation is reset.
type SubscriptionLicenseAllocationResetWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Fires at the start of a billing period when license allocations are replenished.
	// Allocations sharing a billing period are batched into one message.
	Properties SubscriptionLicenseAllocationResetWebhookEventProperties `json:"properties" api:"required"`
	// A lightweight subscription representation for webhook payloads.
	//
	// This avoids the expensive to_subscription_params() call required for full
	// serialization.
	Subscription SubscriptionLicenseAllocationResetWebhookEventSubscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionLicenseAllocationResetWebhookEventType `json:"type" api:"required"`
	JSON subscriptionLicenseAllocationResetWebhookEventJSON `json:"-"`
}

// subscriptionLicenseAllocationResetWebhookEventJSON contains the JSON metadata
// for the struct [SubscriptionLicenseAllocationResetWebhookEvent]
type subscriptionLicenseAllocationResetWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionLicenseAllocationResetWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseAllocationResetWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionLicenseAllocationResetWebhookEvent) implementsUnwrapWebhookEvent() {}

// Fires at the start of a billing period when license allocations are replenished.
// Allocations sharing a billing period are batched into one message.
type SubscriptionLicenseAllocationResetWebhookEventProperties struct {
	ResetAllocations []SubscriptionLicenseAllocationResetWebhookEventPropertiesResetAllocation `json:"reset_allocations" api:"required"`
	TimeframeEnd     time.Time                                                                 `json:"timeframe_end" api:"required" format:"date-time"`
	TimeframeStart   time.Time                                                                 `json:"timeframe_start" api:"required" format:"date-time"`
	JSON             subscriptionLicenseAllocationResetWebhookEventPropertiesJSON              `json:"-"`
}

// subscriptionLicenseAllocationResetWebhookEventPropertiesJSON contains the JSON
// metadata for the struct
// [SubscriptionLicenseAllocationResetWebhookEventProperties]
type subscriptionLicenseAllocationResetWebhookEventPropertiesJSON struct {
	ResetAllocations apijson.Field
	TimeframeEnd     apijson.Field
	TimeframeStart   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SubscriptionLicenseAllocationResetWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseAllocationResetWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// A license allocation replenished at the start of a billing period.
type SubscriptionLicenseAllocationResetWebhookEventPropertiesResetAllocation struct {
	AllocationAmount string                                                                      `json:"allocation_amount" api:"required"`
	LicenseTypeID    string                                                                      `json:"license_type_id" api:"required"`
	PricingUnitID    string                                                                      `json:"pricing_unit_id" api:"required"`
	JSON             subscriptionLicenseAllocationResetWebhookEventPropertiesResetAllocationJSON `json:"-"`
}

// subscriptionLicenseAllocationResetWebhookEventPropertiesResetAllocationJSON
// contains the JSON metadata for the struct
// [SubscriptionLicenseAllocationResetWebhookEventPropertiesResetAllocation]
type subscriptionLicenseAllocationResetWebhookEventPropertiesResetAllocationJSON struct {
	AllocationAmount apijson.Field
	LicenseTypeID    apijson.Field
	PricingUnitID    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SubscriptionLicenseAllocationResetWebhookEventPropertiesResetAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseAllocationResetWebhookEventPropertiesResetAllocationJSON) RawJSON() string {
	return r.raw
}

// A lightweight subscription representation for webhook payloads.
//
// This avoids the expensive to_subscription_params() call required for full
// serialization.
type SubscriptionLicenseAllocationResetWebhookEventSubscription struct {
	ID        string                                                           `json:"id" api:"required"`
	Customer  shared.CustomerMinified                                          `json:"customer" api:"required"`
	EndDate   time.Time                                                        `json:"end_date" api:"required,nullable" format:"date-time"`
	Plan      SubscriptionLicenseAllocationResetWebhookEventSubscriptionPlan   `json:"plan" api:"required,nullable"`
	StartDate time.Time                                                        `json:"start_date" api:"required" format:"date-time"`
	Status    SubscriptionLicenseAllocationResetWebhookEventSubscriptionStatus `json:"status" api:"required"`
	JSON      subscriptionLicenseAllocationResetWebhookEventSubscriptionJSON   `json:"-"`
}

// subscriptionLicenseAllocationResetWebhookEventSubscriptionJSON contains the JSON
// metadata for the struct
// [SubscriptionLicenseAllocationResetWebhookEventSubscription]
type subscriptionLicenseAllocationResetWebhookEventSubscriptionJSON struct {
	ID          apijson.Field
	Customer    apijson.Field
	EndDate     apijson.Field
	Plan        apijson.Field
	StartDate   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionLicenseAllocationResetWebhookEventSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseAllocationResetWebhookEventSubscriptionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionLicenseAllocationResetWebhookEventSubscriptionPlan struct {
	ID string `json:"id" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                                                             `json:"external_plan_id" api:"required,nullable"`
	Name           string                                                             `json:"name" api:"required,nullable"`
	JSON           subscriptionLicenseAllocationResetWebhookEventSubscriptionPlanJSON `json:"-"`
}

// subscriptionLicenseAllocationResetWebhookEventSubscriptionPlanJSON contains the
// JSON metadata for the struct
// [SubscriptionLicenseAllocationResetWebhookEventSubscriptionPlan]
type subscriptionLicenseAllocationResetWebhookEventSubscriptionPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionLicenseAllocationResetWebhookEventSubscriptionPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseAllocationResetWebhookEventSubscriptionPlanJSON) RawJSON() string {
	return r.raw
}

type SubscriptionLicenseAllocationResetWebhookEventSubscriptionStatus string

const (
	SubscriptionLicenseAllocationResetWebhookEventSubscriptionStatusActive   SubscriptionLicenseAllocationResetWebhookEventSubscriptionStatus = "active"
	SubscriptionLicenseAllocationResetWebhookEventSubscriptionStatusEnded    SubscriptionLicenseAllocationResetWebhookEventSubscriptionStatus = "ended"
	SubscriptionLicenseAllocationResetWebhookEventSubscriptionStatusUpcoming SubscriptionLicenseAllocationResetWebhookEventSubscriptionStatus = "upcoming"
)

func (r SubscriptionLicenseAllocationResetWebhookEventSubscriptionStatus) IsKnown() bool {
	switch r {
	case SubscriptionLicenseAllocationResetWebhookEventSubscriptionStatusActive, SubscriptionLicenseAllocationResetWebhookEventSubscriptionStatusEnded, SubscriptionLicenseAllocationResetWebhookEventSubscriptionStatusUpcoming:
		return true
	}
	return false
}

// The event this payload describes.
type SubscriptionLicenseAllocationResetWebhookEventType string

const (
	SubscriptionLicenseAllocationResetWebhookEventTypeSubscriptionLicenseAllocationReset SubscriptionLicenseAllocationResetWebhookEventType = "subscription.license_allocation_reset"
)

func (r SubscriptionLicenseAllocationResetWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionLicenseAllocationResetWebhookEventTypeSubscriptionLicenseAllocationReset:
		return true
	}
	return false
}

// Issued when a license balance threshold is reached.
type SubscriptionLicenseBalanceThresholdReachedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Every license that crossed a threshold for one alert, batched into a single
	// message. For account-wide alerts the licenses may span several license types.
	Properties SubscriptionLicenseBalanceThresholdReachedWebhookEventProperties `json:"properties" api:"required"`
	// A lightweight subscription representation for webhook payloads.
	//
	// This avoids the expensive to_subscription_params() call required for full
	// serialization.
	Subscription SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionLicenseBalanceThresholdReachedWebhookEventType `json:"type" api:"required"`
	JSON subscriptionLicenseBalanceThresholdReachedWebhookEventJSON `json:"-"`
}

// subscriptionLicenseBalanceThresholdReachedWebhookEventJSON contains the JSON
// metadata for the struct [SubscriptionLicenseBalanceThresholdReachedWebhookEvent]
type subscriptionLicenseBalanceThresholdReachedWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionLicenseBalanceThresholdReachedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseBalanceThresholdReachedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionLicenseBalanceThresholdReachedWebhookEvent) implementsUnwrapWebhookEvent() {}

// Every license that crossed a threshold for one alert, batched into a single
// message. For account-wide alerts the licenses may span several license types.
type SubscriptionLicenseBalanceThresholdReachedWebhookEventProperties struct {
	// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
	// usage, or credit balance and trigger webhooks when a threshold is exceeded.
	//
	// Alerts created through the API can be scoped to either customers or
	// subscriptions.
	AlertConfiguration SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfiguration `json:"alert_configuration" api:"required"`
	Licenses           []SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesLicense          `json:"licenses" api:"required"`
	TimeframeEnd       time.Time                                                                          `json:"timeframe_end" api:"required" format:"date-time"`
	TimeframeStart     time.Time                                                                          `json:"timeframe_start" api:"required" format:"date-time"`
	JSON               subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesJSON               `json:"-"`
}

// subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesJSON contains
// the JSON metadata for the struct
// [SubscriptionLicenseBalanceThresholdReachedWebhookEventProperties]
type subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesJSON struct {
	AlertConfiguration apijson.Field
	Licenses           apijson.Field
	TimeframeEnd       apijson.Field
	TimeframeStart     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionLicenseBalanceThresholdReachedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
// usage, or credit balance and trigger webhooks when a threshold is exceeded.
//
// Alerts created through the API can be scoped to either customers or
// subscriptions.
type SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfiguration struct {
	// Also referred to as alert_id in this documentation.
	ID string `json:"id" api:"required"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The name of the currency the credit balance or invoice cost is denominated in.
	Currency string `json:"currency" api:"required,nullable"`
	// The customer the alert applies to.
	Customer shared.CustomerMinified `json:"customer" api:"required,nullable"`
	// Whether the alert is enabled or disabled.
	Enabled bool `json:"enabled" api:"required"`
	// The metric the alert applies to.
	Metric SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationMetric `json:"metric" api:"required,nullable"`
	// The plan the alert applies to.
	Plan SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPlan `json:"plan" api:"required,nullable"`
	// The subscription the alert applies to.
	Subscription shared.SubscriptionMinified `json:"subscription" api:"required,nullable"`
	// The thresholds that define the conditions under which the alert will be
	// triggered.
	Thresholds []SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThreshold `json:"thresholds" api:"required,nullable"`
	// The type of alert. This must be a valid alert type.
	Type SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationType `json:"type" api:"required"`
	// The current status of the alert. This field is only present for credit balance
	// alerts.
	BalanceAlertStatus []SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationBalanceAlertStatus `json:"balance_alert_status" api:"nullable"`
	// The property keys to group cost alerts by. Only present for cost alerts with
	// grouping enabled.
	GroupingKeys []string `json:"grouping_keys" api:"nullable"`
	// Minified license type for alert serialization.
	LicenseType SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationLicenseType `json:"license_type" api:"nullable"`
	// Filters scoping which prices are included in spend and grouped cost alert
	// evaluation. Alerts use the price_id, item_id, and price_type fields only; the
	// alert's pricing unit is reported by currency.
	PriceFilters []SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFilter `json:"price_filters" api:"nullable"`
	// Per-group threshold overrides. Each override maps a specific combination of
	// grouping_keys values to a replacement threshold list. Only present for grouped
	// cost alerts that have at least one override.
	ThresholdOverrides []SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdOverride `json:"threshold_overrides" api:"nullable"`
	JSON               subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationJSON                `json:"-"`
}

// subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationJSON
// contains the JSON metadata for the struct
// [SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfiguration]
type subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Customer           apijson.Field
	Enabled            apijson.Field
	Metric             apijson.Field
	Plan               apijson.Field
	Subscription       apijson.Field
	Thresholds         apijson.Field
	Type               apijson.Field
	BalanceAlertStatus apijson.Field
	GroupingKeys       apijson.Field
	LicenseType        apijson.Field
	PriceFilters       apijson.Field
	ThresholdOverrides apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationJSON) RawJSON() string {
	return r.raw
}

// The metric the alert applies to.
type SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationMetric struct {
	ID   string                                                                                       `json:"id" api:"required"`
	JSON subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationMetricJSON `json:"-"`
}

// subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationMetricJSON
// contains the JSON metadata for the struct
// [SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationMetric]
type subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationMetricJSON) RawJSON() string {
	return r.raw
}

// The plan the alert applies to.
type SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPlan struct {
	ID string `json:"id" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                                                                                     `json:"external_plan_id" api:"required,nullable"`
	Name           string                                                                                     `json:"name" api:"required,nullable"`
	PlanVersion    string                                                                                     `json:"plan_version" api:"required"`
	JSON           subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPlanJSON `json:"-"`
}

// subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPlanJSON
// contains the JSON metadata for the struct
// [SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPlan]
type subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	PlanVersion    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPlanJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                                          `json:"value" api:"required"`
	JSON  subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdJSON `json:"-"`
}

// subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdJSON
// contains the JSON metadata for the struct
// [SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThreshold]
type subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdJSON) RawJSON() string {
	return r.raw
}

// The type of alert. This must be a valid alert type.
type SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationType string

const (
	SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted          SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationType = "credit_balance_depleted"
	SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped           SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationType = "credit_balance_dropped"
	SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered         SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationType = "credit_balance_recovered"
	SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationTypeUsageExceeded                  SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationType = "usage_exceeded"
	SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationTypeCostExceeded                   SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationType = "cost_exceeded"
	SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationTypeSpendExceeded                  SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationType = "spend_exceeded"
	SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationType = "license_balance_threshold_reached"
)

func (r SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationType) IsKnown() bool {
	switch r {
	case SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted, SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped, SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered, SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationTypeUsageExceeded, SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationTypeCostExceeded, SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationTypeSpendExceeded, SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached:
		return true
	}
	return false
}

// Alert status is used to determine if an alert is currently in-alert or not.
type SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationBalanceAlertStatus struct {
	// Whether the alert is currently in-alert or not.
	InAlert bool `json:"in_alert" api:"required"`
	// The value of the threshold that defines the alert status.
	ThresholdValue string                                                                                                   `json:"threshold_value" api:"required"`
	JSON           subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON `json:"-"`
}

// subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON
// contains the JSON metadata for the struct
// [SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationBalanceAlertStatus]
type subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON struct {
	InAlert        apijson.Field
	ThresholdValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationBalanceAlertStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON) RawJSON() string {
	return r.raw
}

// Minified license type for alert serialization.
type SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationLicenseType struct {
	ID   string                                                                                            `json:"id" api:"required"`
	JSON subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationLicenseTypeJSON `json:"-"`
}

// subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationLicenseTypeJSON
// contains the JSON metadata for the struct
// [SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationLicenseType]
type subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationLicenseTypeJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationLicenseType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationLicenseTypeJSON) RawJSON() string {
	return r.raw
}

type SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFilter struct {
	// The property of the price to filter on.
	Field SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                                                          `json:"values" api:"required"`
	JSON   subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFilterJSON `json:"-"`
}

// subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFilter]
type subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersField string

const (
	SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID       SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_id"
	SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID        SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersField = "item_id"
	SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType     SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_type"
	SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency      SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersField = "currency"
	SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersField = "pricing_unit_id"
)

func (r SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID, SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID, SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType, SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency, SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersOperator string

const (
	SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "includes"
	SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "excludes"
)

func (r SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes, SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes:
		return true
	}
	return false
}

// A per-group threshold override on a grouped cost alert.
//
// An empty `thresholds` list means the group is silenced (never fires). A
// non-empty list fully replaces the default thresholds for that group.
type SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdOverride struct {
	// The values identifying this group, ordered to match group_keys when set and the
	// alert's grouping_keys otherwise.
	GroupValues []string `json:"group_values" api:"required"`
	// The thresholds applied to this group. An empty list means the group is silenced.
	Thresholds []SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold `json:"thresholds" api:"required"`
	// The subset of the alert's grouping_keys this override binds. Null when the
	// override targets one exact group across every grouping key.
	GroupKeys []string                                                                                                `json:"group_keys" api:"nullable"`
	JSON      subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON `json:"-"`
}

// subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON
// contains the JSON metadata for the struct
// [SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdOverride]
type subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON struct {
	GroupValues apijson.Field
	Thresholds  apijson.Field
	GroupKeys   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                                                            `json:"value" api:"required"`
	JSON  subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON `json:"-"`
}

// subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON
// contains the JSON metadata for the struct
// [SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold]
type subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON) RawJSON() string {
	return r.raw
}

type SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesLicense struct {
	ExternalLicenseID   string                                                                      `json:"external_license_id" api:"required"`
	LicenseTypeID       string                                                                      `json:"license_type_id" api:"required"`
	ThresholdPercentage string                                                                      `json:"threshold_percentage" api:"required"`
	JSON                subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesLicenseJSON `json:"-"`
}

// subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesLicenseJSON
// contains the JSON metadata for the struct
// [SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesLicense]
type subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesLicenseJSON struct {
	ExternalLicenseID   apijson.Field
	LicenseTypeID       apijson.Field
	ThresholdPercentage apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SubscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesLicense) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseBalanceThresholdReachedWebhookEventPropertiesLicenseJSON) RawJSON() string {
	return r.raw
}

// A lightweight subscription representation for webhook payloads.
//
// This avoids the expensive to_subscription_params() call required for full
// serialization.
type SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscription struct {
	ID        string                                                                   `json:"id" api:"required"`
	Customer  shared.CustomerMinified                                                  `json:"customer" api:"required"`
	EndDate   time.Time                                                                `json:"end_date" api:"required,nullable" format:"date-time"`
	Plan      SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionPlan   `json:"plan" api:"required,nullable"`
	StartDate time.Time                                                                `json:"start_date" api:"required" format:"date-time"`
	Status    SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionStatus `json:"status" api:"required"`
	JSON      subscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionJSON   `json:"-"`
}

// subscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionJSON contains
// the JSON metadata for the struct
// [SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscription]
type subscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionJSON struct {
	ID          apijson.Field
	Customer    apijson.Field
	EndDate     apijson.Field
	Plan        apijson.Field
	StartDate   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionPlan struct {
	ID string `json:"id" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                                                                     `json:"external_plan_id" api:"required,nullable"`
	Name           string                                                                     `json:"name" api:"required,nullable"`
	JSON           subscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionPlanJSON `json:"-"`
}

// subscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionPlanJSON
// contains the JSON metadata for the struct
// [SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionPlan]
type subscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionPlanJSON) RawJSON() string {
	return r.raw
}

type SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionStatus string

const (
	SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionStatusActive   SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionStatus = "active"
	SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionStatusEnded    SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionStatus = "ended"
	SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionStatusUpcoming SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionStatus = "upcoming"
)

func (r SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionStatus) IsKnown() bool {
	switch r {
	case SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionStatusActive, SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionStatusEnded, SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscriptionStatusUpcoming:
		return true
	}
	return false
}

// The event this payload describes.
type SubscriptionLicenseBalanceThresholdReachedWebhookEventType string

const (
	SubscriptionLicenseBalanceThresholdReachedWebhookEventTypeSubscriptionLicenseBalanceThresholdReached SubscriptionLicenseBalanceThresholdReachedWebhookEventType = "subscription.license_balance_threshold_reached"
)

func (r SubscriptionLicenseBalanceThresholdReachedWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionLicenseBalanceThresholdReachedWebhookEventTypeSubscriptionLicenseBalanceThresholdReached:
		return true
	}
	return false
}

// Issued when a subscription plan change is scheduled.
type SubscriptionPlanChangeScheduledWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                             `json:"created_at" api:"required" format:"date-time"`
	Properties SubscriptionPlanChangeScheduledWebhookEventProperties `json:"properties" api:"required"`
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
	Subscription Subscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionPlanChangeScheduledWebhookEventType `json:"type" api:"required"`
	JSON subscriptionPlanChangeScheduledWebhookEventJSON `json:"-"`
}

// subscriptionPlanChangeScheduledWebhookEventJSON contains the JSON metadata for
// the struct [SubscriptionPlanChangeScheduledWebhookEvent]
type subscriptionPlanChangeScheduledWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionPlanChangeScheduledWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPlanChangeScheduledWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionPlanChangeScheduledWebhookEvent) implementsUnwrapWebhookEvent() {}

type SubscriptionPlanChangeScheduledWebhookEventProperties struct {
	ChangeDate     string                                                    `json:"change_date" api:"required"`
	NewPlanID      string                                                    `json:"new_plan_id" api:"required"`
	PreviousPlanID string                                                    `json:"previous_plan_id" api:"required"`
	JSON           subscriptionPlanChangeScheduledWebhookEventPropertiesJSON `json:"-"`
}

// subscriptionPlanChangeScheduledWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [SubscriptionPlanChangeScheduledWebhookEventProperties]
type subscriptionPlanChangeScheduledWebhookEventPropertiesJSON struct {
	ChangeDate     apijson.Field
	NewPlanID      apijson.Field
	PreviousPlanID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionPlanChangeScheduledWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPlanChangeScheduledWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type SubscriptionPlanChangeScheduledWebhookEventType string

const (
	SubscriptionPlanChangeScheduledWebhookEventTypeSubscriptionPlanChangeScheduled SubscriptionPlanChangeScheduledWebhookEventType = "subscription.plan_change_scheduled"
)

func (r SubscriptionPlanChangeScheduledWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionPlanChangeScheduledWebhookEventTypeSubscriptionPlanChangeScheduled:
		return true
	}
	return false
}

// Issued when a subscription transitions from one plan to a different plan.
type SubscriptionPlanChangedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                     `json:"created_at" api:"required" format:"date-time"`
	Properties SubscriptionPlanChangedWebhookEventProperties `json:"properties" api:"required"`
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
	Subscription Subscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionPlanChangedWebhookEventType `json:"type" api:"required"`
	JSON subscriptionPlanChangedWebhookEventJSON `json:"-"`
}

// subscriptionPlanChangedWebhookEventJSON contains the JSON metadata for the
// struct [SubscriptionPlanChangedWebhookEvent]
type subscriptionPlanChangedWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionPlanChangedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPlanChangedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionPlanChangedWebhookEvent) implementsUnwrapWebhookEvent() {}

type SubscriptionPlanChangedWebhookEventProperties struct {
	PreviousPlanID string                                            `json:"previous_plan_id" api:"required"`
	JSON           subscriptionPlanChangedWebhookEventPropertiesJSON `json:"-"`
}

// subscriptionPlanChangedWebhookEventPropertiesJSON contains the JSON metadata for
// the struct [SubscriptionPlanChangedWebhookEventProperties]
type subscriptionPlanChangedWebhookEventPropertiesJSON struct {
	PreviousPlanID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionPlanChangedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPlanChangedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type SubscriptionPlanChangedWebhookEventType string

const (
	SubscriptionPlanChangedWebhookEventTypeSubscriptionPlanChanged SubscriptionPlanChangedWebhookEventType = "subscription.plan_changed"
)

func (r SubscriptionPlanChangedWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionPlanChangedWebhookEventTypeSubscriptionPlanChanged:
		return true
	}
	return false
}

// Issued when a subscription plan version change is scheduled.
type SubscriptionPlanVersionChangeScheduledWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                                    `json:"created_at" api:"required" format:"date-time"`
	Properties SubscriptionPlanVersionChangeScheduledWebhookEventProperties `json:"properties" api:"required"`
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
	Subscription Subscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionPlanVersionChangeScheduledWebhookEventType `json:"type" api:"required"`
	JSON subscriptionPlanVersionChangeScheduledWebhookEventJSON `json:"-"`
}

// subscriptionPlanVersionChangeScheduledWebhookEventJSON contains the JSON
// metadata for the struct [SubscriptionPlanVersionChangeScheduledWebhookEvent]
type subscriptionPlanVersionChangeScheduledWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionPlanVersionChangeScheduledWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPlanVersionChangeScheduledWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionPlanVersionChangeScheduledWebhookEvent) implementsUnwrapWebhookEvent() {}

type SubscriptionPlanVersionChangeScheduledWebhookEventProperties struct {
	EffectiveDate             string                                                           `json:"effective_date" api:"required"`
	NewPlanVersionNumber      int64                                                            `json:"new_plan_version_number" api:"required"`
	PreviousPlanVersionNumber int64                                                            `json:"previous_plan_version_number" api:"required"`
	JSON                      subscriptionPlanVersionChangeScheduledWebhookEventPropertiesJSON `json:"-"`
}

// subscriptionPlanVersionChangeScheduledWebhookEventPropertiesJSON contains the
// JSON metadata for the struct
// [SubscriptionPlanVersionChangeScheduledWebhookEventProperties]
type subscriptionPlanVersionChangeScheduledWebhookEventPropertiesJSON struct {
	EffectiveDate             apijson.Field
	NewPlanVersionNumber      apijson.Field
	PreviousPlanVersionNumber apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionPlanVersionChangeScheduledWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPlanVersionChangeScheduledWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type SubscriptionPlanVersionChangeScheduledWebhookEventType string

const (
	SubscriptionPlanVersionChangeScheduledWebhookEventTypeSubscriptionPlanVersionChangeScheduled SubscriptionPlanVersionChangeScheduledWebhookEventType = "subscription.plan_version_change_scheduled"
)

func (r SubscriptionPlanVersionChangeScheduledWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionPlanVersionChangeScheduledWebhookEventTypeSubscriptionPlanVersionChangeScheduled:
		return true
	}
	return false
}

// Issued when a subscription plan version has changed.
type SubscriptionPlanVersionChangedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                            `json:"created_at" api:"required" format:"date-time"`
	Properties SubscriptionPlanVersionChangedWebhookEventProperties `json:"properties" api:"required"`
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
	Subscription Subscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionPlanVersionChangedWebhookEventType `json:"type" api:"required"`
	JSON subscriptionPlanVersionChangedWebhookEventJSON `json:"-"`
}

// subscriptionPlanVersionChangedWebhookEventJSON contains the JSON metadata for
// the struct [SubscriptionPlanVersionChangedWebhookEvent]
type subscriptionPlanVersionChangedWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionPlanVersionChangedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPlanVersionChangedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionPlanVersionChangedWebhookEvent) implementsUnwrapWebhookEvent() {}

type SubscriptionPlanVersionChangedWebhookEventProperties struct {
	EffectiveDate             string                                                   `json:"effective_date" api:"required"`
	NewPlanVersionNumber      int64                                                    `json:"new_plan_version_number" api:"required"`
	PreviousPlanVersionNumber int64                                                    `json:"previous_plan_version_number" api:"required"`
	JSON                      subscriptionPlanVersionChangedWebhookEventPropertiesJSON `json:"-"`
}

// subscriptionPlanVersionChangedWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [SubscriptionPlanVersionChangedWebhookEventProperties]
type subscriptionPlanVersionChangedWebhookEventPropertiesJSON struct {
	EffectiveDate             apijson.Field
	NewPlanVersionNumber      apijson.Field
	PreviousPlanVersionNumber apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionPlanVersionChangedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPlanVersionChangedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type SubscriptionPlanVersionChangedWebhookEventType string

const (
	SubscriptionPlanVersionChangedWebhookEventTypeSubscriptionPlanVersionChanged SubscriptionPlanVersionChangedWebhookEventType = "subscription.plan_version_changed"
)

func (r SubscriptionPlanVersionChangedWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionPlanVersionChangedWebhookEventTypeSubscriptionPlanVersionChanged:
		return true
	}
	return false
}

// Issued when a subscription's rated spend, before credits and adjustments,
// exceeds a pre-configured amount threshold.
type SubscriptionSpendExceededWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                       `json:"created_at" api:"required" format:"date-time"`
	Properties SubscriptionSpendExceededWebhookEventProperties `json:"properties" api:"required"`
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
	Subscription Subscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionSpendExceededWebhookEventType `json:"type" api:"required"`
	JSON subscriptionSpendExceededWebhookEventJSON `json:"-"`
}

// subscriptionSpendExceededWebhookEventJSON contains the JSON metadata for the
// struct [SubscriptionSpendExceededWebhookEvent]
type subscriptionSpendExceededWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionSpendExceededWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionSpendExceededWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionSpendExceededWebhookEvent) implementsUnwrapWebhookEvent() {}

type SubscriptionSpendExceededWebhookEventProperties struct {
	// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
	// usage, or credit balance and trigger webhooks when a threshold is exceeded.
	//
	// Alerts created through the API can be scoped to either customers or
	// subscriptions.
	AlertConfiguration SubscriptionSpendExceededWebhookEventPropertiesAlertConfiguration `json:"alert_configuration" api:"required"`
	AmountThreshold    string                                                            `json:"amount_threshold" api:"required,nullable"`
	EvaluatedAmount    string                                                            `json:"evaluated_amount" api:"required,nullable"`
	TimeframeEnd       time.Time                                                         `json:"timeframe_end" api:"required" format:"date-time"`
	TimeframeStart     time.Time                                                         `json:"timeframe_start" api:"required" format:"date-time"`
	JSON               subscriptionSpendExceededWebhookEventPropertiesJSON               `json:"-"`
}

// subscriptionSpendExceededWebhookEventPropertiesJSON contains the JSON metadata
// for the struct [SubscriptionSpendExceededWebhookEventProperties]
type subscriptionSpendExceededWebhookEventPropertiesJSON struct {
	AlertConfiguration apijson.Field
	AmountThreshold    apijson.Field
	EvaluatedAmount    apijson.Field
	TimeframeEnd       apijson.Field
	TimeframeStart     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionSpendExceededWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionSpendExceededWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
// usage, or credit balance and trigger webhooks when a threshold is exceeded.
//
// Alerts created through the API can be scoped to either customers or
// subscriptions.
type SubscriptionSpendExceededWebhookEventPropertiesAlertConfiguration struct {
	// Also referred to as alert_id in this documentation.
	ID string `json:"id" api:"required"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The name of the currency the credit balance or invoice cost is denominated in.
	Currency string `json:"currency" api:"required,nullable"`
	// The customer the alert applies to.
	Customer shared.CustomerMinified `json:"customer" api:"required,nullable"`
	// Whether the alert is enabled or disabled.
	Enabled bool `json:"enabled" api:"required"`
	// The metric the alert applies to.
	Metric SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationMetric `json:"metric" api:"required,nullable"`
	// The plan the alert applies to.
	Plan SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPlan `json:"plan" api:"required,nullable"`
	// The subscription the alert applies to.
	Subscription shared.SubscriptionMinified `json:"subscription" api:"required,nullable"`
	// The thresholds that define the conditions under which the alert will be
	// triggered.
	Thresholds []SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThreshold `json:"thresholds" api:"required,nullable"`
	// The type of alert. This must be a valid alert type.
	Type SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationType `json:"type" api:"required"`
	// The current status of the alert. This field is only present for credit balance
	// alerts.
	BalanceAlertStatus []SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatus `json:"balance_alert_status" api:"nullable"`
	// The property keys to group cost alerts by. Only present for cost alerts with
	// grouping enabled.
	GroupingKeys []string `json:"grouping_keys" api:"nullable"`
	// Minified license type for alert serialization.
	LicenseType SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationLicenseType `json:"license_type" api:"nullable"`
	// Filters scoping which prices are included in spend and grouped cost alert
	// evaluation. Alerts use the price_id, item_id, and price_type fields only; the
	// alert's pricing unit is reported by currency.
	PriceFilters []SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFilter `json:"price_filters" api:"nullable"`
	// Per-group threshold overrides. Each override maps a specific combination of
	// grouping_keys values to a replacement threshold list. Only present for grouped
	// cost alerts that have at least one override.
	ThresholdOverrides []SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdOverride `json:"threshold_overrides" api:"nullable"`
	JSON               subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationJSON                `json:"-"`
}

// subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationJSON contains
// the JSON metadata for the struct
// [SubscriptionSpendExceededWebhookEventPropertiesAlertConfiguration]
type subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Customer           apijson.Field
	Enabled            apijson.Field
	Metric             apijson.Field
	Plan               apijson.Field
	Subscription       apijson.Field
	Thresholds         apijson.Field
	Type               apijson.Field
	BalanceAlertStatus apijson.Field
	GroupingKeys       apijson.Field
	LicenseType        apijson.Field
	PriceFilters       apijson.Field
	ThresholdOverrides apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionSpendExceededWebhookEventPropertiesAlertConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationJSON) RawJSON() string {
	return r.raw
}

// The metric the alert applies to.
type SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationMetric struct {
	ID   string                                                                      `json:"id" api:"required"`
	JSON subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationMetricJSON `json:"-"`
}

// subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationMetricJSON
// contains the JSON metadata for the struct
// [SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationMetric]
type subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationMetricJSON) RawJSON() string {
	return r.raw
}

// The plan the alert applies to.
type SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPlan struct {
	ID string `json:"id" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                                                                    `json:"external_plan_id" api:"required,nullable"`
	Name           string                                                                    `json:"name" api:"required,nullable"`
	PlanVersion    string                                                                    `json:"plan_version" api:"required"`
	JSON           subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPlanJSON `json:"-"`
}

// subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPlanJSON
// contains the JSON metadata for the struct
// [SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPlan]
type subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	PlanVersion    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPlanJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                         `json:"value" api:"required"`
	JSON  subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdJSON `json:"-"`
}

// subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdJSON
// contains the JSON metadata for the struct
// [SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThreshold]
type subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdJSON) RawJSON() string {
	return r.raw
}

// The type of alert. This must be a valid alert type.
type SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationType string

const (
	SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted          SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationType = "credit_balance_depleted"
	SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped           SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationType = "credit_balance_dropped"
	SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered         SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationType = "credit_balance_recovered"
	SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationTypeUsageExceeded                  SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationType = "usage_exceeded"
	SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationTypeCostExceeded                   SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationType = "cost_exceeded"
	SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationTypeSpendExceeded                  SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationType = "spend_exceeded"
	SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationType = "license_balance_threshold_reached"
)

func (r SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationType) IsKnown() bool {
	switch r {
	case SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted, SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped, SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered, SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationTypeUsageExceeded, SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationTypeCostExceeded, SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationTypeSpendExceeded, SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached:
		return true
	}
	return false
}

// Alert status is used to determine if an alert is currently in-alert or not.
type SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatus struct {
	// Whether the alert is currently in-alert or not.
	InAlert bool `json:"in_alert" api:"required"`
	// The value of the threshold that defines the alert status.
	ThresholdValue string                                                                                  `json:"threshold_value" api:"required"`
	JSON           subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON `json:"-"`
}

// subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON
// contains the JSON metadata for the struct
// [SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatus]
type subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON struct {
	InAlert        apijson.Field
	ThresholdValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON) RawJSON() string {
	return r.raw
}

// Minified license type for alert serialization.
type SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationLicenseType struct {
	ID   string                                                                           `json:"id" api:"required"`
	JSON subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationLicenseTypeJSON `json:"-"`
}

// subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationLicenseTypeJSON
// contains the JSON metadata for the struct
// [SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationLicenseType]
type subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationLicenseTypeJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationLicenseType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationLicenseTypeJSON) RawJSON() string {
	return r.raw
}

type SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFilter struct {
	// The property of the price to filter on.
	Field SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                                         `json:"values" api:"required"`
	JSON   subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFilterJSON `json:"-"`
}

// subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFilter]
type subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField string

const (
	SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID       SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_id"
	SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID        SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "item_id"
	SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType     SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_type"
	SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency      SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "currency"
	SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "pricing_unit_id"
)

func (r SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID, SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID, SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType, SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency, SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator string

const (
	SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "includes"
	SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "excludes"
)

func (r SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes, SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes:
		return true
	}
	return false
}

// A per-group threshold override on a grouped cost alert.
//
// An empty `thresholds` list means the group is silenced (never fires). A
// non-empty list fully replaces the default thresholds for that group.
type SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdOverride struct {
	// The values identifying this group, ordered to match group_keys when set and the
	// alert's grouping_keys otherwise.
	GroupValues []string `json:"group_values" api:"required"`
	// The thresholds applied to this group. An empty list means the group is silenced.
	Thresholds []SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold `json:"thresholds" api:"required"`
	// The subset of the alert's grouping_keys this override binds. Null when the
	// override targets one exact group across every grouping key.
	GroupKeys []string                                                                               `json:"group_keys" api:"nullable"`
	JSON      subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON `json:"-"`
}

// subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON
// contains the JSON metadata for the struct
// [SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdOverride]
type subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON struct {
	GroupValues apijson.Field
	Thresholds  apijson.Field
	GroupKeys   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                                           `json:"value" api:"required"`
	JSON  subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON `json:"-"`
}

// subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON
// contains the JSON metadata for the struct
// [SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold]
type subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionSpendExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type SubscriptionSpendExceededWebhookEventType string

const (
	SubscriptionSpendExceededWebhookEventTypeSubscriptionSpendExceeded SubscriptionSpendExceededWebhookEventType = "subscription.spend_exceeded"
)

func (r SubscriptionSpendExceededWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionSpendExceededWebhookEventTypeSubscriptionSpendExceeded:
		return true
	}
	return false
}

// Issued when a subscription begins.
type SubscriptionStartedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time   `json:"created_at" api:"required" format:"date-time"`
	Properties interface{} `json:"properties" api:"required"`
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
	Subscription Subscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionStartedWebhookEventType `json:"type" api:"required"`
	JSON subscriptionStartedWebhookEventJSON `json:"-"`
}

// subscriptionStartedWebhookEventJSON contains the JSON metadata for the struct
// [SubscriptionStartedWebhookEvent]
type subscriptionStartedWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionStartedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionStartedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionStartedWebhookEvent) implementsUnwrapWebhookEvent() {}

// The event this payload describes.
type SubscriptionStartedWebhookEventType string

const (
	SubscriptionStartedWebhookEventTypeSubscriptionStarted SubscriptionStartedWebhookEventType = "subscription.started"
)

func (r SubscriptionStartedWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionStartedWebhookEventTypeSubscriptionStarted:
		return true
	}
	return false
}

// Issued when a subscription trial ends.
type SubscriptionTrialEndedWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time   `json:"created_at" api:"required" format:"date-time"`
	Properties interface{} `json:"properties" api:"required"`
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
	Subscription Subscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionTrialEndedWebhookEventType `json:"type" api:"required"`
	JSON subscriptionTrialEndedWebhookEventJSON `json:"-"`
}

// subscriptionTrialEndedWebhookEventJSON contains the JSON metadata for the struct
// [SubscriptionTrialEndedWebhookEvent]
type subscriptionTrialEndedWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionTrialEndedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionTrialEndedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionTrialEndedWebhookEvent) implementsUnwrapWebhookEvent() {}

// The event this payload describes.
type SubscriptionTrialEndedWebhookEventType string

const (
	SubscriptionTrialEndedWebhookEventTypeSubscriptionTrialEnded SubscriptionTrialEndedWebhookEventType = "subscription.trial_ended"
)

func (r SubscriptionTrialEndedWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionTrialEndedWebhookEventTypeSubscriptionTrialEnded:
		return true
	}
	return false
}

// Issued when a billable metric in a subscription exceeds a pre-configured
// quantity threshold.
type SubscriptionUsageExceededWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt  time.Time                                       `json:"created_at" api:"required" format:"date-time"`
	Properties SubscriptionUsageExceededWebhookEventProperties `json:"properties" api:"required"`
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
	Subscription Subscription `json:"subscription" api:"required"`
	// The event this payload describes.
	Type SubscriptionUsageExceededWebhookEventType `json:"type" api:"required"`
	JSON subscriptionUsageExceededWebhookEventJSON `json:"-"`
}

// subscriptionUsageExceededWebhookEventJSON contains the JSON metadata for the
// struct [SubscriptionUsageExceededWebhookEvent]
type subscriptionUsageExceededWebhookEventJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Properties   apijson.Field
	Subscription apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SubscriptionUsageExceededWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageExceededWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionUsageExceededWebhookEvent) implementsUnwrapWebhookEvent() {}

type SubscriptionUsageExceededWebhookEventProperties struct {
	// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
	// usage, or credit balance and trigger webhooks when a threshold is exceeded.
	//
	// Alerts created through the API can be scoped to either customers or
	// subscriptions.
	AlertConfiguration SubscriptionUsageExceededWebhookEventPropertiesAlertConfiguration `json:"alert_configuration" api:"required"`
	BillableMetricID   string                                                            `json:"billable_metric_id" api:"required"`
	EvaluatedQuantity  float64                                                           `json:"evaluated_quantity" api:"required,nullable"`
	QuantityThreshold  float64                                                           `json:"quantity_threshold" api:"required,nullable"`
	TimeframeEnd       time.Time                                                         `json:"timeframe_end" api:"required" format:"date-time"`
	TimeframeStart     time.Time                                                         `json:"timeframe_start" api:"required" format:"date-time"`
	JSON               subscriptionUsageExceededWebhookEventPropertiesJSON               `json:"-"`
}

// subscriptionUsageExceededWebhookEventPropertiesJSON contains the JSON metadata
// for the struct [SubscriptionUsageExceededWebhookEventProperties]
type subscriptionUsageExceededWebhookEventPropertiesJSON struct {
	AlertConfiguration apijson.Field
	BillableMetricID   apijson.Field
	EvaluatedQuantity  apijson.Field
	QuantityThreshold  apijson.Field
	TimeframeEnd       apijson.Field
	TimeframeStart     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionUsageExceededWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageExceededWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
// usage, or credit balance and trigger webhooks when a threshold is exceeded.
//
// Alerts created through the API can be scoped to either customers or
// subscriptions.
type SubscriptionUsageExceededWebhookEventPropertiesAlertConfiguration struct {
	// Also referred to as alert_id in this documentation.
	ID string `json:"id" api:"required"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The name of the currency the credit balance or invoice cost is denominated in.
	Currency string `json:"currency" api:"required,nullable"`
	// The customer the alert applies to.
	Customer shared.CustomerMinified `json:"customer" api:"required,nullable"`
	// Whether the alert is enabled or disabled.
	Enabled bool `json:"enabled" api:"required"`
	// The metric the alert applies to.
	Metric SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationMetric `json:"metric" api:"required,nullable"`
	// The plan the alert applies to.
	Plan SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPlan `json:"plan" api:"required,nullable"`
	// The subscription the alert applies to.
	Subscription shared.SubscriptionMinified `json:"subscription" api:"required,nullable"`
	// The thresholds that define the conditions under which the alert will be
	// triggered.
	Thresholds []SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThreshold `json:"thresholds" api:"required,nullable"`
	// The type of alert. This must be a valid alert type.
	Type SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationType `json:"type" api:"required"`
	// The current status of the alert. This field is only present for credit balance
	// alerts.
	BalanceAlertStatus []SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatus `json:"balance_alert_status" api:"nullable"`
	// The property keys to group cost alerts by. Only present for cost alerts with
	// grouping enabled.
	GroupingKeys []string `json:"grouping_keys" api:"nullable"`
	// Minified license type for alert serialization.
	LicenseType SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationLicenseType `json:"license_type" api:"nullable"`
	// Filters scoping which prices are included in spend and grouped cost alert
	// evaluation. Alerts use the price_id, item_id, and price_type fields only; the
	// alert's pricing unit is reported by currency.
	PriceFilters []SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFilter `json:"price_filters" api:"nullable"`
	// Per-group threshold overrides. Each override maps a specific combination of
	// grouping_keys values to a replacement threshold list. Only present for grouped
	// cost alerts that have at least one override.
	ThresholdOverrides []SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdOverride `json:"threshold_overrides" api:"nullable"`
	JSON               subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationJSON                `json:"-"`
}

// subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationJSON contains
// the JSON metadata for the struct
// [SubscriptionUsageExceededWebhookEventPropertiesAlertConfiguration]
type subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Customer           apijson.Field
	Enabled            apijson.Field
	Metric             apijson.Field
	Plan               apijson.Field
	Subscription       apijson.Field
	Thresholds         apijson.Field
	Type               apijson.Field
	BalanceAlertStatus apijson.Field
	GroupingKeys       apijson.Field
	LicenseType        apijson.Field
	PriceFilters       apijson.Field
	ThresholdOverrides apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionUsageExceededWebhookEventPropertiesAlertConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationJSON) RawJSON() string {
	return r.raw
}

// The metric the alert applies to.
type SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationMetric struct {
	ID   string                                                                      `json:"id" api:"required"`
	JSON subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationMetricJSON `json:"-"`
}

// subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationMetricJSON
// contains the JSON metadata for the struct
// [SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationMetric]
type subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationMetricJSON) RawJSON() string {
	return r.raw
}

// The plan the alert applies to.
type SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPlan struct {
	ID string `json:"id" api:"required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                                                                    `json:"external_plan_id" api:"required,nullable"`
	Name           string                                                                    `json:"name" api:"required,nullable"`
	PlanVersion    string                                                                    `json:"plan_version" api:"required"`
	JSON           subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPlanJSON `json:"-"`
}

// subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPlanJSON
// contains the JSON metadata for the struct
// [SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPlan]
type subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	PlanVersion    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPlanJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                         `json:"value" api:"required"`
	JSON  subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdJSON `json:"-"`
}

// subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdJSON
// contains the JSON metadata for the struct
// [SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThreshold]
type subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdJSON) RawJSON() string {
	return r.raw
}

// The type of alert. This must be a valid alert type.
type SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationType string

const (
	SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted          SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationType = "credit_balance_depleted"
	SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped           SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationType = "credit_balance_dropped"
	SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered         SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationType = "credit_balance_recovered"
	SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationTypeUsageExceeded                  SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationType = "usage_exceeded"
	SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationTypeCostExceeded                   SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationType = "cost_exceeded"
	SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationTypeSpendExceeded                  SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationType = "spend_exceeded"
	SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationType = "license_balance_threshold_reached"
)

func (r SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationType) IsKnown() bool {
	switch r {
	case SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDepleted, SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceDropped, SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationTypeCreditBalanceRecovered, SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationTypeUsageExceeded, SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationTypeCostExceeded, SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationTypeSpendExceeded, SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationTypeLicenseBalanceThresholdReached:
		return true
	}
	return false
}

// Alert status is used to determine if an alert is currently in-alert or not.
type SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatus struct {
	// Whether the alert is currently in-alert or not.
	InAlert bool `json:"in_alert" api:"required"`
	// The value of the threshold that defines the alert status.
	ThresholdValue string                                                                                  `json:"threshold_value" api:"required"`
	JSON           subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON `json:"-"`
}

// subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON
// contains the JSON metadata for the struct
// [SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatus]
type subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON struct {
	InAlert        apijson.Field
	ThresholdValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationBalanceAlertStatusJSON) RawJSON() string {
	return r.raw
}

// Minified license type for alert serialization.
type SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationLicenseType struct {
	ID   string                                                                           `json:"id" api:"required"`
	JSON subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationLicenseTypeJSON `json:"-"`
}

// subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationLicenseTypeJSON
// contains the JSON metadata for the struct
// [SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationLicenseType]
type subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationLicenseTypeJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationLicenseType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationLicenseTypeJSON) RawJSON() string {
	return r.raw
}

type SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFilter struct {
	// The property of the price to filter on.
	Field SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                                         `json:"values" api:"required"`
	JSON   subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFilterJSON `json:"-"`
}

// subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFilter]
type subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField string

const (
	SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID       SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_id"
	SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID        SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "item_id"
	SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType     SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "price_type"
	SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency      SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "currency"
	SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField = "pricing_unit_id"
)

func (r SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceID, SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldItemID, SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPriceType, SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldCurrency, SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator string

const (
	SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "includes"
	SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator = "excludes"
)

func (r SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorIncludes, SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationPriceFiltersOperatorExcludes:
		return true
	}
	return false
}

// A per-group threshold override on a grouped cost alert.
//
// An empty `thresholds` list means the group is silenced (never fires). A
// non-empty list fully replaces the default thresholds for that group.
type SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdOverride struct {
	// The values identifying this group, ordered to match group_keys when set and the
	// alert's grouping_keys otherwise.
	GroupValues []string `json:"group_values" api:"required"`
	// The thresholds applied to this group. An empty list means the group is silenced.
	Thresholds []SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold `json:"thresholds" api:"required"`
	// The subset of the alert's grouping_keys this override binds. Null when the
	// override targets one exact group across every grouping key.
	GroupKeys []string                                                                               `json:"group_keys" api:"nullable"`
	JSON      subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON `json:"-"`
}

// subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON
// contains the JSON metadata for the struct
// [SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdOverride]
type subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON struct {
	GroupValues apijson.Field
	Thresholds  apijson.Field
	GroupKeys   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdOverrideJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value string                                                                                           `json:"value" api:"required"`
	JSON  subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON `json:"-"`
}

// subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON
// contains the JSON metadata for the struct
// [SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold]
type subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageExceededWebhookEventPropertiesAlertConfigurationThresholdOverridesThresholdJSON) RawJSON() string {
	return r.raw
}

// The event this payload describes.
type SubscriptionUsageExceededWebhookEventType string

const (
	SubscriptionUsageExceededWebhookEventTypeSubscriptionUsageExceeded SubscriptionUsageExceededWebhookEventType = "subscription.usage_exceeded"
)

func (r SubscriptionUsageExceededWebhookEventType) IsKnown() bool {
	switch r {
	case SubscriptionUsageExceededWebhookEventTypeSubscriptionUsageExceeded:
		return true
	}
	return false
}

// Issued when a transaction accounting sync fails.
type TransactionAccountingSyncFailedWebhookEvent struct {
	// The ID of this webhook event.
	ID                   string                                                          `json:"id" api:"required"`
	AccountingSyncRecord TransactionAccountingSyncFailedWebhookEventAccountingSyncRecord `json:"accounting_sync_record" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt   time.Time                                              `json:"created_at" api:"required" format:"date-time"`
	Properties  TransactionAccountingSyncFailedWebhookEventProperties  `json:"properties" api:"required"`
	Transaction TransactionAccountingSyncFailedWebhookEventTransaction `json:"transaction" api:"required"`
	// The event this payload describes.
	Type TransactionAccountingSyncFailedWebhookEventType `json:"type" api:"required"`
	JSON transactionAccountingSyncFailedWebhookEventJSON `json:"-"`
}

// transactionAccountingSyncFailedWebhookEventJSON contains the JSON metadata for
// the struct [TransactionAccountingSyncFailedWebhookEvent]
type transactionAccountingSyncFailedWebhookEventJSON struct {
	ID                   apijson.Field
	AccountingSyncRecord apijson.Field
	CreatedAt            apijson.Field
	Properties           apijson.Field
	Transaction          apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TransactionAccountingSyncFailedWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAccountingSyncFailedWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r TransactionAccountingSyncFailedWebhookEvent) implementsUnwrapWebhookEvent() {}

type TransactionAccountingSyncFailedWebhookEventAccountingSyncRecord struct {
	ID                  string                                                                    `json:"id" api:"required"`
	CustomerID          string                                                                    `json:"customer_id" api:"required"`
	RecordType          TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType `json:"record_type" api:"required"`
	ErrorDetails        map[string]interface{}                                                    `json:"error_details" api:"nullable"`
	ProviderCustomerID  string                                                                    `json:"provider_customer_id" api:"nullable"`
	Status              string                                                                    `json:"status" api:"nullable"`
	SyncAction          string                                                                    `json:"sync_action" api:"nullable"`
	TransactionRecordID string                                                                    `json:"transaction_record_id" api:"nullable"`
	JSON                transactionAccountingSyncFailedWebhookEventAccountingSyncRecordJSON       `json:"-"`
}

// transactionAccountingSyncFailedWebhookEventAccountingSyncRecordJSON contains the
// JSON metadata for the struct
// [TransactionAccountingSyncFailedWebhookEventAccountingSyncRecord]
type transactionAccountingSyncFailedWebhookEventAccountingSyncRecordJSON struct {
	ID                  apijson.Field
	CustomerID          apijson.Field
	RecordType          apijson.Field
	ErrorDetails        apijson.Field
	ProviderCustomerID  apijson.Field
	Status              apijson.Field
	SyncAction          apijson.Field
	TransactionRecordID apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TransactionAccountingSyncFailedWebhookEventAccountingSyncRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAccountingSyncFailedWebhookEventAccountingSyncRecordJSON) RawJSON() string {
	return r.raw
}

type TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType string

const (
	TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomer                   TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "customer"
	TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeInvoice                    TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "invoice"
	TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeTransaction                TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "transaction"
	TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "customer_balance_transaction"
	TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCreditNote                 TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "credit_note"
	TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSubscription               TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "subscription"
	TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSalesOrder                 TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "sales_order"
	TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeBlock                      TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType = "block"
)

func (r TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordType) IsKnown() bool {
	switch r {
	case TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomer, TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeInvoice, TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeTransaction, TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction, TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeCreditNote, TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSubscription, TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeSalesOrder, TransactionAccountingSyncFailedWebhookEventAccountingSyncRecordRecordTypeBlock:
		return true
	}
	return false
}

type TransactionAccountingSyncFailedWebhookEventProperties struct {
	ConnectionType string                                                    `json:"connection_type" api:"required"`
	FailureReason  string                                                    `json:"failure_reason" api:"required"`
	JSON           transactionAccountingSyncFailedWebhookEventPropertiesJSON `json:"-"`
}

// transactionAccountingSyncFailedWebhookEventPropertiesJSON contains the JSON
// metadata for the struct [TransactionAccountingSyncFailedWebhookEventProperties]
type transactionAccountingSyncFailedWebhookEventPropertiesJSON struct {
	ConnectionType apijson.Field
	FailureReason  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionAccountingSyncFailedWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAccountingSyncFailedWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

type TransactionAccountingSyncFailedWebhookEventTransaction struct {
	// The ID of the payment attempt.
	ID string `json:"id" api:"required"`
	// The amount of the payment attempt.
	Amount string `json:"amount" api:"required"`
	// The time at which the payment attempt was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The payment provider that attempted to collect the payment.
	PaymentProvider TransactionAccountingSyncFailedWebhookEventTransactionPaymentProvider `json:"payment_provider" api:"required,nullable"`
	// The ID of the payment attempt in the payment provider.
	PaymentProviderID string `json:"payment_provider_id" api:"required,nullable"`
	// URL to the downloadable PDF version of the receipt. This field will be `null`
	// for payment attempts that did not succeed.
	ReceiptPdf string `json:"receipt_pdf" api:"required,nullable"`
	// Whether the payment attempt succeeded.
	Succeeded bool                                                       `json:"succeeded" api:"required"`
	JSON      transactionAccountingSyncFailedWebhookEventTransactionJSON `json:"-"`
}

// transactionAccountingSyncFailedWebhookEventTransactionJSON contains the JSON
// metadata for the struct [TransactionAccountingSyncFailedWebhookEventTransaction]
type transactionAccountingSyncFailedWebhookEventTransactionJSON struct {
	ID                apijson.Field
	Amount            apijson.Field
	CreatedAt         apijson.Field
	PaymentProvider   apijson.Field
	PaymentProviderID apijson.Field
	ReceiptPdf        apijson.Field
	Succeeded         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TransactionAccountingSyncFailedWebhookEventTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAccountingSyncFailedWebhookEventTransactionJSON) RawJSON() string {
	return r.raw
}

// The payment provider that attempted to collect the payment.
type TransactionAccountingSyncFailedWebhookEventTransactionPaymentProvider string

const (
	TransactionAccountingSyncFailedWebhookEventTransactionPaymentProviderStripe TransactionAccountingSyncFailedWebhookEventTransactionPaymentProvider = "stripe"
	TransactionAccountingSyncFailedWebhookEventTransactionPaymentProviderAdyen  TransactionAccountingSyncFailedWebhookEventTransactionPaymentProvider = "adyen"
)

func (r TransactionAccountingSyncFailedWebhookEventTransactionPaymentProvider) IsKnown() bool {
	switch r {
	case TransactionAccountingSyncFailedWebhookEventTransactionPaymentProviderStripe, TransactionAccountingSyncFailedWebhookEventTransactionPaymentProviderAdyen:
		return true
	}
	return false
}

// The event this payload describes.
type TransactionAccountingSyncFailedWebhookEventType string

const (
	TransactionAccountingSyncFailedWebhookEventTypeTransactionAccountingSyncFailed TransactionAccountingSyncFailedWebhookEventType = "transaction.accounting_sync_failed"
)

func (r TransactionAccountingSyncFailedWebhookEventType) IsKnown() bool {
	switch r {
	case TransactionAccountingSyncFailedWebhookEventTypeTransactionAccountingSyncFailed:
		return true
	}
	return false
}

// Issued when a transaction accounting sync succeeds.
type TransactionAccountingSyncSucceededWebhookEvent struct {
	// The ID of this webhook event.
	ID                   string                                                             `json:"id" api:"required"`
	AccountingSyncRecord TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecord `json:"accounting_sync_record" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt   time.Time                                                 `json:"created_at" api:"required" format:"date-time"`
	Properties  TransactionAccountingSyncSucceededWebhookEventProperties  `json:"properties" api:"required"`
	Transaction TransactionAccountingSyncSucceededWebhookEventTransaction `json:"transaction" api:"required"`
	// The event this payload describes.
	Type TransactionAccountingSyncSucceededWebhookEventType `json:"type" api:"required"`
	JSON transactionAccountingSyncSucceededWebhookEventJSON `json:"-"`
}

// transactionAccountingSyncSucceededWebhookEventJSON contains the JSON metadata
// for the struct [TransactionAccountingSyncSucceededWebhookEvent]
type transactionAccountingSyncSucceededWebhookEventJSON struct {
	ID                   apijson.Field
	AccountingSyncRecord apijson.Field
	CreatedAt            apijson.Field
	Properties           apijson.Field
	Transaction          apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TransactionAccountingSyncSucceededWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAccountingSyncSucceededWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r TransactionAccountingSyncSucceededWebhookEvent) implementsUnwrapWebhookEvent() {}

type TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecord struct {
	ID                  string                                                                       `json:"id" api:"required"`
	CustomerID          string                                                                       `json:"customer_id" api:"required"`
	RecordType          TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType `json:"record_type" api:"required"`
	ErrorDetails        map[string]interface{}                                                       `json:"error_details" api:"nullable"`
	ProviderCustomerID  string                                                                       `json:"provider_customer_id" api:"nullable"`
	Status              string                                                                       `json:"status" api:"nullable"`
	SyncAction          string                                                                       `json:"sync_action" api:"nullable"`
	TransactionRecordID string                                                                       `json:"transaction_record_id" api:"nullable"`
	JSON                transactionAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON       `json:"-"`
}

// transactionAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON contains
// the JSON metadata for the struct
// [TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecord]
type transactionAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON struct {
	ID                  apijson.Field
	CustomerID          apijson.Field
	RecordType          apijson.Field
	ErrorDetails        apijson.Field
	ProviderCustomerID  apijson.Field
	Status              apijson.Field
	SyncAction          apijson.Field
	TransactionRecordID apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAccountingSyncSucceededWebhookEventAccountingSyncRecordJSON) RawJSON() string {
	return r.raw
}

type TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType string

const (
	TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomer                   TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "customer"
	TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeInvoice                    TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "invoice"
	TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeTransaction                TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "transaction"
	TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "customer_balance_transaction"
	TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCreditNote                 TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "credit_note"
	TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSubscription               TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "subscription"
	TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSalesOrder                 TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "sales_order"
	TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeBlock                      TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType = "block"
)

func (r TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordType) IsKnown() bool {
	switch r {
	case TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomer, TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeInvoice, TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeTransaction, TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCustomerBalanceTransaction, TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeCreditNote, TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSubscription, TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeSalesOrder, TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecordRecordTypeBlock:
		return true
	}
	return false
}

type TransactionAccountingSyncSucceededWebhookEventProperties struct {
	ConnectionType string                                                       `json:"connection_type" api:"required"`
	JSON           transactionAccountingSyncSucceededWebhookEventPropertiesJSON `json:"-"`
}

// transactionAccountingSyncSucceededWebhookEventPropertiesJSON contains the JSON
// metadata for the struct
// [TransactionAccountingSyncSucceededWebhookEventProperties]
type transactionAccountingSyncSucceededWebhookEventPropertiesJSON struct {
	ConnectionType apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionAccountingSyncSucceededWebhookEventProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAccountingSyncSucceededWebhookEventPropertiesJSON) RawJSON() string {
	return r.raw
}

type TransactionAccountingSyncSucceededWebhookEventTransaction struct {
	// The ID of the payment attempt.
	ID string `json:"id" api:"required"`
	// The amount of the payment attempt.
	Amount string `json:"amount" api:"required"`
	// The time at which the payment attempt was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The payment provider that attempted to collect the payment.
	PaymentProvider TransactionAccountingSyncSucceededWebhookEventTransactionPaymentProvider `json:"payment_provider" api:"required,nullable"`
	// The ID of the payment attempt in the payment provider.
	PaymentProviderID string `json:"payment_provider_id" api:"required,nullable"`
	// URL to the downloadable PDF version of the receipt. This field will be `null`
	// for payment attempts that did not succeed.
	ReceiptPdf string `json:"receipt_pdf" api:"required,nullable"`
	// Whether the payment attempt succeeded.
	Succeeded bool                                                          `json:"succeeded" api:"required"`
	JSON      transactionAccountingSyncSucceededWebhookEventTransactionJSON `json:"-"`
}

// transactionAccountingSyncSucceededWebhookEventTransactionJSON contains the JSON
// metadata for the struct
// [TransactionAccountingSyncSucceededWebhookEventTransaction]
type transactionAccountingSyncSucceededWebhookEventTransactionJSON struct {
	ID                apijson.Field
	Amount            apijson.Field
	CreatedAt         apijson.Field
	PaymentProvider   apijson.Field
	PaymentProviderID apijson.Field
	ReceiptPdf        apijson.Field
	Succeeded         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TransactionAccountingSyncSucceededWebhookEventTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionAccountingSyncSucceededWebhookEventTransactionJSON) RawJSON() string {
	return r.raw
}

// The payment provider that attempted to collect the payment.
type TransactionAccountingSyncSucceededWebhookEventTransactionPaymentProvider string

const (
	TransactionAccountingSyncSucceededWebhookEventTransactionPaymentProviderStripe TransactionAccountingSyncSucceededWebhookEventTransactionPaymentProvider = "stripe"
	TransactionAccountingSyncSucceededWebhookEventTransactionPaymentProviderAdyen  TransactionAccountingSyncSucceededWebhookEventTransactionPaymentProvider = "adyen"
)

func (r TransactionAccountingSyncSucceededWebhookEventTransactionPaymentProvider) IsKnown() bool {
	switch r {
	case TransactionAccountingSyncSucceededWebhookEventTransactionPaymentProviderStripe, TransactionAccountingSyncSucceededWebhookEventTransactionPaymentProviderAdyen:
		return true
	}
	return false
}

// The event this payload describes.
type TransactionAccountingSyncSucceededWebhookEventType string

const (
	TransactionAccountingSyncSucceededWebhookEventTypeTransactionAccountingSyncSucceeded TransactionAccountingSyncSucceededWebhookEventType = "transaction.accounting_sync_succeeded"
)

func (r TransactionAccountingSyncSucceededWebhookEventType) IsKnown() bool {
	switch r {
	case TransactionAccountingSyncSucceededWebhookEventTypeTransactionAccountingSyncSucceeded:
		return true
	}
	return false
}

// Issued when a backfill is closed and its events are reflected into usage.
type UnwrapWebhookEvent struct {
	// The ID of this webhook event.
	ID string `json:"id" api:"required"`
	// The time at which this event was created, to the second.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// This field can have the runtime type of [interface{}],
	// [BillableMetricEditedWebhookEventProperties],
	// [CreditBlockAccountingSyncFailedWebhookEventProperties],
	// [CreditBlockAccountingSyncSucceededWebhookEventProperties],
	// [CreditNoteAccountingSyncFailedWebhookEventProperties],
	// [CreditNoteAccountingSyncSucceededWebhookEventProperties],
	// [CustomerAccountingSyncFailedWebhookEventProperties],
	// [CustomerAccountingSyncSucceededWebhookEventProperties],
	// [CustomerBalanceTransactionCreatedWebhookEventProperties],
	// [CustomerCreditBalanceDepletedWebhookEventProperties],
	// [CustomerCreditBalanceDroppedWebhookEventProperties],
	// [CustomerCreditBalanceRecoveredWebhookEventProperties],
	// [CustomerCreditLedgerIncrementedWebhookEventProperties],
	// [CustomerEditedWebhookEventProperties],
	// [DataExportsTransferErrorWebhookEventProperties],
	// [DataExportsTransferSuccessWebhookEventProperties],
	// [EventUnmatchedEventWebhookEventProperties],
	// [IngestionUnmatchedCustomerIDsWebhookEventProperties],
	// [InvoiceAccountingSyncFailedWebhookEventProperties],
	// [InvoiceAccountingSyncSucceededWebhookEventProperties],
	// [InvoiceAutomationScheduleStepExecutedWebhookEventProperties],
	// [InvoiceCostDataExportedWebhookEventProperties],
	// [InvoiceDunningScheduleCreatedWebhookEventProperties],
	// [InvoiceDunningScheduleEndedWebhookEventProperties],
	// [InvoiceDunningScheduleResetWebhookEventProperties],
	// [InvoiceDunningScheduleStepExecutedWebhookEventProperties],
	// [InvoiceEditedWebhookEventProperties],
	// [InvoiceInvoiceDateElapsedWebhookEventProperties],
	// [InvoiceIssueFailedWebhookEventProperties],
	// [InvoiceIssuedWebhookEventProperties],
	// [InvoiceIssuedSummaryWebhookEventProperties],
	// [InvoiceManuallyMarkedAsPaidWebhookEventProperties],
	// [InvoicePaymentFailedWebhookEventProperties],
	// [InvoicePaymentProcessingWebhookEventProperties],
	// [InvoicePaymentSucceededWebhookEventProperties],
	// [InvoiceSyncFailedWebhookEventProperties],
	// [InvoiceSyncSucceededWebhookEventProperties],
	// [InvoiceDueDateRecalculationCanceledWebhookEventProperties],
	// [InvoiceDueDateRecalculationCompletedWebhookEventProperties],
	// [InvoiceDueDateRecalculationStartedWebhookEventProperties],
	// [MetricEventsDroppedByWatermarkWebhookEventProperties],
	// [PlanDefaultVersionSetWebhookEventProperties],
	// [PlanVersionCreatedWebhookEventProperties], [PriceEditedWebhookEventProperties],
	// [ResourceEventTestWebhookEventProperties],
	// [SalesOrderAccountingSyncFailedWebhookEventProperties],
	// [SalesOrderAccountingSyncSucceededWebhookEventProperties],
	// [SubscriptionAccountingSyncFailedWebhookEventProperties],
	// [SubscriptionAccountingSyncSucceededWebhookEventProperties],
	// [SubscriptionAlertDisabledWebhookEventProperties],
	// [SubscriptionCancellationScheduledWebhookEventProperties],
	// [SubscriptionCancellationUnscheduledWebhookEventProperties],
	// [SubscriptionCostExceededWebhookEventProperties],
	// [SubscriptionEditedWebhookEventProperties],
	// [SubscriptionFixedFeeQuantityUpdatedWebhookEventProperties],
	// [SubscriptionGroupedCostExceededWebhookEventProperties],
	// [SubscriptionInvoicingThresholdExceededWebhookEventProperties],
	// [SubscriptionLicenseAllocationResetWebhookEventProperties],
	// [SubscriptionLicenseBalanceThresholdReachedWebhookEventProperties],
	// [SubscriptionPlanChangeScheduledWebhookEventProperties],
	// [SubscriptionPlanChangedWebhookEventProperties],
	// [SubscriptionPlanVersionChangeScheduledWebhookEventProperties],
	// [SubscriptionPlanVersionChangedWebhookEventProperties],
	// [SubscriptionSpendExceededWebhookEventProperties],
	// [SubscriptionUsageExceededWebhookEventProperties],
	// [TransactionAccountingSyncFailedWebhookEventProperties],
	// [TransactionAccountingSyncSucceededWebhookEventProperties].
	Properties interface{} `json:"properties" api:"required"`
	// The event this payload describes.
	Type UnwrapWebhookEventType `json:"type" api:"required"`
	// This field can have the runtime type of
	// [CreditBlockAccountingSyncFailedWebhookEventAccountingSyncRecord],
	// [CreditBlockAccountingSyncSucceededWebhookEventAccountingSyncRecord],
	// [CreditNoteAccountingSyncFailedWebhookEventAccountingSyncRecord],
	// [CreditNoteAccountingSyncSucceededWebhookEventAccountingSyncRecord],
	// [CustomerAccountingSyncFailedWebhookEventAccountingSyncRecord],
	// [CustomerAccountingSyncSucceededWebhookEventAccountingSyncRecord],
	// [InvoiceAccountingSyncFailedWebhookEventAccountingSyncRecord],
	// [InvoiceAccountingSyncSucceededWebhookEventAccountingSyncRecord],
	// [SalesOrderAccountingSyncFailedWebhookEventAccountingSyncRecord],
	// [SalesOrderAccountingSyncSucceededWebhookEventAccountingSyncRecord],
	// [SubscriptionAccountingSyncFailedWebhookEventAccountingSyncRecord],
	// [SubscriptionAccountingSyncSucceededWebhookEventAccountingSyncRecord],
	// [TransactionAccountingSyncFailedWebhookEventAccountingSyncRecord],
	// [TransactionAccountingSyncSucceededWebhookEventAccountingSyncRecord].
	AccountingSyncRecord interface{} `json:"accounting_sync_record"`
	// This field can have the runtime type of [BackfillReflectedWebhookEventBackfill],
	// [BackfillRevertedWebhookEventBackfill].
	Backfill interface{} `json:"backfill"`
	// The Metric resource represents a calculation of a quantity based on events.
	// Metrics are defined by the query that transforms raw usage events into
	// meaningful values for your customers.
	BillableMetric BillableMetric `json:"billable_metric"`
	// This field can have the runtime type of
	// [CreditBlockAccountingSyncFailedWebhookEventBlock],
	// [CreditBlockAccountingSyncSucceededWebhookEventBlock].
	Block interface{} `json:"block"`
	// The [Credit Note](/invoicing/credit-notes) resource represents a credit that has
	// been applied to a particular invoice.
	CreditNote shared.CreditNote `json:"credit_note"`
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
	Customer Customer `json:"customer"`
	// This field can have the runtime type of [shared.Invoice],
	// [InvoiceAutomationScheduleStepExecutedWebhookEventInvoice], [string],
	// [InvoiceInvoiceDateElapsedWebhookEventInvoice],
	// [InvoiceIssuedSummaryWebhookEventInvoice].
	Invoice interface{} `json:"invoice"`
	// This field can have the runtime type of
	// [PaymentMethodCreatedWebhookEventPaymentMethod],
	// [PaymentMethodDeletedWebhookEventPaymentMethod].
	PaymentMethod interface{} `json:"payment_method"`
	// This field can have the runtime type of [PlanDefaultVersionSetWebhookEventPlan],
	// [PlanVersionCreatedWebhookEventPlan].
	Plan interface{} `json:"plan"`
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
	Price shared.Price `json:"price"`
	// This field can have the runtime type of [Subscription],
	// [SubscriptionAlertDisabledWebhookEventSubscription],
	// [SubscriptionGroupedCostExceededWebhookEventSubscription],
	// [SubscriptionInvoicingThresholdExceededWebhookEventSubscription],
	// [SubscriptionLicenseAllocationResetWebhookEventSubscription],
	// [SubscriptionLicenseBalanceThresholdReachedWebhookEventSubscription].
	Subscription interface{} `json:"subscription"`
	// This field can have the runtime type of
	// [TransactionAccountingSyncFailedWebhookEventTransaction],
	// [TransactionAccountingSyncSucceededWebhookEventTransaction].
	Transaction interface{}            `json:"transaction"`
	JSON        unwrapWebhookEventJSON `json:"-"`
	union       UnwrapWebhookEventUnion
}

// unwrapWebhookEventJSON contains the JSON metadata for the struct
// [UnwrapWebhookEvent]
type unwrapWebhookEventJSON struct {
	ID                   apijson.Field
	CreatedAt            apijson.Field
	Properties           apijson.Field
	Type                 apijson.Field
	AccountingSyncRecord apijson.Field
	Backfill             apijson.Field
	BillableMetric       apijson.Field
	Block                apijson.Field
	CreditNote           apijson.Field
	Customer             apijson.Field
	Invoice              apijson.Field
	PaymentMethod        apijson.Field
	Plan                 apijson.Field
	Price                apijson.Field
	Subscription         apijson.Field
	Transaction          apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r unwrapWebhookEventJSON) RawJSON() string {
	return r.raw
}

func (r *UnwrapWebhookEvent) UnmarshalJSON(data []byte) (err error) {
	*r = UnwrapWebhookEvent{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UnwrapWebhookEventUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [BackfillReflectedWebhookEvent],
// [BackfillRevertedWebhookEvent], [BillableMetricEditedWebhookEvent],
// [CreditBlockAccountingSyncFailedWebhookEvent],
// [CreditBlockAccountingSyncSucceededWebhookEvent],
// [CreditNoteAccountingSyncFailedWebhookEvent],
// [CreditNoteAccountingSyncSucceededWebhookEvent], [CreditNoteIssuedWebhookEvent],
// [CreditNoteMarkedAsVoidWebhookEvent],
// [CustomerAccountingSyncFailedWebhookEvent],
// [CustomerAccountingSyncSucceededWebhookEvent],
// [CustomerBalanceTransactionCreatedWebhookEvent], [CustomerCreatedWebhookEvent],
// [CustomerCreditBalanceDepletedWebhookEvent],
// [CustomerCreditBalanceDroppedWebhookEvent],
// [CustomerCreditBalanceRecoveredWebhookEvent],
// [CustomerCreditLedgerIncrementedWebhookEvent], [CustomerEditedWebhookEvent],
// [DataExportsTransferErrorWebhookEvent],
// [DataExportsTransferSuccessWebhookEvent], [EventUnmatchedEventWebhookEvent],
// [IngestionUnmatchedCustomerIDsWebhookEvent],
// [InvoiceAccountingSyncFailedWebhookEvent],
// [InvoiceAccountingSyncSucceededWebhookEvent],
// [InvoiceAutomationScheduleStepExecutedWebhookEvent],
// [InvoiceCostDataExportedWebhookEvent],
// [InvoiceDunningScheduleCreatedWebhookEvent],
// [InvoiceDunningScheduleEndedWebhookEvent],
// [InvoiceDunningScheduleResetWebhookEvent],
// [InvoiceDunningScheduleStepExecutedWebhookEvent], [InvoiceEditedWebhookEvent],
// [InvoiceInvoiceDateElapsedWebhookEvent], [InvoiceIssueFailedWebhookEvent],
// [InvoiceIssuedWebhookEvent], [InvoiceIssuedSummaryWebhookEvent],
// [InvoiceManuallyMarkedAsPaidWebhookEvent],
// [InvoiceManuallyMarkedAsVoidWebhookEvent], [InvoicePaymentFailedWebhookEvent],
// [InvoicePaymentProcessingWebhookEvent], [InvoicePaymentSucceededWebhookEvent],
// [InvoiceSyncFailedWebhookEvent], [InvoiceSyncSucceededWebhookEvent],
// [InvoiceUndoMarkAsPaidWebhookEvent],
// [InvoiceDueDateRecalculationCanceledWebhookEvent],
// [InvoiceDueDateRecalculationCompletedWebhookEvent],
// [InvoiceDueDateRecalculationStartedWebhookEvent],
// [MetricEventsDroppedByWatermarkWebhookEvent],
// [PaymentMethodCreatedWebhookEvent], [PaymentMethodDeletedWebhookEvent],
// [PlanDefaultVersionSetWebhookEvent], [PlanVersionCreatedWebhookEvent],
// [PriceEditedWebhookEvent], [ResourceEventTestWebhookEvent],
// [SalesOrderAccountingSyncFailedWebhookEvent],
// [SalesOrderAccountingSyncSucceededWebhookEvent],
// [SubscriptionAccountingSyncFailedWebhookEvent],
// [SubscriptionAccountingSyncSucceededWebhookEvent],
// [SubscriptionAlertDisabledWebhookEvent],
// [SubscriptionCancellationScheduledWebhookEvent],
// [SubscriptionCancellationUnscheduledWebhookEvent],
// [SubscriptionCostExceededWebhookEvent], [SubscriptionCreatedWebhookEvent],
// [SubscriptionEditedWebhookEvent], [SubscriptionEndedWebhookEvent],
// [SubscriptionFixedFeeQuantityUpdatedWebhookEvent],
// [SubscriptionGroupedCostExceededWebhookEvent],
// [SubscriptionInvoicingThresholdExceededWebhookEvent],
// [SubscriptionLicenseAllocationResetWebhookEvent],
// [SubscriptionLicenseBalanceThresholdReachedWebhookEvent],
// [SubscriptionPlanChangeScheduledWebhookEvent],
// [SubscriptionPlanChangedWebhookEvent],
// [SubscriptionPlanVersionChangeScheduledWebhookEvent],
// [SubscriptionPlanVersionChangedWebhookEvent],
// [SubscriptionSpendExceededWebhookEvent], [SubscriptionStartedWebhookEvent],
// [SubscriptionTrialEndedWebhookEvent], [SubscriptionUsageExceededWebhookEvent],
// [TransactionAccountingSyncFailedWebhookEvent],
// [TransactionAccountingSyncSucceededWebhookEvent].
func (r UnwrapWebhookEvent) AsUnion() UnwrapWebhookEventUnion {
	return r.union
}

// Issued when a backfill is closed and its events are reflected into usage.
//
// Union satisfied by [BackfillReflectedWebhookEvent],
// [BackfillRevertedWebhookEvent], [BillableMetricEditedWebhookEvent],
// [CreditBlockAccountingSyncFailedWebhookEvent],
// [CreditBlockAccountingSyncSucceededWebhookEvent],
// [CreditNoteAccountingSyncFailedWebhookEvent],
// [CreditNoteAccountingSyncSucceededWebhookEvent], [CreditNoteIssuedWebhookEvent],
// [CreditNoteMarkedAsVoidWebhookEvent],
// [CustomerAccountingSyncFailedWebhookEvent],
// [CustomerAccountingSyncSucceededWebhookEvent],
// [CustomerBalanceTransactionCreatedWebhookEvent], [CustomerCreatedWebhookEvent],
// [CustomerCreditBalanceDepletedWebhookEvent],
// [CustomerCreditBalanceDroppedWebhookEvent],
// [CustomerCreditBalanceRecoveredWebhookEvent],
// [CustomerCreditLedgerIncrementedWebhookEvent], [CustomerEditedWebhookEvent],
// [DataExportsTransferErrorWebhookEvent],
// [DataExportsTransferSuccessWebhookEvent], [EventUnmatchedEventWebhookEvent],
// [IngestionUnmatchedCustomerIDsWebhookEvent],
// [InvoiceAccountingSyncFailedWebhookEvent],
// [InvoiceAccountingSyncSucceededWebhookEvent],
// [InvoiceAutomationScheduleStepExecutedWebhookEvent],
// [InvoiceCostDataExportedWebhookEvent],
// [InvoiceDunningScheduleCreatedWebhookEvent],
// [InvoiceDunningScheduleEndedWebhookEvent],
// [InvoiceDunningScheduleResetWebhookEvent],
// [InvoiceDunningScheduleStepExecutedWebhookEvent], [InvoiceEditedWebhookEvent],
// [InvoiceInvoiceDateElapsedWebhookEvent], [InvoiceIssueFailedWebhookEvent],
// [InvoiceIssuedWebhookEvent], [InvoiceIssuedSummaryWebhookEvent],
// [InvoiceManuallyMarkedAsPaidWebhookEvent],
// [InvoiceManuallyMarkedAsVoidWebhookEvent], [InvoicePaymentFailedWebhookEvent],
// [InvoicePaymentProcessingWebhookEvent], [InvoicePaymentSucceededWebhookEvent],
// [InvoiceSyncFailedWebhookEvent], [InvoiceSyncSucceededWebhookEvent],
// [InvoiceUndoMarkAsPaidWebhookEvent],
// [InvoiceDueDateRecalculationCanceledWebhookEvent],
// [InvoiceDueDateRecalculationCompletedWebhookEvent],
// [InvoiceDueDateRecalculationStartedWebhookEvent],
// [MetricEventsDroppedByWatermarkWebhookEvent],
// [PaymentMethodCreatedWebhookEvent], [PaymentMethodDeletedWebhookEvent],
// [PlanDefaultVersionSetWebhookEvent], [PlanVersionCreatedWebhookEvent],
// [PriceEditedWebhookEvent], [ResourceEventTestWebhookEvent],
// [SalesOrderAccountingSyncFailedWebhookEvent],
// [SalesOrderAccountingSyncSucceededWebhookEvent],
// [SubscriptionAccountingSyncFailedWebhookEvent],
// [SubscriptionAccountingSyncSucceededWebhookEvent],
// [SubscriptionAlertDisabledWebhookEvent],
// [SubscriptionCancellationScheduledWebhookEvent],
// [SubscriptionCancellationUnscheduledWebhookEvent],
// [SubscriptionCostExceededWebhookEvent], [SubscriptionCreatedWebhookEvent],
// [SubscriptionEditedWebhookEvent], [SubscriptionEndedWebhookEvent],
// [SubscriptionFixedFeeQuantityUpdatedWebhookEvent],
// [SubscriptionGroupedCostExceededWebhookEvent],
// [SubscriptionInvoicingThresholdExceededWebhookEvent],
// [SubscriptionLicenseAllocationResetWebhookEvent],
// [SubscriptionLicenseBalanceThresholdReachedWebhookEvent],
// [SubscriptionPlanChangeScheduledWebhookEvent],
// [SubscriptionPlanChangedWebhookEvent],
// [SubscriptionPlanVersionChangeScheduledWebhookEvent],
// [SubscriptionPlanVersionChangedWebhookEvent],
// [SubscriptionSpendExceededWebhookEvent], [SubscriptionStartedWebhookEvent],
// [SubscriptionTrialEndedWebhookEvent], [SubscriptionUsageExceededWebhookEvent],
// [TransactionAccountingSyncFailedWebhookEvent] or
// [TransactionAccountingSyncSucceededWebhookEvent].
type UnwrapWebhookEventUnion interface {
	implementsUnwrapWebhookEvent()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UnwrapWebhookEventUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BackfillReflectedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BackfillRevertedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BillableMetricEditedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditBlockAccountingSyncFailedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditBlockAccountingSyncSucceededWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditNoteAccountingSyncFailedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditNoteAccountingSyncSucceededWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditNoteIssuedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditNoteMarkedAsVoidWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerAccountingSyncFailedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerAccountingSyncSucceededWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerBalanceTransactionCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerCreditBalanceDepletedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerCreditBalanceDroppedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerCreditBalanceRecoveredWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerCreditLedgerIncrementedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CustomerEditedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DataExportsTransferErrorWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DataExportsTransferSuccessWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EventUnmatchedEventWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IngestionUnmatchedCustomerIDsWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceAccountingSyncFailedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceAccountingSyncSucceededWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceAutomationScheduleStepExecutedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceCostDataExportedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceDunningScheduleCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceDunningScheduleEndedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceDunningScheduleResetWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceDunningScheduleStepExecutedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceEditedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceInvoiceDateElapsedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceIssueFailedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceIssuedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceIssuedSummaryWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceManuallyMarkedAsPaidWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceManuallyMarkedAsVoidWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoicePaymentFailedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoicePaymentProcessingWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoicePaymentSucceededWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceSyncFailedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceSyncSucceededWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceUndoMarkAsPaidWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceDueDateRecalculationCanceledWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceDueDateRecalculationCompletedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InvoiceDueDateRecalculationStartedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MetricEventsDroppedByWatermarkWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PaymentMethodCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PaymentMethodDeletedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PlanDefaultVersionSetWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PlanVersionCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PriceEditedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResourceEventTestWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SalesOrderAccountingSyncFailedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SalesOrderAccountingSyncSucceededWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionAccountingSyncFailedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionAccountingSyncSucceededWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionAlertDisabledWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionCancellationScheduledWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionCancellationUnscheduledWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionCostExceededWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionCreatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionEditedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionEndedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionFixedFeeQuantityUpdatedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionGroupedCostExceededWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionInvoicingThresholdExceededWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionLicenseAllocationResetWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionLicenseBalanceThresholdReachedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionPlanChangeScheduledWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionPlanChangedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionPlanVersionChangeScheduledWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionPlanVersionChangedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionSpendExceededWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionStartedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionTrialEndedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionUsageExceededWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionAccountingSyncFailedWebhookEvent{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TransactionAccountingSyncSucceededWebhookEvent{}),
		},
	)
}

// The event this payload describes.
type UnwrapWebhookEventType string

const (
	UnwrapWebhookEventTypeBackfillReflected                          UnwrapWebhookEventType = "backfill.reflected"
	UnwrapWebhookEventTypeBackfillReverted                           UnwrapWebhookEventType = "backfill.reverted"
	UnwrapWebhookEventTypeBillableMetricEdited                       UnwrapWebhookEventType = "billable_metric.edited"
	UnwrapWebhookEventTypeCreditBlockAccountingSyncFailed            UnwrapWebhookEventType = "credit_block.accounting_sync_failed"
	UnwrapWebhookEventTypeCreditBlockAccountingSyncSucceeded         UnwrapWebhookEventType = "credit_block.accounting_sync_succeeded"
	UnwrapWebhookEventTypeCreditNoteAccountingSyncFailed             UnwrapWebhookEventType = "credit_note.accounting_sync_failed"
	UnwrapWebhookEventTypeCreditNoteAccountingSyncSucceeded          UnwrapWebhookEventType = "credit_note.accounting_sync_succeeded"
	UnwrapWebhookEventTypeCreditNoteIssued                           UnwrapWebhookEventType = "credit_note.issued"
	UnwrapWebhookEventTypeCreditNoteMarkedAsVoid                     UnwrapWebhookEventType = "credit_note.marked_as_void"
	UnwrapWebhookEventTypeCustomerAccountingSyncFailed               UnwrapWebhookEventType = "customer.accounting_sync_failed"
	UnwrapWebhookEventTypeCustomerAccountingSyncSucceeded            UnwrapWebhookEventType = "customer.accounting_sync_succeeded"
	UnwrapWebhookEventTypeCustomerBalanceTransactionCreated          UnwrapWebhookEventType = "customer.balance_transaction_created"
	UnwrapWebhookEventTypeCustomerCreated                            UnwrapWebhookEventType = "customer.created"
	UnwrapWebhookEventTypeCustomerCreditBalanceDepleted              UnwrapWebhookEventType = "customer.credit_balance_depleted"
	UnwrapWebhookEventTypeCustomerCreditBalanceDropped               UnwrapWebhookEventType = "customer.credit_balance_dropped"
	UnwrapWebhookEventTypeCustomerCreditBalanceRecovered             UnwrapWebhookEventType = "customer.credit_balance_recovered"
	UnwrapWebhookEventTypeCustomerCreditLedgerIncremented            UnwrapWebhookEventType = "customer.credit_ledger_incremented"
	UnwrapWebhookEventTypeCustomerEdited                             UnwrapWebhookEventType = "customer.edited"
	UnwrapWebhookEventTypeDataExportsTransferError                   UnwrapWebhookEventType = "data_exports.transfer_error"
	UnwrapWebhookEventTypeDataExportsTransferSuccess                 UnwrapWebhookEventType = "data_exports.transfer_success"
	UnwrapWebhookEventTypeEventUnmatchedEvent                        UnwrapWebhookEventType = "event.unmatched_event"
	UnwrapWebhookEventTypeIngestionUnmatchedCustomerIDs              UnwrapWebhookEventType = "ingestion.unmatched_customer_ids"
	UnwrapWebhookEventTypeInvoiceAccountingSyncFailed                UnwrapWebhookEventType = "invoice.accounting_sync_failed"
	UnwrapWebhookEventTypeInvoiceAccountingSyncSucceeded             UnwrapWebhookEventType = "invoice.accounting_sync_succeeded"
	UnwrapWebhookEventTypeInvoiceAutomationScheduleStepExecuted      UnwrapWebhookEventType = "invoice.automation_schedule_step_executed"
	UnwrapWebhookEventTypeInvoiceCostDataExported                    UnwrapWebhookEventType = "invoice.cost_data_exported"
	UnwrapWebhookEventTypeInvoiceDunningScheduleCreated              UnwrapWebhookEventType = "invoice.dunning_schedule_created"
	UnwrapWebhookEventTypeInvoiceDunningScheduleEnded                UnwrapWebhookEventType = "invoice.dunning_schedule_ended"
	UnwrapWebhookEventTypeInvoiceDunningScheduleReset                UnwrapWebhookEventType = "invoice.dunning_schedule_reset"
	UnwrapWebhookEventTypeInvoiceDunningScheduleStepExecuted         UnwrapWebhookEventType = "invoice.dunning_schedule_step_executed"
	UnwrapWebhookEventTypeInvoiceEdited                              UnwrapWebhookEventType = "invoice.edited"
	UnwrapWebhookEventTypeInvoiceInvoiceDateElapsed                  UnwrapWebhookEventType = "invoice.invoice_date_elapsed"
	UnwrapWebhookEventTypeInvoiceIssueFailed                         UnwrapWebhookEventType = "invoice.issue_failed"
	UnwrapWebhookEventTypeInvoiceIssued                              UnwrapWebhookEventType = "invoice.issued"
	UnwrapWebhookEventTypeInvoiceIssuedSummary                       UnwrapWebhookEventType = "invoice.issued_summary"
	UnwrapWebhookEventTypeInvoiceManuallyMarkedAsPaid                UnwrapWebhookEventType = "invoice.manually_marked_as_paid"
	UnwrapWebhookEventTypeInvoiceManuallyMarkedAsVoid                UnwrapWebhookEventType = "invoice.manually_marked_as_void"
	UnwrapWebhookEventTypeInvoicePaymentFailed                       UnwrapWebhookEventType = "invoice.payment_failed"
	UnwrapWebhookEventTypeInvoicePaymentProcessing                   UnwrapWebhookEventType = "invoice.payment_processing"
	UnwrapWebhookEventTypeInvoicePaymentSucceeded                    UnwrapWebhookEventType = "invoice.payment_succeeded"
	UnwrapWebhookEventTypeInvoiceSyncFailed                          UnwrapWebhookEventType = "invoice.sync_failed"
	UnwrapWebhookEventTypeInvoiceSyncSucceeded                       UnwrapWebhookEventType = "invoice.sync_succeeded"
	UnwrapWebhookEventTypeInvoiceUndoMarkAsPaid                      UnwrapWebhookEventType = "invoice.undo_mark_as_paid"
	UnwrapWebhookEventTypeInvoiceDueDateRecalculationCanceled        UnwrapWebhookEventType = "invoice_due_date_recalculation.canceled"
	UnwrapWebhookEventTypeInvoiceDueDateRecalculationCompleted       UnwrapWebhookEventType = "invoice_due_date_recalculation.completed"
	UnwrapWebhookEventTypeInvoiceDueDateRecalculationStarted         UnwrapWebhookEventType = "invoice_due_date_recalculation.started"
	UnwrapWebhookEventTypeMetricEventsDroppedByWatermark             UnwrapWebhookEventType = "metric.events_dropped_by_watermark"
	UnwrapWebhookEventTypePaymentMethodCreated                       UnwrapWebhookEventType = "payment_method.created"
	UnwrapWebhookEventTypePaymentMethodDeleted                       UnwrapWebhookEventType = "payment_method.deleted"
	UnwrapWebhookEventTypePlanDefaultVersionSet                      UnwrapWebhookEventType = "plan.default_version_set"
	UnwrapWebhookEventTypePlanVersionCreated                         UnwrapWebhookEventType = "plan.version_created"
	UnwrapWebhookEventTypePriceEdited                                UnwrapWebhookEventType = "price.edited"
	UnwrapWebhookEventTypeResourceEventTest                          UnwrapWebhookEventType = "resource_event.test"
	UnwrapWebhookEventTypeSalesOrderAccountingSyncFailed             UnwrapWebhookEventType = "sales_order.accounting_sync_failed"
	UnwrapWebhookEventTypeSalesOrderAccountingSyncSucceeded          UnwrapWebhookEventType = "sales_order.accounting_sync_succeeded"
	UnwrapWebhookEventTypeSubscriptionAccountingSyncFailed           UnwrapWebhookEventType = "subscription.accounting_sync_failed"
	UnwrapWebhookEventTypeSubscriptionAccountingSyncSucceeded        UnwrapWebhookEventType = "subscription.accounting_sync_succeeded"
	UnwrapWebhookEventTypeSubscriptionAlertDisabled                  UnwrapWebhookEventType = "subscription.alert_disabled"
	UnwrapWebhookEventTypeSubscriptionCancellationScheduled          UnwrapWebhookEventType = "subscription.cancellation_scheduled"
	UnwrapWebhookEventTypeSubscriptionCancellationUnscheduled        UnwrapWebhookEventType = "subscription.cancellation_unscheduled"
	UnwrapWebhookEventTypeSubscriptionCostExceeded                   UnwrapWebhookEventType = "subscription.cost_exceeded"
	UnwrapWebhookEventTypeSubscriptionCreated                        UnwrapWebhookEventType = "subscription.created"
	UnwrapWebhookEventTypeSubscriptionEdited                         UnwrapWebhookEventType = "subscription.edited"
	UnwrapWebhookEventTypeSubscriptionEnded                          UnwrapWebhookEventType = "subscription.ended"
	UnwrapWebhookEventTypeSubscriptionFixedFeeQuantityUpdated        UnwrapWebhookEventType = "subscription.fixed_fee_quantity_updated"
	UnwrapWebhookEventTypeSubscriptionGroupedCostExceeded            UnwrapWebhookEventType = "subscription.grouped_cost_exceeded"
	UnwrapWebhookEventTypeSubscriptionInvoicingThresholdExceeded     UnwrapWebhookEventType = "subscription.invoicing_threshold_exceeded"
	UnwrapWebhookEventTypeSubscriptionLicenseAllocationReset         UnwrapWebhookEventType = "subscription.license_allocation_reset"
	UnwrapWebhookEventTypeSubscriptionLicenseBalanceThresholdReached UnwrapWebhookEventType = "subscription.license_balance_threshold_reached"
	UnwrapWebhookEventTypeSubscriptionPlanChangeScheduled            UnwrapWebhookEventType = "subscription.plan_change_scheduled"
	UnwrapWebhookEventTypeSubscriptionPlanChanged                    UnwrapWebhookEventType = "subscription.plan_changed"
	UnwrapWebhookEventTypeSubscriptionPlanVersionChangeScheduled     UnwrapWebhookEventType = "subscription.plan_version_change_scheduled"
	UnwrapWebhookEventTypeSubscriptionPlanVersionChanged             UnwrapWebhookEventType = "subscription.plan_version_changed"
	UnwrapWebhookEventTypeSubscriptionSpendExceeded                  UnwrapWebhookEventType = "subscription.spend_exceeded"
	UnwrapWebhookEventTypeSubscriptionStarted                        UnwrapWebhookEventType = "subscription.started"
	UnwrapWebhookEventTypeSubscriptionTrialEnded                     UnwrapWebhookEventType = "subscription.trial_ended"
	UnwrapWebhookEventTypeSubscriptionUsageExceeded                  UnwrapWebhookEventType = "subscription.usage_exceeded"
	UnwrapWebhookEventTypeTransactionAccountingSyncFailed            UnwrapWebhookEventType = "transaction.accounting_sync_failed"
	UnwrapWebhookEventTypeTransactionAccountingSyncSucceeded         UnwrapWebhookEventType = "transaction.accounting_sync_succeeded"
)

func (r UnwrapWebhookEventType) IsKnown() bool {
	switch r {
	case UnwrapWebhookEventTypeBackfillReflected, UnwrapWebhookEventTypeBackfillReverted, UnwrapWebhookEventTypeBillableMetricEdited, UnwrapWebhookEventTypeCreditBlockAccountingSyncFailed, UnwrapWebhookEventTypeCreditBlockAccountingSyncSucceeded, UnwrapWebhookEventTypeCreditNoteAccountingSyncFailed, UnwrapWebhookEventTypeCreditNoteAccountingSyncSucceeded, UnwrapWebhookEventTypeCreditNoteIssued, UnwrapWebhookEventTypeCreditNoteMarkedAsVoid, UnwrapWebhookEventTypeCustomerAccountingSyncFailed, UnwrapWebhookEventTypeCustomerAccountingSyncSucceeded, UnwrapWebhookEventTypeCustomerBalanceTransactionCreated, UnwrapWebhookEventTypeCustomerCreated, UnwrapWebhookEventTypeCustomerCreditBalanceDepleted, UnwrapWebhookEventTypeCustomerCreditBalanceDropped, UnwrapWebhookEventTypeCustomerCreditBalanceRecovered, UnwrapWebhookEventTypeCustomerCreditLedgerIncremented, UnwrapWebhookEventTypeCustomerEdited, UnwrapWebhookEventTypeDataExportsTransferError, UnwrapWebhookEventTypeDataExportsTransferSuccess, UnwrapWebhookEventTypeEventUnmatchedEvent, UnwrapWebhookEventTypeIngestionUnmatchedCustomerIDs, UnwrapWebhookEventTypeInvoiceAccountingSyncFailed, UnwrapWebhookEventTypeInvoiceAccountingSyncSucceeded, UnwrapWebhookEventTypeInvoiceAutomationScheduleStepExecuted, UnwrapWebhookEventTypeInvoiceCostDataExported, UnwrapWebhookEventTypeInvoiceDunningScheduleCreated, UnwrapWebhookEventTypeInvoiceDunningScheduleEnded, UnwrapWebhookEventTypeInvoiceDunningScheduleReset, UnwrapWebhookEventTypeInvoiceDunningScheduleStepExecuted, UnwrapWebhookEventTypeInvoiceEdited, UnwrapWebhookEventTypeInvoiceInvoiceDateElapsed, UnwrapWebhookEventTypeInvoiceIssueFailed, UnwrapWebhookEventTypeInvoiceIssued, UnwrapWebhookEventTypeInvoiceIssuedSummary, UnwrapWebhookEventTypeInvoiceManuallyMarkedAsPaid, UnwrapWebhookEventTypeInvoiceManuallyMarkedAsVoid, UnwrapWebhookEventTypeInvoicePaymentFailed, UnwrapWebhookEventTypeInvoicePaymentProcessing, UnwrapWebhookEventTypeInvoicePaymentSucceeded, UnwrapWebhookEventTypeInvoiceSyncFailed, UnwrapWebhookEventTypeInvoiceSyncSucceeded, UnwrapWebhookEventTypeInvoiceUndoMarkAsPaid, UnwrapWebhookEventTypeInvoiceDueDateRecalculationCanceled, UnwrapWebhookEventTypeInvoiceDueDateRecalculationCompleted, UnwrapWebhookEventTypeInvoiceDueDateRecalculationStarted, UnwrapWebhookEventTypeMetricEventsDroppedByWatermark, UnwrapWebhookEventTypePaymentMethodCreated, UnwrapWebhookEventTypePaymentMethodDeleted, UnwrapWebhookEventTypePlanDefaultVersionSet, UnwrapWebhookEventTypePlanVersionCreated, UnwrapWebhookEventTypePriceEdited, UnwrapWebhookEventTypeResourceEventTest, UnwrapWebhookEventTypeSalesOrderAccountingSyncFailed, UnwrapWebhookEventTypeSalesOrderAccountingSyncSucceeded, UnwrapWebhookEventTypeSubscriptionAccountingSyncFailed, UnwrapWebhookEventTypeSubscriptionAccountingSyncSucceeded, UnwrapWebhookEventTypeSubscriptionAlertDisabled, UnwrapWebhookEventTypeSubscriptionCancellationScheduled, UnwrapWebhookEventTypeSubscriptionCancellationUnscheduled, UnwrapWebhookEventTypeSubscriptionCostExceeded, UnwrapWebhookEventTypeSubscriptionCreated, UnwrapWebhookEventTypeSubscriptionEdited, UnwrapWebhookEventTypeSubscriptionEnded, UnwrapWebhookEventTypeSubscriptionFixedFeeQuantityUpdated, UnwrapWebhookEventTypeSubscriptionGroupedCostExceeded, UnwrapWebhookEventTypeSubscriptionInvoicingThresholdExceeded, UnwrapWebhookEventTypeSubscriptionLicenseAllocationReset, UnwrapWebhookEventTypeSubscriptionLicenseBalanceThresholdReached, UnwrapWebhookEventTypeSubscriptionPlanChangeScheduled, UnwrapWebhookEventTypeSubscriptionPlanChanged, UnwrapWebhookEventTypeSubscriptionPlanVersionChangeScheduled, UnwrapWebhookEventTypeSubscriptionPlanVersionChanged, UnwrapWebhookEventTypeSubscriptionSpendExceeded, UnwrapWebhookEventTypeSubscriptionStarted, UnwrapWebhookEventTypeSubscriptionTrialEnded, UnwrapWebhookEventTypeSubscriptionUsageExceeded, UnwrapWebhookEventTypeTransactionAccountingSyncFailed, UnwrapWebhookEventTypeTransactionAccountingSyncSucceeded:
		return true
	}
	return false
}

// Validates whether or not the webhook payload was sent by Orb. Pass an empty string to use the secret defined at the
// client level.
//
// An error will be raised if the webhook payload was not sent by Orb.
func (r *WebhookService) VerifySignature(payload []byte, headers http.Header, secret string, now time.Time) (err error) {
	return r.verifySignatureImpl(payload, headers, secret, now, 5*time.Minute)
}

// Identical to VerifySignature, but allows you to pass in a WebhookVerifySignatureParams struct to specify extra
// parameters such as the tolerance for X-Orb-Timestamp.
func (r *WebhookService) VerifySignatureWithParams(params WebhookVerifySignatureParams) (err error) {
	if params.Now.IsZero() {
		params.Now = time.Now()
	}
	if params.Tolerance == 0 {
		params.Tolerance = 5 * time.Minute
	}
	return r.verifySignatureImpl(params.Payload, params.Headers, params.Secret, params.Now, params.Tolerance)
}

func (r *WebhookService) verifySignatureImpl(payload []byte, headers http.Header, secret string, now time.Time, tolerance time.Duration) (err error) {
	if secret == "" {
		secret = r.webhookSecret
	}
	if secret == "" {
		return errors.New("The webhook secret must either be set using the env var, ORB_WEBHOOK_SECRET, on the client class, orb.NewClient(option.WithWebhookSecret(\"123\")}), or passed to this function")
	}

	msgSignature := headers.Values("X-Orb-Signature")
	if len(msgSignature) == 0 {
		return errors.New("could not find X-Orb-Signature header")
	}
	msgTimestamp := headers.Get("X-Orb-Timestamp")
	if len(msgTimestamp) == 0 {
		return errors.New("could not find X-Orb-Timestamp header")
	}

	timestamp, err := time.Parse(WebhookHeaderTimestampFormat, msgTimestamp)
	if err != nil {
		return fmt.Errorf("invalid timestamp headers: %s", err)
	}

	if timestamp.Before(now.Add(-tolerance)) {
		return fmt.Errorf("value from X-Orb-Timestamp header too old, tolerance=-%s", tolerance.String())
	}
	if timestamp.After(now.Add(tolerance)) {
		return fmt.Errorf("value from X-Orb-Timestamp header too new, tolerance=%s", tolerance.String())
	}

	secretBytes := []byte(secret)
	mac := hmac.New(sha256.New, secretBytes)
	mac.Write([]byte("v1:"))
	mac.Write([]byte(msgTimestamp))
	mac.Write([]byte(":"))
	mac.Write(payload)
	expected := mac.Sum(nil)

	for _, part := range msgSignature {
		parts := strings.Split(part, "=")
		if len(parts) != 2 {
			continue
		}
		if parts[0] != "v1" {
			continue
		}
		signature, err := hex.DecodeString(parts[1])
		if err != nil {
			continue
		}
		if hmac.Equal(signature, expected) {
			return nil
		}
	}

	return errors.New("None of the given webhook signatures match the expected signature")
}

type WebhookVerifySignatureParams struct {
	Payload   []byte
	Headers   http.Header
	Secret    string
	Now       time.Time
	Tolerance time.Duration
}
