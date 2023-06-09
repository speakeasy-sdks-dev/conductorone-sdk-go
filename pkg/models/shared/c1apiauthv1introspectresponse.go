// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// C1APIAuthV1IntrospectResponse - The IntrospectResponse message.
type C1APIAuthV1IntrospectResponse struct {
	// The features field.
	Features []string `json:"features,omitempty"`
	// The permissions field.
	Permissions []string `json:"permissions,omitempty"`
	// The principleId field.
	PrincipleID *string `json:"principleId,omitempty"`
	// The roles field.
	Roles []string `json:"roles,omitempty"`
	// The userId field.
	UserID *string `json:"userId,omitempty"`
}
