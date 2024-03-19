// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

func (r errorJSON) RawJSON() string {
	return r.raw
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
	ErrorStatus429 ErrorStatus = 429
	ErrorStatus413 ErrorStatus = 413
	ErrorStatus409 ErrorStatus = 409
	ErrorStatus404 ErrorStatus = 404
	ErrorStatus400 ErrorStatus = 400
	ErrorStatus401 ErrorStatus = 401
)

func (r ErrorStatus) IsKnown() bool {
	switch r {
	case ErrorStatus500, ErrorStatus429, ErrorStatus413, ErrorStatus409, ErrorStatus404, ErrorStatus400, ErrorStatus401:
		return true
	}
	return false
}

type ErrorType string

const (
	ErrorTypeOrbInternalServerError                                                 ErrorType = "https://docs.withorb.com/reference/error-responses#500-internal-server-error"
	ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses429TooManyRequests           ErrorType = "https://docs.withorb.com/reference/error-responses#429-too-many-requests"
	ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses413ResourceTooLarge          ErrorType = "https://docs.withorb.com/reference/error-responses#413-resource-too-large"
	ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses413RequestTooLarge           ErrorType = "https://docs.withorb.com/reference/error-responses#413-request-too-large"
	ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses409ResourceConflict          ErrorType = "https://docs.withorb.com/reference/error-responses#409-resource-conflict"
	ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses404URLNotFound               ErrorType = "https://docs.withorb.com/reference/error-responses#404-url-not-found"
	ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses404ResourceNotFound          ErrorType = "https://docs.withorb.com/reference/error-responses#404-resource-not-found"
	ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses404FeatureNotAvailable       ErrorType = "https://docs.withorb.com/reference/error-responses#404-feature-not-available"
	ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses401AuthenticationError       ErrorType = "https://docs.withorb.com/reference/error-responses#401-authentication-error"
	ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses400RequestValidationErrors   ErrorType = "https://docs.withorb.com/reference/error-responses#400-request-validation-errors"
	ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses400DuplicateResourceCreation ErrorType = "https://docs.withorb.com/reference/error-responses#400-duplicate-resource-creation"
	ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses400ConstraintViolation       ErrorType = "https://docs.withorb.com/reference/error-responses#400-constraint-violation"
)

func (r ErrorType) IsKnown() bool {
	switch r {
	case ErrorTypeOrbInternalServerError, ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses429TooManyRequests, ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses413ResourceTooLarge, ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses413RequestTooLarge, ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses409ResourceConflict, ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses404URLNotFound, ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses404ResourceNotFound, ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses404FeatureNotAvailable, ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses401AuthenticationError, ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses400RequestValidationErrors, ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses400DuplicateResourceCreation, ErrorTypeHTTPSDocsWithorbComReferenceErrorResponses400ConstraintViolation:
		return true
	}
	return false
}
