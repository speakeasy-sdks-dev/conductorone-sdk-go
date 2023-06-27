// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AppResourceTypeServiceListResponse - The AppResourceTypeServiceListResponse message.
type AppResourceTypeServiceListResponse struct {
	// The expanded field.
	Expanded []map[string]interface{} `json:"expanded,omitempty"`
	// The list field.
	List []AppResourceTypeView `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// The notificationToken field.
	NotificationToken *string `json:"notificationToken,omitempty"`
}
