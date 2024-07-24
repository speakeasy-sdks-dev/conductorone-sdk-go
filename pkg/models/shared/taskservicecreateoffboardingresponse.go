// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
)

// TaskServiceCreateOffboardingResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type TaskServiceCreateOffboardingResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string        `json:"@type,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

func (t TaskServiceCreateOffboardingResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskServiceCreateOffboardingResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskServiceCreateOffboardingResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *TaskServiceCreateOffboardingResponseExpanded) GetAdditionalProperties() map[string]any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The TaskServiceCreateOffboardingResponse message.
type TaskServiceCreateOffboardingResponse struct {
	// Contains a task and JSONPATH expressions that describe where in the expanded array related objects are located. This view can be used to display a fully-detailed dashboard of task information.
	TaskView *TaskView `json:"taskView,omitempty"`
	// The expanded field.
	Expanded []TaskServiceCreateOffboardingResponseExpanded `json:"expanded,omitempty"`
}

func (o *TaskServiceCreateOffboardingResponse) GetTaskView() *TaskView {
	if o == nil {
		return nil
	}
	return o.TaskView
}

func (o *TaskServiceCreateOffboardingResponse) GetExpanded() []TaskServiceCreateOffboardingResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}
