// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppsCreateResponse struct {
	ContentType string
	// Returns the new app's values.
	CreateAppResponse *shared.CreateAppResponse
	StatusCode        int
	RawResponse       *http.Response
}
