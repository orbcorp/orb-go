// File generated from our OpenAPI spec by Stainless.

package apierror

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/orbcorp/orb-go/internal/apijson"
)

// Error represents an error that originates from the API, i.e. when a request is
// made and the API returns a response with a HTTP status code. Other errors are
// not wrapped by this SDK.
type Error struct {
	Status           ErrorStatus   `json:"status,required"`
	Title            string        `json:"title,required,nullable"`
	Type             ErrorType     `json:"type,required"`
	ValidationErrors []interface{} `json:"validation_errors,required"`
	Detail           string        `json:"detail,nullable"`
	JSON             errorJSON     `json:"-"`
	StatusCode       int
	Request          *http.Request
	Response         *http.Response
}

// errorJSON contains the JSON metadata for the struct [Error]
type errorJSON struct {
	Status           apijson.Field
	Title            apijson.Field
	Type             apijson.Field
	ValidationErrors apijson.Field
	Detail           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r *Error) Error() string {
	body, _ := io.ReadAll(r.Response.Body)
	return fmt.Sprintf("%s \"%s\": %d %s %s", r.Request.Method, r.Request.URL, r.Response.StatusCode, http.StatusText(r.Response.StatusCode), string(body))
}

func (r *Error) DumpRequest(body bool) []byte {
	if r.Request.GetBody != nil {
		r.Request.Body, _ = r.Request.GetBody()
	}
	out, _ := httputil.DumpRequestOut(r.Request, body)
	return out
}

func (r *Error) DumpResponse(body bool) []byte {
	out, _ := httputil.DumpResponse(r.Response, body)
	return out
}

type ErrorStatus int64

const (
	ErrorStatus500 ErrorStatus = 500
)

type ErrorType string

const (
	ErrorTypeOrbInternalServerError ErrorType = "https://docs.withorb.com/reference/error-responses#500-internal-server-error"
)
