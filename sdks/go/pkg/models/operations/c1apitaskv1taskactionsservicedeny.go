// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APITaskV1TaskActionsServiceDenyRequest struct {
	TaskActionsServiceDenyRequest *shared.TaskActionsServiceDenyRequest `request:"mediaType=application/json"`
	TaskID                        string                                `pathParam:"style=simple,explode=false,name=task_id"`
}

type C1APITaskV1TaskActionsServiceDenyResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The TaskActionsServiceDenyResponse returns a task view with paths indicating the location of expanded items in the array.
	TaskActionsServiceDenyResponse *shared.TaskActionsServiceDenyResponse
}
