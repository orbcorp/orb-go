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

// ItemService contains methods and other services that help with interacting with
// the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewItemService] method instead.
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

// This endpoint is used to create an [Item](/core-concepts#item).
func (r *ItemService) New(ctx context.Context, body ItemNewParams, opts ...option.RequestOption) (res *Item, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "items"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint can be used to update properties on the Item.
func (r *ItemService) Update(ctx context.Context, itemID string, body ItemUpdateParams, opts ...option.RequestOption) (res *Item, err error) {
	opts = slices.Concat(r.Options, opts)
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return
	}
	path := fmt.Sprintf("items/%s", itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// This endpoint returns a list of all Items, ordered in descending order by
// creation time.
func (r *ItemService) List(ctx context.Context, query ItemListParams, opts ...option.RequestOption) (res *pagination.Page[Item], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
func (r *ItemService) ListAutoPaging(ctx context.Context, query ItemListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Item] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Archive item
func (r *ItemService) Archive(ctx context.Context, itemID string, opts ...option.RequestOption) (res *Item, err error) {
	opts = slices.Concat(r.Options, opts)
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return
	}
	path := fmt.Sprintf("items/%s/archive", itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This endpoint returns an item identified by its item_id.
func (r *ItemService) Fetch(ctx context.Context, itemID string, opts ...option.RequestOption) (res *Item, err error) {
	opts = slices.Concat(r.Options, opts)
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return
	}
	path := fmt.Sprintf("items/%s", itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// The Item resource represents a sellable product or good. Items are associated
// with all line items, billable metrics, and prices and are used for defining
// external sync behavior for invoices and tax calculation purposes.
type Item struct {
	// The Orb-assigned unique identifier for the item.
	ID string `json:"id,required"`
	// The time at which the item was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A list of external connections for this item, used to sync with external
	// invoicing and tax systems.
	ExternalConnections []ItemExternalConnection `json:"external_connections,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// The name of the item.
	Name string `json:"name,required"`
	// The time at which the item was archived. If null, the item is not archived.
	ArchivedAt time.Time `json:"archived_at,nullable" format:"date-time"`
	JSON       itemJSON  `json:"-"`
}

// itemJSON contains the JSON metadata for the struct [Item]
type itemJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	ExternalConnections apijson.Field
	Metadata            apijson.Field
	Name                apijson.Field
	ArchivedAt          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Item) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemJSON) RawJSON() string {
	return r.raw
}

// Represents a connection between an Item and an external system for invoicing or
// tax calculation purposes.
type ItemExternalConnection struct {
	// The name of the external system this item is connected to.
	ExternalConnectionName ItemExternalConnectionsExternalConnectionName `json:"external_connection_name,required"`
	// The identifier of this item in the external system.
	ExternalEntityID string                     `json:"external_entity_id,required"`
	JSON             itemExternalConnectionJSON `json:"-"`
}

// itemExternalConnectionJSON contains the JSON metadata for the struct
// [ItemExternalConnection]
type itemExternalConnectionJSON struct {
	ExternalConnectionName apijson.Field
	ExternalEntityID       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ItemExternalConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemExternalConnectionJSON) RawJSON() string {
	return r.raw
}

// The name of the external system this item is connected to.
type ItemExternalConnectionsExternalConnectionName string

const (
	ItemExternalConnectionsExternalConnectionNameStripe     ItemExternalConnectionsExternalConnectionName = "stripe"
	ItemExternalConnectionsExternalConnectionNameQuickbooks ItemExternalConnectionsExternalConnectionName = "quickbooks"
	ItemExternalConnectionsExternalConnectionNameBillCom    ItemExternalConnectionsExternalConnectionName = "bill.com"
	ItemExternalConnectionsExternalConnectionNameNetsuite   ItemExternalConnectionsExternalConnectionName = "netsuite"
	ItemExternalConnectionsExternalConnectionNameTaxjar     ItemExternalConnectionsExternalConnectionName = "taxjar"
	ItemExternalConnectionsExternalConnectionNameAvalara    ItemExternalConnectionsExternalConnectionName = "avalara"
	ItemExternalConnectionsExternalConnectionNameAnrok      ItemExternalConnectionsExternalConnectionName = "anrok"
	ItemExternalConnectionsExternalConnectionNameNumeral    ItemExternalConnectionsExternalConnectionName = "numeral"
)

func (r ItemExternalConnectionsExternalConnectionName) IsKnown() bool {
	switch r {
	case ItemExternalConnectionsExternalConnectionNameStripe, ItemExternalConnectionsExternalConnectionNameQuickbooks, ItemExternalConnectionsExternalConnectionNameBillCom, ItemExternalConnectionsExternalConnectionNameNetsuite, ItemExternalConnectionsExternalConnectionNameTaxjar, ItemExternalConnectionsExternalConnectionNameAvalara, ItemExternalConnectionsExternalConnectionNameAnrok, ItemExternalConnectionsExternalConnectionNameNumeral:
		return true
	}
	return false
}

type ItemNewParams struct {
	// The name of the item.
	Name param.Field[string] `json:"name,required"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r ItemNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ItemUpdateParams struct {
	ExternalConnections param.Field[[]ItemUpdateParamsExternalConnection] `json:"external_connections"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	Name     param.Field[string]            `json:"name"`
}

func (r ItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Represents a connection between an Item and an external system for invoicing or
// tax calculation purposes.
type ItemUpdateParamsExternalConnection struct {
	// The name of the external system this item is connected to.
	ExternalConnectionName param.Field[ItemUpdateParamsExternalConnectionsExternalConnectionName] `json:"external_connection_name,required"`
	// The identifier of this item in the external system.
	ExternalEntityID param.Field[string] `json:"external_entity_id,required"`
}

func (r ItemUpdateParamsExternalConnection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name of the external system this item is connected to.
type ItemUpdateParamsExternalConnectionsExternalConnectionName string

const (
	ItemUpdateParamsExternalConnectionsExternalConnectionNameStripe     ItemUpdateParamsExternalConnectionsExternalConnectionName = "stripe"
	ItemUpdateParamsExternalConnectionsExternalConnectionNameQuickbooks ItemUpdateParamsExternalConnectionsExternalConnectionName = "quickbooks"
	ItemUpdateParamsExternalConnectionsExternalConnectionNameBillCom    ItemUpdateParamsExternalConnectionsExternalConnectionName = "bill.com"
	ItemUpdateParamsExternalConnectionsExternalConnectionNameNetsuite   ItemUpdateParamsExternalConnectionsExternalConnectionName = "netsuite"
	ItemUpdateParamsExternalConnectionsExternalConnectionNameTaxjar     ItemUpdateParamsExternalConnectionsExternalConnectionName = "taxjar"
	ItemUpdateParamsExternalConnectionsExternalConnectionNameAvalara    ItemUpdateParamsExternalConnectionsExternalConnectionName = "avalara"
	ItemUpdateParamsExternalConnectionsExternalConnectionNameAnrok      ItemUpdateParamsExternalConnectionsExternalConnectionName = "anrok"
	ItemUpdateParamsExternalConnectionsExternalConnectionNameNumeral    ItemUpdateParamsExternalConnectionsExternalConnectionName = "numeral"
)

func (r ItemUpdateParamsExternalConnectionsExternalConnectionName) IsKnown() bool {
	switch r {
	case ItemUpdateParamsExternalConnectionsExternalConnectionNameStripe, ItemUpdateParamsExternalConnectionsExternalConnectionNameQuickbooks, ItemUpdateParamsExternalConnectionsExternalConnectionNameBillCom, ItemUpdateParamsExternalConnectionsExternalConnectionNameNetsuite, ItemUpdateParamsExternalConnectionsExternalConnectionNameTaxjar, ItemUpdateParamsExternalConnectionsExternalConnectionNameAvalara, ItemUpdateParamsExternalConnectionsExternalConnectionNameAnrok, ItemUpdateParamsExternalConnectionsExternalConnectionNameNumeral:
		return true
	}
	return false
}

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
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
