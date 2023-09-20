// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// UserDirectoryStatus - The status of the user in the directory.
type UserDirectoryStatus string

const (
	UserDirectoryStatusUnknown  UserDirectoryStatus = "UNKNOWN"
	UserDirectoryStatusEnabled  UserDirectoryStatus = "ENABLED"
	UserDirectoryStatusDisabled UserDirectoryStatus = "DISABLED"
	UserDirectoryStatusDeleted  UserDirectoryStatus = "DELETED"
)

func (e UserDirectoryStatus) ToPointer() *UserDirectoryStatus {
	return &e
}

func (e *UserDirectoryStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNKNOWN":
		fallthrough
	case "ENABLED":
		fallthrough
	case "DISABLED":
		fallthrough
	case "DELETED":
		*e = UserDirectoryStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UserDirectoryStatus: %v", v)
	}
}

// UserStatus - The status of the user in the system.
type UserStatus string

const (
	UserStatusUnknown  UserStatus = "UNKNOWN"
	UserStatusEnabled  UserStatus = "ENABLED"
	UserStatusDisabled UserStatus = "DISABLED"
	UserStatusDeleted  UserStatus = "DELETED"
)

func (e UserStatus) ToPointer() *UserStatus {
	return &e
}

func (e *UserStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNKNOWN":
		fallthrough
	case "ENABLED":
		fallthrough
	case "DISABLED":
		fallthrough
	case "DELETED":
		*e = UserStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UserStatus: %v", v)
	}
}

// The User object provides all of the details for an user, as well as some configuration.
type User struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The id of the user to whom tasks will be automatically reassigned to.
	DelegatedUserID *string    `json:"delegatedUserId,omitempty"`
	DeletedAt       *time.Time `json:"deletedAt,omitempty"`
	// The department which the user belongs to in the organization.
	Department *string `json:"department,omitempty"`
	// A list of objects mapped based on department attribute mappings configured in the system.
	DepartmentSources []UserAttributeMappingSource `json:"departmentSources,omitempty"`
	// A list of unique ids that represent different directories.
	DirectoryIds []string `json:"directoryIds,omitempty"`
	// The status of the user in the directory.
	DirectoryStatus *UserDirectoryStatus `json:"directoryStatus,omitempty"`
	// A list of objects mapped based on directoryStatus attribute mappings configured in the system.
	DirectoryStatusSources []UserAttributeMappingSource `json:"directoryStatusSources,omitempty"`
	// The display name of the user.
	DisplayName *string `json:"displayName,omitempty"`
	// This is the user's email.
	Email *string `json:"email,omitempty"`
	// The users employment status.
	EmploymentStatus *string `json:"employmentStatus,omitempty"`
	// A list of objects mapped based on employmentStatus attribute mappings configured in the system.
	EmploymentStatusSources []UserAttributeMappingSource `json:"employmentStatusSources,omitempty"`
	// The employment type of the user.
	EmploymentType *string `json:"employmentType,omitempty"`
	// A list of objects mapped based on employmentType attribute mappings configured in the system.
	EmploymentTypeSources []UserAttributeMappingSource `json:"employmentTypeSources,omitempty"`
	// A unique identifier of the user.
	ID *string `json:"id,omitempty"`
	// The job title of the user.
	JobTitle *string `json:"jobTitle,omitempty"`
	// A list of objects mapped based on jobTitle attribute mappings configured in the system.
	JobTitleSources []UserAttributeMappingSource `json:"jobTitleSources,omitempty"`
	// A list of ids of the user's managers.
	ManagerIds []string `json:"managerIds,omitempty"`
	// A list of objects mapped based on managerId attribute mappings configured in the system.
	ManagerSources []UserAttributeMappingSource `json:"managerSources,omitempty"`
	// A list of unique identifiers that maps to ConductorOne’s user roles let you assign users permissions tailored to the work they do in the software.
	RoleIds []string `json:"roleIds,omitempty"`
	// The status of the user in the system.
	Status    *UserStatus `json:"status,omitempty"`
	UpdatedAt *time.Time  `json:"updatedAt,omitempty"`
}

func (o *User) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *User) GetDelegatedUserID() *string {
	if o == nil {
		return nil
	}
	return o.DelegatedUserID
}

func (o *User) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *User) GetDepartment() *string {
	if o == nil {
		return nil
	}
	return o.Department
}

func (o *User) GetDepartmentSources() []UserAttributeMappingSource {
	if o == nil {
		return nil
	}
	return o.DepartmentSources
}

func (o *User) GetDirectoryIds() []string {
	if o == nil {
		return nil
	}
	return o.DirectoryIds
}

func (o *User) GetDirectoryStatus() *UserDirectoryStatus {
	if o == nil {
		return nil
	}
	return o.DirectoryStatus
}

func (o *User) GetDirectoryStatusSources() []UserAttributeMappingSource {
	if o == nil {
		return nil
	}
	return o.DirectoryStatusSources
}

func (o *User) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *User) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *User) GetEmploymentStatus() *string {
	if o == nil {
		return nil
	}
	return o.EmploymentStatus
}

func (o *User) GetEmploymentStatusSources() []UserAttributeMappingSource {
	if o == nil {
		return nil
	}
	return o.EmploymentStatusSources
}

func (o *User) GetEmploymentType() *string {
	if o == nil {
		return nil
	}
	return o.EmploymentType
}

func (o *User) GetEmploymentTypeSources() []UserAttributeMappingSource {
	if o == nil {
		return nil
	}
	return o.EmploymentTypeSources
}

func (o *User) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *User) GetJobTitle() *string {
	if o == nil {
		return nil
	}
	return o.JobTitle
}

func (o *User) GetJobTitleSources() []UserAttributeMappingSource {
	if o == nil {
		return nil
	}
	return o.JobTitleSources
}

func (o *User) GetManagerIds() []string {
	if o == nil {
		return nil
	}
	return o.ManagerIds
}

func (o *User) GetManagerSources() []UserAttributeMappingSource {
	if o == nil {
		return nil
	}
	return o.ManagerSources
}

func (o *User) GetRoleIds() []string {
	if o == nil {
		return nil
	}
	return o.RoleIds
}

func (o *User) GetStatus() *UserStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *User) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
