// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the orb API. You should not instantiate this client directly,
// and instead use the [NewClient] method instead.
type Client struct {
	Options  []option.RequestOption
	TopLevel *TopLevelService
	// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
	// subscribed to by a customer. Plans define the billing behavior of the
	// subscription. You can see more about how to configure prices in the
	// [Price resource](/reference/price).
	Beta *BetaService
	// A coupon represents a reusable discount configuration that can be applied either
	// as a fixed or percentage amount to an invoice or subscription. Coupons are
	// activated using a redemption code, which applies the discount to a subscription
	// or invoice. The duration of a coupon determines how long it remains available
	// for use by end users.
	Coupons *CouponService
	// The [Credit Note](/invoicing/credit-notes) resource represents a credit that has
	// been applied to a particular invoice.
	CreditNotes *CreditNoteService
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
	Customers *CustomerService
	// The [Event](/core-concepts#event) resource represents a usage event that has
	// been created for a customer. Events are the core of Orb's usage-based billing
	// model, and are used to calculate the usage charges for a given billing period.
	Events *EventService
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	InvoiceLineItems *InvoiceLineItemService
	// An [`Invoice`](/core-concepts#invoice) is a fundamental billing entity,
	// representing the request for payment for a single subscription. This includes a
	// set of line items, which correspond to prices in the subscription's plan and can
	// represent fixed recurring fees or usage-based fees. They are generated at the
	// end of a billing period, or as the result of an action, such as a cancellation.
	Invoices *InvoiceService
	// The Item resource represents a sellable product or good. Items are associated
	// with all line items, billable metrics, and prices and are used for defining
	// external sync behavior for invoices and tax calculation purposes.
	Items *ItemService
	// The Metric resource represents a calculation of a quantity based on events.
	// Metrics are defined by the query that transforms raw usage events into
	// meaningful values for your customers.
	Metrics *MetricService
	// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
	// subscribed to by a customer. Plans define the billing behavior of the
	// subscription. You can see more about how to configure prices in the
	// [Price resource](/reference/price).
	Plans *PlanService
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
	Prices        *PriceService
	Subscriptions *SubscriptionService
	// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
	// usage, or credit balance and trigger webhooks when a threshold is exceeded.
	//
	// Alerts created through the API can be scoped to either customers or
	// subscriptions.
	Alerts                 *AlertService
	DimensionalPriceGroups *DimensionalPriceGroupService
	SubscriptionChanges    *SubscriptionChangeService
	// The [Credit Ledger Entry resource](/product-catalog/prepurchase) models prepaid
	// credits within Orb.
	CreditBlocks *CreditBlockService
	// The LicenseType resource represents a type of license that can be assigned to
	// users. License types are used during billing by grouping metrics on the
	// configured grouping key.
	LicenseTypes *LicenseTypeService
	Licenses     *LicenseService
}

// DefaultClientOptions read from the environment (ORB_API_KEY, ORB_WEBHOOK_SECRET,
// ORB_BASE_URL). This should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithHTTPClient(defaultHTTPClient()), option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("ORB_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("ORB_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("ORB_WEBHOOK_SECRET"); ok {
		defaults = append(defaults, option.WithWebhookSecret(o))
	}
	if o, ok := os.LookupEnv("ORB_CUSTOM_HEADERS"); ok {
		for _, line := range strings.Split(o, "\n") {
			colon := strings.Index(line, ":")
			if colon >= 0 {
				defaults = append(defaults, option.WithHeader(strings.TrimSpace(line[:colon]), strings.TrimSpace(line[colon+1:])))
			}
		}
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
	r.LicenseTypes = NewLicenseTypeService(opts...)
	r.Licenses = NewLicenseService(opts...)

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
