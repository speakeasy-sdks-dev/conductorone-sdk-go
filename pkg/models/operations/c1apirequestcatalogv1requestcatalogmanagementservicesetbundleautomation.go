// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceSetBundleAutomationRequest struct {
	SetBundleAutomationRequest *shared.SetBundleAutomationRequest `request:"mediaType=application/json"`
	RequestCatalogID           string                             `pathParam:"style=simple,explode=false,name=request_catalog_id"`
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceSetBundleAutomationRequest) GetSetBundleAutomationRequest() *shared.SetBundleAutomationRequest {
	if o == nil {
		return nil
	}
	return o.SetBundleAutomationRequest
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceSetBundleAutomationRequest) GetRequestCatalogID() string {
	if o == nil {
		return ""
	}
	return o.RequestCatalogID
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceSetBundleAutomationResponse struct {
	// Successful response
	BundleAutomation *shared.BundleAutomation
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceSetBundleAutomationResponse) GetBundleAutomation() *shared.BundleAutomation {
	if o == nil {
		return nil
	}
	return o.BundleAutomation
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceSetBundleAutomationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceSetBundleAutomationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceSetBundleAutomationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
