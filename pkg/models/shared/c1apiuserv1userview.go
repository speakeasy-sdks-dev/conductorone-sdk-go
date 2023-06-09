// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// C1APIUserV1UserView - The UserView message.
type C1APIUserV1UserView struct {
	// The delegatedUserPath field.
	DelegatedUserPath *string `json:"delegatedUserPath,omitempty"`
	// The directoriesPath field.
	DirectoriesPath *string `json:"directoriesPath,omitempty"`
	// The managersPath field.
	ManagersPath *string `json:"managersPath,omitempty"`
	// The rolesPath field.
	RolesPath *string `json:"rolesPath,omitempty"`
	// The User message.
	User *C1APIUserV1User `json:"user,omitempty"`
}
