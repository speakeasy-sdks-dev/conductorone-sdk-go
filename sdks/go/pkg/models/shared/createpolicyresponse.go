// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreatePolicyResponse - The CreatePolicyResponse message contains the created policy object.
type CreatePolicyResponse struct {
	// A policy describes the behavior of the ConductorOne system when processing a task. You can describe the type, approvers, fallback behavior, and escalation processes.
	Policy *Policy `json:"policy,omitempty"`
}
