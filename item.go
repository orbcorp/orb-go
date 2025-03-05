// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
	"github.com/orbcorp/orb-go/shared"
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
func (r *ItemService) New(ctx context.Context, body ItemNewParams, opts ...option.RequestOption) (res *shared.ItemModel, err error) {
	opts = append(r.Options[:], opts...)
	path := "items"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint can be used to update properties on the Item.
func (r *ItemService) Update(ctx context.Context, itemID string, body ItemUpdateParams, opts ...option.RequestOption) (res *shared.ItemModel, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *ItemService) List(ctx context.Context, query ItemListParams, opts ...option.RequestOption) (res *pagination.Page[shared.ItemModel], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *ItemService) ListAutoPaging(ctx context.Context, query ItemListParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.ItemModel] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint returns an item identified by its item_id.
func (r *ItemService) Fetch(ctx context.Context, itemID string, opts ...option.RequestOption) (res *shared.ItemModel, err error) {
	opts = append(r.Options[:], opts...)
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return
	}
	path := fmt.Sprintf("items/%s", itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ItemNewParams struct {
	// The name of the item.
	Name param.Field[string] `json:"name,required"`
}

func (r ItemNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ItemUpdateParams struct {
	ExternalConnections param.Field[[]shared.ItemExternalConnectionModelParam] `json:"external_connections"`
	Name                param.Field[string]                                    `json:"name"`
}

func (r ItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
