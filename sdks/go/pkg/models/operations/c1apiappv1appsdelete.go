// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppsDeleteRequest struct {
	DeleteAppRequest *shared.DeleteAppRequest `request:"mediaType=application/json"`
	ID               string                   `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIAppV1AppsDeleteRequest) GetDeleteAppRequest() *shared.DeleteAppRequest {
	if o == nil {
		return nil
	}
	return o.DeleteAppRequest
}

func (o *C1APIAppV1AppsDeleteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIAppV1AppsDeleteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Empty response body. Status code indicates success.
	DeleteAppResponse *shared.DeleteAppResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAppV1AppsDeleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppsDeleteResponse) GetDeleteAppResponse() *shared.DeleteAppResponse {
	if o == nil {
		return nil
	}
	return o.DeleteAppResponse
}

func (o *C1APIAppV1AppsDeleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppsDeleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
