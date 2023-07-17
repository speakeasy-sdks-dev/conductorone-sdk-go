// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// PolicyPolicyType - The policyType field.
type PolicyPolicyType string

const (
	PolicyPolicyTypePolicyTypeUnspecified   PolicyPolicyType = "POLICY_TYPE_UNSPECIFIED"
	PolicyPolicyTypePolicyTypeGrant         PolicyPolicyType = "POLICY_TYPE_GRANT"
	PolicyPolicyTypePolicyTypeRevoke        PolicyPolicyType = "POLICY_TYPE_REVOKE"
	PolicyPolicyTypePolicyTypeCertify       PolicyPolicyType = "POLICY_TYPE_CERTIFY"
	PolicyPolicyTypePolicyTypeAccessRequest PolicyPolicyType = "POLICY_TYPE_ACCESS_REQUEST"
	PolicyPolicyTypePolicyTypeProvision     PolicyPolicyType = "POLICY_TYPE_PROVISION"
)

func (e PolicyPolicyType) ToPointer() *PolicyPolicyType {
	return &e
}

func (e *PolicyPolicyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "POLICY_TYPE_UNSPECIFIED":
		fallthrough
	case "POLICY_TYPE_GRANT":
		fallthrough
	case "POLICY_TYPE_REVOKE":
		fallthrough
	case "POLICY_TYPE_CERTIFY":
		fallthrough
	case "POLICY_TYPE_ACCESS_REQUEST":
		fallthrough
	case "POLICY_TYPE_PROVISION":
		*e = PolicyPolicyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PolicyPolicyType: %v", v)
	}
}

// Policy - The Policy message.
type Policy struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The id field.
	ID *string `json:"id,omitempty"`
	// The policySteps field.
	PolicySteps map[string]PolicySteps `json:"policySteps,omitempty"`
	// The policyType field.
	PolicyType *PolicyPolicyType `json:"policyType,omitempty"`
	// The postActions field.
	PostActions []PolicyPostActions `json:"postActions,omitempty"`
	// The reassignTasksToDelegates field.
	ReassignTasksToDelegates *bool `json:"reassignTasksToDelegates,omitempty"`
	// The systemBuiltin field.
	SystemBuiltin *bool      `json:"systemBuiltin,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
}

func (o *Policy) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Policy) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *Policy) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Policy) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *Policy) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Policy) GetPolicySteps() map[string]PolicySteps {
	if o == nil {
		return nil
	}
	return o.PolicySteps
}

func (o *Policy) GetPolicyType() *PolicyPolicyType {
	if o == nil {
		return nil
	}
	return o.PolicyType
}

func (o *Policy) GetPostActions() []PolicyPostActions {
	if o == nil {
		return nil
	}
	return o.PostActions
}

func (o *Policy) GetReassignTasksToDelegates() *bool {
	if o == nil {
		return nil
	}
	return o.ReassignTasksToDelegates
}

func (o *Policy) GetSystemBuiltin() *bool {
	if o == nil {
		return nil
	}
	return o.SystemBuiltin
}

func (o *Policy) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
