// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppsGetRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIAppV1AppsGetResponse struct {
	ContentType string
	// The GetAppResponse message contains the details of the requested app in the app field.
	GetAppResponse *shared.GetAppResponse
	StatusCode     int
	RawResponse    *http.Response
}
