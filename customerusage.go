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
)

// CustomerUsageService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerUsageService] method instead.
type CustomerUsageService struct {
	Options []option.RequestOption
}

// NewCustomerUsageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerUsageService(opts ...option.RequestOption) (r *CustomerUsageService) {
	r = &CustomerUsageService{}
	r.Options = opts
	return
}

// This endpoint is used to amend usage within a timeframe for a customer that has
// an active subscription.
//
// This endpoint will mark _all_ existing events within
// `[timeframe_start, timeframe_end)` as _ignored_ for billing purposes, and Orb
// will only use the _new_ events passed in the body of this request as the source
// of truth for that timeframe moving forwards. Note that a given time period can
// be amended any number of times, so events can be overwritten in subsequent calls
// to th is endpoint.
//
// This is a powerful and audit-safe mechanism to retroactively change usage data
// in cases where you need to:
//
//   - decrease historical usage consumption because of degraded service availability
//     in your systems
//   - account for gaps from your usage reporting mechanism
//   - make point-in-time fixes for specific event records, while retaining the
//     original time of usage and associated metadata. This amendment API is designed
//     with two explicit goals:
//
//  1. Amendments are **always audit-safe**. The amendment process will still retain
//     original events in the timeframe, though they will be ignored for billing
//     calculations. For auditing a nd data fidelity purposes, Orb never overwrites
//     or permanently deletes ingested usage data.
//  2. Amendments always preserve data **consistency**. In other words, either an
//     amendment is fully processed by the system (and the new events for the
//     timeframe are honored rather than the existing ones) or the amendment request
//     fails. To maintain this important property, Orb prevents _partial event
//     ingestion_ on this endpoint.
//
// ## Response semantics
//
//   - Either all events are ingested successfully, or all fail to ingest (returning
//     a `4xx` or `5xx` response code).
//   - Any event that fails schema validation will lead to a `4xx` response. In this
//     case, to maintain data consistency, Orb will not ingest any events and will
//     also not deprecate existing events in the time period.
//   - You can assume that the amendment is successful on receipt of a `2xx`
//     response.While a successful response from this endpoint indicates that the new
//     events have been ingested, updating usage totals happens asynchronously and
//     may be delayed by a few minutes.
//
// As emphasized above, Orb will never show an inconsistent state (e.g. in invoice
// previews or dashboards); either it will show the existing state (before the
// amendment) or the new state (with new events in the requested timeframe).
//
// ## Sample request body
//
// ```json
//
//	{
//	  "events": [
//	    {
//	      "event_name": "payment_processed",
//	      "timestamp": "2022-03-24T07:15:00Z",
//	      "properties": {
//	        "amount": 100
//	      }
//	    },
//	    {
//	      "event_name": "payment_failed",
//	      "timestamp": "2022-03-24T07:15:00Z",
//	      "properties": {
//	        "amount": 100
//	      }
//	    }
//	  ]
//	}
//
// ```
//
// ## Request Validation
//
//   - The `timestamp` of each event reported must fall within the bounds of
//     `timeframe_start` and `timeframe_end`. As with ingestion, all timesta mps must
//     be sent in ISO8601 format with UTC timezone offset.
//   - Orb **does not accept an `idempotency_key`** with each event in this endpoint,
//     since the entirety of the event list must be ingested to ensure consistency.
//     On retryable errors , you should retry the request in its entirety, and assume
//     that the amendment operation has not succeeded until receipt of a `2xx`.
//
//   - Both `timeframe_start` and `timeframe_end` must be timestamps in the past.
//     Furthermore, Orb will genera lly validate that the `timeframe_start` and
//     `timeframe_end` fall within the customer's _current_ subscription billing pe
//     riod. However, Orb does allow amendments while in the grace period of the
//     previous billing period; in this instance, the timeframe can start before the
//     current period.
//
// ## API Limits
//
// Note that Orb does not currently enforce a hard rate- limit for API usage or a
// maximum request payload size. Similar to the event ingestion API, this API is
// architected for h igh-throughput ingestion. It is also safe to
// _programmatically_ call this endpoint if your system can automatically dete ct a
// need for historical amendment.
//
// In order to overwrite timeframes with a very large number of events, we suggest
// using multiple calls with small adjacent (e.g. every hour) timeframes.
//
// Deprecated: This method will be removed in a future release. Please use the
// 'events.backfills.create' instead.
func (r *CustomerUsageService) Update(ctx context.Context, id string, params CustomerUsageUpdateParams, opts ...option.RequestOption) (res *CustomerUsageUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/usage", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// This endpoint is used to amend usage within a timeframe for a customer that has
// an active subscription.
//
// This endpoint will mark _all_ existing events within
// `[timeframe_start, timeframe_end)` as _ignored_ for billing purposes, and Orb
// will only use the _new_ events passed in the body of this request as the source
// of truth for that timeframe moving forwards. Note that a given time period can
// be amended any number of times, so events can be overwritten in subsequent calls
// to th is endpoint.
//
// This is a powerful and audit-safe mechanism to retroactively change usage data
// in cases where you need to:
//
//   - decrease historical usage consumption because of degraded service availability
//     in your systems
//   - account for gaps from your usage reporting mechanism
//   - make point-in-time fixes for specific event records, while retaining the
//     original time of usage and associated metadata. This amendment API is designed
//     with two explicit goals:
//
//  1. Amendments are **always audit-safe**. The amendment process will still retain
//     original events in the timeframe, though they will be ignored for billing
//     calculations. For auditing a nd data fidelity purposes, Orb never overwrites
//     or permanently deletes ingested usage data.
//  2. Amendments always preserve data **consistency**. In other words, either an
//     amendment is fully processed by the system (and the new events for the
//     timeframe are honored rather than the existing ones) or the amendment request
//     fails. To maintain this important property, Orb prevents _partial event
//     ingestion_ on this endpoint.
//
// ## Response semantics
//
//   - Either all events are ingested successfully, or all fail to ingest (returning
//     a `4xx` or `5xx` response code).
//   - Any event that fails schema validation will lead to a `4xx` response. In this
//     case, to maintain data consistency, Orb will not ingest any events and will
//     also not deprecate existing events in the time period.
//   - You can assume that the amendment is successful on receipt of a `2xx`
//     response.While a successful response from this endpoint indicates that the new
//     events have been ingested, updating usage totals happens asynchronously and
//     may be delayed by a few minutes.
//
// As emphasized above, Orb will never show an inconsistent state (e.g. in invoice
// previews or dashboards); either it will show the existing state (before the
// amendment) or the new state (with new events in the requested timeframe).
//
// ## Sample request body
//
// ```json
//
//	{
//	  "events": [
//	    {
//	      "event_name": "payment_processed",
//	      "timestamp": "2022-03-24T07:15:00Z",
//	      "properties": {
//	        "amount": 100
//	      }
//	    },
//	    {
//	      "event_name": "payment_failed",
//	      "timestamp": "2022-03-24T07:15:00Z",
//	      "properties": {
//	        "amount": 100
//	      }
//	    }
//	  ]
//	}
//
// ```
//
// ## Request Validation
//
//   - The `timestamp` of each event reported must fall within the bounds of
//     `timeframe_start` and `timeframe_end`. As with ingestion, all timesta mps must
//     be sent in ISO8601 format with UTC timezone offset.
//   - Orb **does not accept an `idempotency_key`** with each event in this endpoint,
//     since the entirety of the event list must be ingested to ensure consistency.
//     On retryable errors , you should retry the request in its entirety, and assume
//     that the amendment operation has not succeeded until receipt of a `2xx`.
//
//   - Both `timeframe_start` and `timeframe_end` must be timestamps in the past.
//     Furthermore, Orb will genera lly validate that the `timeframe_start` and
//     `timeframe_end` fall within the customer's _current_ subscription billing pe
//     riod. However, Orb does allow amendments while in the grace period of the
//     previous billing period; in this instance, the timeframe can start before the
//     current period.
//
// ## API Limits
//
// Note that Orb does not currently enforce a hard rate- limit for API usage or a
// maximum request payload size. Similar to the event ingestion API, this API is
// architected for h igh-throughput ingestion. It is also safe to
// _programmatically_ call this endpoint if your system can automatically dete ct a
// need for historical amendment.
//
// In order to overwrite timeframes with a very large number of events, we suggest
// using multiple calls with small adjacent (e.g. every hour) timeframes.
//
// Deprecated: This method will be removed in a future release. Please use the
// 'events.backfills.create' instead.
func (r *CustomerUsageService) UpdateByExternalID(ctx context.Context, id string, params CustomerUsageUpdateByExternalIDParams, opts ...option.RequestOption) (res *CustomerUsageUpdateByExternalIDResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("customers/external_customer_id/%s/usage", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type CustomerUsageUpdateResponse struct {
	// An array of strings, corresponding to idempotency_key's marked as duplicates
	// (previously ingested)
	Duplicate []string `json:"duplicate,required"`
	// An array of strings, corresponding to idempotency_key's which were successfully
	// ingested.
	Ingested []string                        `json:"ingested,required"`
	JSON     customerUsageUpdateResponseJSON `json:"-"`
}

// customerUsageUpdateResponseJSON contains the JSON metadata for the struct
// [CustomerUsageUpdateResponse]
type customerUsageUpdateResponseJSON struct {
	Duplicate   apijson.Field
	Ingested    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerUsageUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerUsageUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerUsageUpdateByExternalIDResponse struct {
	// An array of strings, corresponding to idempotency_key's marked as duplicates
	// (previously ingested)
	Duplicate []string `json:"duplicate,required"`
	// An array of strings, corresponding to idempotency_key's which were successfully
	// ingested.
	Ingested []string                                    `json:"ingested,required"`
	JSON     customerUsageUpdateByExternalIDResponseJSON `json:"-"`
}

// customerUsageUpdateByExternalIDResponseJSON contains the JSON metadata for the
// struct [CustomerUsageUpdateByExternalIDResponse]
type customerUsageUpdateByExternalIDResponseJSON struct {
	Duplicate   apijson.Field
	Ingested    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerUsageUpdateByExternalIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerUsageUpdateByExternalIDResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerUsageUpdateParams struct {
	// Events to update
	Events param.Field[[]CustomerUsageUpdateParamsEvent] `json:"events,required"`
	// This bound is exclusive (i.e. events before this timestamp will be updated)
	TimeframeEnd param.Field[time.Time] `query:"timeframe_end" format:"date-time"`
	// This bound is inclusive (i.e. events with this timestamp onward, inclusive will
	// be updated)
	TimeframeStart param.Field[time.Time] `query:"timeframe_start" format:"date-time"`
}

func (r CustomerUsageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [CustomerUsageUpdateParams]'s query parameters as
// `url.Values`.
func (r CustomerUsageUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerUsageUpdateParamsEvent struct {
	// A name to meaningfully identify the action or event type.
	EventName param.Field[string] `json:"event_name,required"`
	// A dictionary of custom properties. Values in this dictionary must be numeric,
	// boolean, or strings. Nested dictionaries are disallowed.
	Properties param.Field[interface{}] `json:"properties,required"`
	// An ISO 8601 format date with no timezone offset (i.e. UTC). This should
	// represent the time that usage was recorded, and is particularly important to
	// attribute usage to a given billing period.
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	// The Orb Customer identifier
	CustomerID param.Field[string] `json:"customer_id"`
	// An alias for the Orb customer, whose mapping is specified when creating the
	// customer
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
}

func (r CustomerUsageUpdateParamsEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUsageUpdateByExternalIDParams struct {
	// Events to update
	Events param.Field[[]CustomerUsageUpdateByExternalIDParamsEvent] `json:"events,required"`
	// This bound is exclusive (i.e. events before this timestamp will be updated)
	TimeframeEnd param.Field[time.Time] `query:"timeframe_end" format:"date-time"`
	// This bound is inclusive (i.e. events with this timestamp onward, inclusive will
	// be updated)
	TimeframeStart param.Field[time.Time] `query:"timeframe_start" format:"date-time"`
}

func (r CustomerUsageUpdateByExternalIDParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [CustomerUsageUpdateByExternalIDParams]'s query parameters
// as `url.Values`.
func (r CustomerUsageUpdateByExternalIDParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerUsageUpdateByExternalIDParamsEvent struct {
	// A name to meaningfully identify the action or event type.
	EventName param.Field[string] `json:"event_name,required"`
	// A dictionary of custom properties. Values in this dictionary must be numeric,
	// boolean, or strings. Nested dictionaries are disallowed.
	Properties param.Field[interface{}] `json:"properties,required"`
	// An ISO 8601 format date with no timezone offset (i.e. UTC). This should
	// represent the time that usage was recorded, and is particularly important to
	// attribute usage to a given billing period.
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	// The Orb Customer identifier
	CustomerID param.Field[string] `json:"customer_id"`
	// An alias for the Orb customer, whose mapping is specified when creating the
	// customer
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
}

func (r CustomerUsageUpdateByExternalIDParamsEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
