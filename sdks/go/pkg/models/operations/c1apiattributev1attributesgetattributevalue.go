// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAttributeV1AttributesGetAttributeValueRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIAttributeV1AttributesGetAttributeValueResponse struct {
	ContentType string
	// GetAttributeValueResponse is the response for getting an attribute value by id.
	GetAttributeValueResponse *shared.GetAttributeValueResponse
	StatusCode                int
	RawResponse               *http.Response
}
