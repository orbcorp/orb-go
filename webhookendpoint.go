// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
)

// A webhook endpoint is an HTTPS URL that Orb sends
// [webhook events](/integrations-and-exports/webhooks) to.
//
// WebhookEndpointService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookEndpointService] method instead.
type WebhookEndpointService struct {
	Options []option.RequestOption
}

// NewWebhookEndpointService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookEndpointService(opts ...option.RequestOption) (r *WebhookEndpointService) {
	r = &WebhookEndpointService{}
	r.Options = opts
	return
}

// This endpoint creates a webhook endpoint and returns the secret used to sign its
// webhook events. The secret isn't returned by any other endpoint. An account can
// have at most 20 webhook endpoints by default.
func (r *WebhookEndpointService) New(ctx context.Context, body WebhookEndpointNewParams, opts ...option.RequestOption) (res *WebhookEndpointWithSecret, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "webhook_endpoints"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint returns a webhook endpoint identified by its ID.
func (r *WebhookEndpointService) Get(ctx context.Context, webhookEndpointID string, opts ...option.RequestOption) (res *WebhookEndpoint, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookEndpointID == "" {
		err = errors.New("missing required webhook_endpoint_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhook_endpoints/%s", webhookEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This endpoint updates a webhook endpoint's URL or the event types sent to it.
func (r *WebhookEndpointService) Update(ctx context.Context, webhookEndpointID string, body WebhookEndpointUpdateParams, opts ...option.RequestOption) (res *WebhookEndpoint, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookEndpointID == "" {
		err = errors.New("missing required webhook_endpoint_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhook_endpoints/%s", webhookEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// This endpoint returns the account's webhook endpoints, ordered starting from the
// most recently created.
func (r *WebhookEndpointService) List(ctx context.Context, query WebhookEndpointListParams, opts ...option.RequestOption) (res *pagination.Page[WebhookEndpoint], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "webhook_endpoints"
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

// This endpoint returns the account's webhook endpoints, ordered starting from the
// most recently created.
func (r *WebhookEndpointService) ListAutoPaging(ctx context.Context, query WebhookEndpointListParams, opts ...option.RequestOption) *pagination.PageAutoPager[WebhookEndpoint] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint deletes a webhook endpoint. Orb stops sending webhook events to it
// immediately.
func (r *WebhookEndpointService) Delete(ctx context.Context, webhookEndpointID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if webhookEndpointID == "" {
		err = errors.New("missing required webhook_endpoint_id parameter")
		return err
	}
	path := fmt.Sprintf("webhook_endpoints/%s", webhookEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// This endpoint disables a webhook endpoint, so Orb stops sending webhook events
// to it until it is enabled again.
func (r *WebhookEndpointService) Disable(ctx context.Context, webhookEndpointID string, opts ...option.RequestOption) (res *WebhookEndpoint, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookEndpointID == "" {
		err = errors.New("missing required webhook_endpoint_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhook_endpoints/%s/disable", webhookEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// This endpoint enables a webhook endpoint, so Orb resumes sending webhook events
// to it.
func (r *WebhookEndpointService) Enable(ctx context.Context, webhookEndpointID string, opts ...option.RequestOption) (res *WebhookEndpoint, err error) {
	opts = slices.Concat(r.Options, opts)
	if webhookEndpointID == "" {
		err = errors.New("missing required webhook_endpoint_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("webhook_endpoints/%s/enable", webhookEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// A webhook endpoint is an HTTPS URL that Orb sends
// [webhook events](/integrations-and-exports/webhooks) to.
type WebhookEndpoint struct {
	// The Orb-assigned unique identifier for the webhook endpoint.
	ID string `json:"id" api:"required"`
	// The time the webhook endpoint was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether Orb sends webhook events to this endpoint.
	Enabled bool `json:"enabled" api:"required"`
	// The webhook event types sent to this endpoint. If null, all event types are
	// sent.
	EnabledEventTypes []string `json:"enabled_event_types" api:"required,nullable"`
	// The payload style of the webhook events sent to this endpoint. `standard`
	// payloads include each object's own fields and references to related objects,
	// without invoice line items or subscription price intervals. `expanded` payloads
	// include every object fully expanded.
	PayloadStyle WebhookEndpointPayloadStyle `json:"payload_style" api:"required"`
	// The HTTPS URL that Orb sends webhook events to.
	URL  string              `json:"url" api:"required"`
	JSON webhookEndpointJSON `json:"-"`
}

// webhookEndpointJSON contains the JSON metadata for the struct [WebhookEndpoint]
type webhookEndpointJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	Enabled           apijson.Field
	EnabledEventTypes apijson.Field
	PayloadStyle      apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *WebhookEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookEndpointJSON) RawJSON() string {
	return r.raw
}

// The payload style of the webhook events sent to this endpoint. `standard`
// payloads include each object's own fields and references to related objects,
// without invoice line items or subscription price intervals. `expanded` payloads
// include every object fully expanded.
type WebhookEndpointPayloadStyle string

const (
	WebhookEndpointPayloadStyleExpanded WebhookEndpointPayloadStyle = "expanded"
	WebhookEndpointPayloadStyleStandard WebhookEndpointPayloadStyle = "standard"
)

func (r WebhookEndpointPayloadStyle) IsKnown() bool {
	switch r {
	case WebhookEndpointPayloadStyleExpanded, WebhookEndpointPayloadStyleStandard:
		return true
	}
	return false
}

// A newly created webhook endpoint, along with the secret used to sign its webhook
// events.
type WebhookEndpointWithSecret struct {
	// The Orb-assigned unique identifier for the webhook endpoint.
	ID string `json:"id" api:"required"`
	// The time the webhook endpoint was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether Orb sends webhook events to this endpoint.
	Enabled bool `json:"enabled" api:"required"`
	// The webhook event types sent to this endpoint. If null, all event types are
	// sent.
	EnabledEventTypes []string `json:"enabled_event_types" api:"required,nullable"`
	// The payload style of the webhook events sent to this endpoint. `standard`
	// payloads include each object's own fields and references to related objects,
	// without invoice line items or subscription price intervals. `expanded` payloads
	// include every object fully expanded.
	PayloadStyle WebhookEndpointWithSecretPayloadStyle `json:"payload_style" api:"required"`
	// The secret used to sign the webhook events sent to this endpoint. Only returned
	// when the endpoint is created.
	Secret string `json:"secret" api:"required"`
	// The HTTPS URL that Orb sends webhook events to.
	URL  string                        `json:"url" api:"required"`
	JSON webhookEndpointWithSecretJSON `json:"-"`
}

// webhookEndpointWithSecretJSON contains the JSON metadata for the struct
// [WebhookEndpointWithSecret]
type webhookEndpointWithSecretJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	Enabled           apijson.Field
	EnabledEventTypes apijson.Field
	PayloadStyle      apijson.Field
	Secret            apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *WebhookEndpointWithSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookEndpointWithSecretJSON) RawJSON() string {
	return r.raw
}

// The payload style of the webhook events sent to this endpoint. `standard`
// payloads include each object's own fields and references to related objects,
// without invoice line items or subscription price intervals. `expanded` payloads
// include every object fully expanded.
type WebhookEndpointWithSecretPayloadStyle string

const (
	WebhookEndpointWithSecretPayloadStyleExpanded WebhookEndpointWithSecretPayloadStyle = "expanded"
	WebhookEndpointWithSecretPayloadStyleStandard WebhookEndpointWithSecretPayloadStyle = "standard"
)

func (r WebhookEndpointWithSecretPayloadStyle) IsKnown() bool {
	switch r {
	case WebhookEndpointWithSecretPayloadStyleExpanded, WebhookEndpointWithSecretPayloadStyleStandard:
		return true
	}
	return false
}

type WebhookEndpointNewParams struct {
	// The HTTPS URL that Orb sends webhook events to.
	URL param.Field[string] `json:"url" api:"required"`
	// The webhook event types to send to this endpoint. If omitted or null, all event
	// types are sent.
	EnabledEventTypes param.Field[[]string] `json:"enabled_event_types"`
	// The payload style of the webhook events sent to this endpoint. If omitted or
	// null, the account's default style is used.
	PayloadStyle param.Field[WebhookEndpointNewParamsPayloadStyle] `json:"payload_style"`
}

func (r WebhookEndpointNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The payload style of the webhook events sent to this endpoint. If omitted or
// null, the account's default style is used.
type WebhookEndpointNewParamsPayloadStyle string

const (
	WebhookEndpointNewParamsPayloadStyleExpanded WebhookEndpointNewParamsPayloadStyle = "expanded"
	WebhookEndpointNewParamsPayloadStyleStandard WebhookEndpointNewParamsPayloadStyle = "standard"
)

func (r WebhookEndpointNewParamsPayloadStyle) IsKnown() bool {
	switch r {
	case WebhookEndpointNewParamsPayloadStyleExpanded, WebhookEndpointNewParamsPayloadStyleStandard:
		return true
	}
	return false
}

type WebhookEndpointUpdateParams struct {
	// The webhook event types to send to this endpoint. Null sends all event types;
	// omit to leave them unchanged.
	EnabledEventTypes param.Field[[]string] `json:"enabled_event_types"`
	// The payload style of the webhook events sent to this endpoint. Applies to events
	// created after the change; omit to leave it unchanged.
	PayloadStyle param.Field[WebhookEndpointUpdateParamsPayloadStyle] `json:"payload_style"`
	// The HTTPS URL that Orb sends webhook events to.
	URL param.Field[string] `json:"url"`
}

func (r WebhookEndpointUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The payload style of the webhook events sent to this endpoint. Applies to events
// created after the change; omit to leave it unchanged.
type WebhookEndpointUpdateParamsPayloadStyle string

const (
	WebhookEndpointUpdateParamsPayloadStyleExpanded WebhookEndpointUpdateParamsPayloadStyle = "expanded"
	WebhookEndpointUpdateParamsPayloadStyleStandard WebhookEndpointUpdateParamsPayloadStyle = "standard"
)

func (r WebhookEndpointUpdateParamsPayloadStyle) IsKnown() bool {
	switch r {
	case WebhookEndpointUpdateParamsPayloadStyleExpanded, WebhookEndpointUpdateParamsPayloadStyleStandard:
		return true
	}
	return false
}

type WebhookEndpointListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [WebhookEndpointListParams]'s query parameters as
// `url.Values`.
func (r WebhookEndpointListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
