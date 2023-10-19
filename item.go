// File generated from our OpenAPI spec by Stainless.

package orb

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/internal/shared"
	"github.com/orbcorp/orb-go/option"
)

// ItemService contains methods and other services that help with interacting with
// the orb API. Note, unlike clients, this service does not read variables from the
// environment automatically. You should not instantiate this service directly, and
// instead use the [NewItemService] method instead.
type ItemService struct {
	Options []option.RequestOption
}

// NewItemService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewItemService(opts ...option.RequestOption) (r *ItemService) {
	r = &ItemService{}
	r.Options = opts
	return
}

// This endpoint returns a list of all Items, ordered in descending order by
// creation time.
func (r *ItemService) List(ctx context.Context, query ItemListParams, opts ...option.RequestOption) (res *shared.Page[ItemListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "items"
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

// This endpoint returns a list of all Items, ordered in descending order by
// creation time.
func (r *ItemService) ListAutoPaging(ctx context.Context, query ItemListParams, opts ...option.RequestOption) *shared.PageAutoPager[ItemListResponse] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint returns an item identified by its item_id.
func (r *ItemService) Fetch(ctx context.Context, itemID string, opts ...option.RequestOption) (res *ItemFetchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("items/%s", itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// The Item resource represents a sellable product or good. Items are associated
// with all line items, billable metrics, and prices and are used for defining
// external sync behavior for invoices and tax calculation purposes.
type ItemListResponse struct {
	ID                  string                               `json:"id,required"`
	CreatedAt           time.Time                            `json:"created_at,required" format:"date-time"`
	ExternalConnections []ItemListResponseExternalConnection `json:"external_connections,required"`
	Name                string                               `json:"name,required"`
	JSON                itemListResponseJSON
}

// itemListResponseJSON contains the JSON metadata for the struct
// [ItemListResponse]
type itemListResponseJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	ExternalConnections apijson.Field
	Name                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ItemListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ItemListResponseExternalConnection struct {
	ExternalConnectionName ItemListResponseExternalConnectionsExternalConnectionName `json:"external_connection_name,required"`
	ExternalEntityID       string                                                    `json:"external_entity_id,required"`
	JSON                   itemListResponseExternalConnectionJSON
}

// itemListResponseExternalConnectionJSON contains the JSON metadata for the struct
// [ItemListResponseExternalConnection]
type itemListResponseExternalConnectionJSON struct {
	ExternalConnectionName apijson.Field
	ExternalEntityID       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ItemListResponseExternalConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ItemListResponseExternalConnectionsExternalConnectionName string

const (
	ItemListResponseExternalConnectionsExternalConnectionNameStripe     ItemListResponseExternalConnectionsExternalConnectionName = "stripe"
	ItemListResponseExternalConnectionsExternalConnectionNameQuickbooks ItemListResponseExternalConnectionsExternalConnectionName = "quickbooks"
	ItemListResponseExternalConnectionsExternalConnectionNameBillCom    ItemListResponseExternalConnectionsExternalConnectionName = "bill.com"
	ItemListResponseExternalConnectionsExternalConnectionNameNetsuite   ItemListResponseExternalConnectionsExternalConnectionName = "netsuite"
	ItemListResponseExternalConnectionsExternalConnectionNameTaxjar     ItemListResponseExternalConnectionsExternalConnectionName = "taxjar"
	ItemListResponseExternalConnectionsExternalConnectionNameAvalara    ItemListResponseExternalConnectionsExternalConnectionName = "avalara"
	ItemListResponseExternalConnectionsExternalConnectionNameAnrok      ItemListResponseExternalConnectionsExternalConnectionName = "anrok"
)

// The Item resource represents a sellable product or good. Items are associated
// with all line items, billable metrics, and prices and are used for defining
// external sync behavior for invoices and tax calculation purposes.
type ItemFetchResponse struct {
	ID                  string                                `json:"id,required"`
	CreatedAt           time.Time                             `json:"created_at,required" format:"date-time"`
	ExternalConnections []ItemFetchResponseExternalConnection `json:"external_connections,required"`
	Name                string                                `json:"name,required"`
	JSON                itemFetchResponseJSON
}

// itemFetchResponseJSON contains the JSON metadata for the struct
// [ItemFetchResponse]
type itemFetchResponseJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	ExternalConnections apijson.Field
	Name                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ItemFetchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ItemFetchResponseExternalConnection struct {
	ExternalConnectionName ItemFetchResponseExternalConnectionsExternalConnectionName `json:"external_connection_name,required"`
	ExternalEntityID       string                                                     `json:"external_entity_id,required"`
	JSON                   itemFetchResponseExternalConnectionJSON
}

// itemFetchResponseExternalConnectionJSON contains the JSON metadata for the
// struct [ItemFetchResponseExternalConnection]
type itemFetchResponseExternalConnectionJSON struct {
	ExternalConnectionName apijson.Field
	ExternalEntityID       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ItemFetchResponseExternalConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ItemFetchResponseExternalConnectionsExternalConnectionName string

const (
	ItemFetchResponseExternalConnectionsExternalConnectionNameStripe     ItemFetchResponseExternalConnectionsExternalConnectionName = "stripe"
	ItemFetchResponseExternalConnectionsExternalConnectionNameQuickbooks ItemFetchResponseExternalConnectionsExternalConnectionName = "quickbooks"
	ItemFetchResponseExternalConnectionsExternalConnectionNameBillCom    ItemFetchResponseExternalConnectionsExternalConnectionName = "bill.com"
	ItemFetchResponseExternalConnectionsExternalConnectionNameNetsuite   ItemFetchResponseExternalConnectionsExternalConnectionName = "netsuite"
	ItemFetchResponseExternalConnectionsExternalConnectionNameTaxjar     ItemFetchResponseExternalConnectionsExternalConnectionName = "taxjar"
	ItemFetchResponseExternalConnectionsExternalConnectionNameAvalara    ItemFetchResponseExternalConnectionsExternalConnectionName = "avalara"
	ItemFetchResponseExternalConnectionsExternalConnectionNameAnrok      ItemFetchResponseExternalConnectionsExternalConnectionName = "anrok"
)

type ItemListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [ItemListParams]'s query parameters as `url.Values`.
func (r ItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
