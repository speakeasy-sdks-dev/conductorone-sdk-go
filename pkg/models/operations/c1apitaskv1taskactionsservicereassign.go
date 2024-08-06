// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APITaskV1TaskActionsServiceReassignRequest struct {
	TaskActionsServiceReassignRequest *shared.TaskActionsServiceReassignRequest `request:"mediaType=application/json"`
	TaskID                            string                                    `pathParam:"style=simple,explode=false,name=task_id"`
}

func (o *C1APITaskV1TaskActionsServiceReassignRequest) GetTaskActionsServiceReassignRequest() *shared.TaskActionsServiceReassignRequest {
	if o == nil {
		return nil
	}
	return o.TaskActionsServiceReassignRequest
}

func (o *C1APITaskV1TaskActionsServiceReassignRequest) GetTaskID() string {
	if o == nil {
		return ""
	}
	return o.TaskID
}

type C1APITaskV1TaskActionsServiceReassignResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The TaskActionsServiceReassignResponse returns a task view with paths indicating the location of expanded items in the array.
	TaskActionsServiceReassignResponse *shared.TaskActionsServiceReassignResponse
}

func (o *C1APITaskV1TaskActionsServiceReassignResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APITaskV1TaskActionsServiceReassignResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APITaskV1TaskActionsServiceReassignResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APITaskV1TaskActionsServiceReassignResponse) GetTaskActionsServiceReassignResponse() *shared.TaskActionsServiceReassignResponse {
	if o == nil {
		return nil
	}
	return o.TaskActionsServiceReassignResponse
}
