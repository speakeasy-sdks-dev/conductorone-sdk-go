// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The Approval message.
//
// This message contains a oneof named typ. Only a single field of the following list may be set at a time:
//   - users
//   - manager
//   - appOwners
//   - group
//   - self
//   - entitlementOwners
//   - expression
//   - webhook
type Approval struct {
	// The AppGroupApproval object provides the configuration for setting a group as the approvers of an approval policy step.
	AppGroupApproval *AppGroupApproval `json:"group,omitempty"`
	// App owner approval provides the configuration for an approval step when the app owner is the target.
	AppOwnerApproval *AppOwnerApproval `json:"appOwners,omitempty"`
	// The entitlement owner approval allows configuration of the approval step when the target approvers are the entitlement owners.
	EntitlementOwnerApproval *EntitlementOwnerApproval `json:"entitlementOwners,omitempty"`
	// The ExpressionApproval message.
	ExpressionApproval *ExpressionApproval `json:"expression,omitempty"`
	// The manager approval object provides configuration options for approval when the target of the approval is the manager of the user in the task.
	ManagerApproval *ManagerApproval `json:"manager,omitempty"`
	// The self approval object describes the configuration of a policy step that needs to be approved by the target of the request.
	SelfApproval *SelfApproval `json:"self,omitempty"`
	// The user approval object describes the approval configuration of a policy step that needs to be approved by a specific list of users.
	UserApproval *UserApproval `json:"users,omitempty"`
	// The WebhookApproval message.
	WebhookApproval *WebhookApproval `json:"webhook,omitempty"`
	// Configuration to allow reassignment by reviewers during this step.
	AllowReassignment *bool `json:"allowReassignment,omitempty"`
	// A field indicating whether this step is assigned.
	Assigned *bool `json:"assigned,omitempty"`
	// Configuration to require a reason when approving this step.
	RequireApprovalReason *bool `json:"requireApprovalReason,omitempty"`
	// Configuration to require a reason when denying this step.
	RequireDenialReason *bool `json:"requireDenialReason,omitempty"`
	// Configuration to require a reason when reassigning this step.
	RequireReassignmentReason *bool `json:"requireReassignmentReason,omitempty"`
}

func (o *Approval) GetAppGroupApproval() *AppGroupApproval {
	if o == nil {
		return nil
	}
	return o.AppGroupApproval
}

func (o *Approval) GetAppOwnerApproval() *AppOwnerApproval {
	if o == nil {
		return nil
	}
	return o.AppOwnerApproval
}

func (o *Approval) GetEntitlementOwnerApproval() *EntitlementOwnerApproval {
	if o == nil {
		return nil
	}
	return o.EntitlementOwnerApproval
}

func (o *Approval) GetExpressionApproval() *ExpressionApproval {
	if o == nil {
		return nil
	}
	return o.ExpressionApproval
}

func (o *Approval) GetManagerApproval() *ManagerApproval {
	if o == nil {
		return nil
	}
	return o.ManagerApproval
}

func (o *Approval) GetSelfApproval() *SelfApproval {
	if o == nil {
		return nil
	}
	return o.SelfApproval
}

func (o *Approval) GetUserApproval() *UserApproval {
	if o == nil {
		return nil
	}
	return o.UserApproval
}

func (o *Approval) GetWebhookApproval() *WebhookApproval {
	if o == nil {
		return nil
	}
	return o.WebhookApproval
}

func (o *Approval) GetAllowReassignment() *bool {
	if o == nil {
		return nil
	}
	return o.AllowReassignment
}

func (o *Approval) GetAssigned() *bool {
	if o == nil {
		return nil
	}
	return o.Assigned
}

func (o *Approval) GetRequireApprovalReason() *bool {
	if o == nil {
		return nil
	}
	return o.RequireApprovalReason
}

func (o *Approval) GetRequireDenialReason() *bool {
	if o == nil {
		return nil
	}
	return o.RequireDenialReason
}

func (o *Approval) GetRequireReassignmentReason() *bool {
	if o == nil {
		return nil
	}
	return o.RequireReassignmentReason
}
