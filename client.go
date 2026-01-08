// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"net/http"
	"os"
	"slices"

	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the orb API. You should not instantiate this client directly,
// and instead use the [NewClient] method instead.
type Client struct {
	Options                []option.RequestOption
	TopLevel               *TopLevelService
	Beta                   *BetaService
	Coupons                *CouponService
	CreditNotes            *CreditNoteService
	Customers              *CustomerService
	Events                 *EventService
	InvoiceLineItems       *InvoiceLineItemService
	Invoices               *InvoiceService
	Items                  *ItemService
	Metrics                *MetricService
	Plans                  *PlanService
	Prices                 *PriceService
	Subscriptions          *SubscriptionService
	Alerts                 *AlertService
	DimensionalPriceGroups *DimensionalPriceGroupService
	SubscriptionChanges    *SubscriptionChangeService
	CreditBlocks           *CreditBlockService
}

// DefaultClientOptions read from the environment (ORB_API_KEY, ORB_WEBHOOK_SECRET,
// ORB_BASE_URL). This should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("ORB_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("ORB_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("ORB_WEBHOOK_SECRET"); ok {
		defaults = append(defaults, option.WithWebhookSecret(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (ORB_API_KEY, ORB_WEBHOOK_SECRET, ORB_BASE_URL). The option passed
// in as arguments are applied after these default arguments, and all option will
// be passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = &Client{Options: opts}

	r.TopLevel = NewTopLevelService(opts...)
	r.Beta = NewBetaService(opts...)
	r.Coupons = NewCouponService(opts...)
	r.CreditNotes = NewCreditNoteService(opts...)
	r.Customers = NewCustomerService(opts...)
	r.Events = NewEventService(opts...)
	r.InvoiceLineItems = NewInvoiceLineItemService(opts...)
	r.Invoices = NewInvoiceService(opts...)
	r.Items = NewItemService(opts...)
	r.Metrics = NewMetricService(opts...)
	r.Plans = NewPlanService(opts...)
	r.Prices = NewPriceService(opts...)
	r.Subscriptions = NewSubscriptionService(opts...)
	r.Alerts = NewAlertService(opts...)
	r.DimensionalPriceGroups = NewDimensionalPriceGroupService(opts...)
	r.SubscriptionChanges = NewSubscriptionChangeService(opts...)
	r.CreditBlocks = NewCreditBlockService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	opts = slices.Concat(r.Options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
