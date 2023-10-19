// File generated from our OpenAPI spec by Stainless.

package orb

import (
	"os"

	"github.com/orbcorp/orb-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the orb API. You should not instantiate this client directly,
// and instead use the [NewClient] method instead.
type Client struct {
	Options          []option.RequestOption
	TopLevel         *TopLevelService
	Coupons          *CouponService
	CreditNotes      *CreditNoteService
	Customers        *CustomerService
	Events           *EventService
	InvoiceLineItems *InvoiceLineItemService
	Invoices         *InvoiceService
	Items            *ItemService
	Metrics          *MetricService
	Plans            *PlanService
	Prices           *PriceService
	Subscriptions    *SubscriptionService
}

// NewClient generates a new client with the default option read from the
// environment (ORB_API_KEY). The option passed in as arguments are applied after
// these default arguments, and all option will be passed down to the services and
// requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("ORB_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.TopLevel = NewTopLevelService(opts...)
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

	return
}
