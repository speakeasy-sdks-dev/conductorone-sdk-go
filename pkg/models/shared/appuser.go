// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
	"time"
)

// AppUserAppUserType - The appplication user type. Type can be user, system or service.
type AppUserAppUserType string

const (
	AppUserAppUserTypeAppUserTypeUnspecified    AppUserAppUserType = "APP_USER_TYPE_UNSPECIFIED"
	AppUserAppUserTypeAppUserTypeUser           AppUserAppUserType = "APP_USER_TYPE_USER"
	AppUserAppUserTypeAppUserTypeServiceAccount AppUserAppUserType = "APP_USER_TYPE_SERVICE_ACCOUNT"
	AppUserAppUserTypeAppUserTypeSystemAccount  AppUserAppUserType = "APP_USER_TYPE_SYSTEM_ACCOUNT"
)

func (e AppUserAppUserType) ToPointer() *AppUserAppUserType {
	return &e
}

func (e *AppUserAppUserType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "APP_USER_TYPE_UNSPECIFIED":
		fallthrough
	case "APP_USER_TYPE_USER":
		fallthrough
	case "APP_USER_TYPE_SERVICE_ACCOUNT":
		fallthrough
	case "APP_USER_TYPE_SYSTEM_ACCOUNT":
		*e = AppUserAppUserType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AppUserAppUserType: %v", v)
	}
}

// AppUserInput - Application User that represents an account in the application.
type AppUserInput struct {
	// The satus of the applicaiton user.
	AppUserStatus *AppUserStatusInput `json:"status,omitempty"`
	// The appplication user type. Type can be user, system or service.
	AppUserType *AppUserAppUserType `json:"appUserType,omitempty"`
}

func (o *AppUserInput) GetAppUserStatus() *AppUserStatusInput {
	if o == nil {
		return nil
	}
	return o.AppUserStatus
}

func (o *AppUserInput) GetAppUserType() *AppUserAppUserType {
	if o == nil {
		return nil
	}
	return o.AppUserType
}

type AppUserProfile3 struct {
}

type AppUserProfileType string

const (
	AppUserProfileTypeStr             AppUserProfileType = "str"
	AppUserProfileTypeNumber          AppUserProfileType = "number"
	AppUserProfileTypeAppUserProfile3 AppUserProfileType = "AppUser_profile_3"
	AppUserProfileTypeArrayOfany      AppUserProfileType = "arrayOfany"
	AppUserProfileTypeBoolean         AppUserProfileType = "boolean"
)

type AppUserProfile struct {
	Str             *string
	Number          *float64
	AppUserProfile3 *AppUserProfile3
	ArrayOfany      []interface{}
	Boolean         *bool

	Type AppUserProfileType
}

func CreateAppUserProfileStr(str string) AppUserProfile {
	typ := AppUserProfileTypeStr

	return AppUserProfile{
		Str:  &str,
		Type: typ,
	}
}

func CreateAppUserProfileNumber(number float64) AppUserProfile {
	typ := AppUserProfileTypeNumber

	return AppUserProfile{
		Number: &number,
		Type:   typ,
	}
}

func CreateAppUserProfileAppUserProfile3(appUserProfile3 AppUserProfile3) AppUserProfile {
	typ := AppUserProfileTypeAppUserProfile3

	return AppUserProfile{
		AppUserProfile3: &appUserProfile3,
		Type:            typ,
	}
}

func CreateAppUserProfileArrayOfany(arrayOfany []interface{}) AppUserProfile {
	typ := AppUserProfileTypeArrayOfany

	return AppUserProfile{
		ArrayOfany: arrayOfany,
		Type:       typ,
	}
}

func CreateAppUserProfileBoolean(boolean bool) AppUserProfile {
	typ := AppUserProfileTypeBoolean

	return AppUserProfile{
		Boolean: &boolean,
		Type:    typ,
	}
}

func (u *AppUserProfile) UnmarshalJSON(data []byte) error {

	appUserProfile3 := new(AppUserProfile3)
	if err := utils.UnmarshalJSON(data, &appUserProfile3, "", true, true); err == nil {
		u.AppUserProfile3 = appUserProfile3
		u.Type = AppUserProfileTypeAppUserProfile3
		return nil
	}

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = AppUserProfileTypeStr
		return nil
	}

	number := new(float64)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = number
		u.Type = AppUserProfileTypeNumber
		return nil
	}

	arrayOfany := []interface{}{}
	if err := utils.UnmarshalJSON(data, &arrayOfany, "", true, true); err == nil {
		u.ArrayOfany = arrayOfany
		u.Type = AppUserProfileTypeArrayOfany
		return nil
	}

	boolean := new(bool)
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = boolean
		u.Type = AppUserProfileTypeBoolean
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u AppUserProfile) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	if u.AppUserProfile3 != nil {
		return utils.MarshalJSON(u.AppUserProfile3, "", true)
	}

	if u.ArrayOfany != nil {
		return utils.MarshalJSON(u.ArrayOfany, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// AppUser - Application User that represents an account in the application.
type AppUser struct {
	// The satus of the applicaiton user.
	AppUserStatus *AppUserStatus `json:"status,omitempty"`
	// The ID of the application.
	AppID *string `json:"appId,omitempty"`
	// The appplication user type. Type can be user, system or service.
	AppUserType *AppUserAppUserType `json:"appUserType,omitempty"`
	CreatedAt   *time.Time          `json:"createdAt,omitempty"`
	DeletedAt   *time.Time          `json:"deletedAt,omitempty"`
	// The display name of the application user.
	DisplayName *string `json:"displayName,omitempty"`
	// The email field of the application user.
	Email *string `json:"email,omitempty"`
	// The emails field of the application user.
	Emails []string `json:"emails,omitempty"`
	// A unique idenditfier of the application user.
	ID *string `json:"id,omitempty"`
	// The conductor one user ID of the account owner.
	IdentityUserID *string                   `json:"identityUserId,omitempty"`
	Profile        map[string]AppUserProfile `json:"profile,omitempty"`
	UpdatedAt      *time.Time                `json:"updatedAt,omitempty"`
	// The username field of the application user.
	Username *string `json:"username,omitempty"`
	// The usernames field of the application user.
	Usernames []string `json:"usernames,omitempty"`
}

func (a AppUser) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppUser) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppUser) GetAppUserStatus() *AppUserStatus {
	if o == nil {
		return nil
	}
	return o.AppUserStatus
}

func (o *AppUser) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *AppUser) GetAppUserType() *AppUserAppUserType {
	if o == nil {
		return nil
	}
	return o.AppUserType
}

func (o *AppUser) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AppUser) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *AppUser) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *AppUser) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *AppUser) GetEmails() []string {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *AppUser) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AppUser) GetIdentityUserID() *string {
	if o == nil {
		return nil
	}
	return o.IdentityUserID
}

func (o *AppUser) GetProfile() map[string]AppUserProfile {
	if o == nil {
		return nil
	}
	return o.Profile
}

func (o *AppUser) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *AppUser) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *AppUser) GetUsernames() []string {
	if o == nil {
		return nil
	}
	return o.Usernames
}
