// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
	"time"
)

// The BundleAutomation message.
//
// This message contains a oneof named conditions. Only a single field of the following list may be set at a time:
//   - entitlements
type BundleAutomation struct {
	// The BundleAutomationLastRunState message.
	BundleAutomationLastRunState *BundleAutomationLastRunState `json:"state,omitempty"`
	// The BundleAutomationRuleEntitlement message.
	BundleAutomationRuleEntitlement *BundleAutomationRuleEntitlement `json:"entitlements,omitempty"`
	CreatedAt                       *time.Time                       `json:"createdAt,omitempty"`
	DeletedAt                       *time.Time                       `json:"deletedAt,omitempty"`
	// The requestCatalogId field.
	RequestCatalogID *string `json:"requestCatalogId,omitempty"`
	// The tenantId field.
	TenantID  *string    `json:"tenantId,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

func (b BundleAutomation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BundleAutomation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BundleAutomation) GetBundleAutomationLastRunState() *BundleAutomationLastRunState {
	if o == nil {
		return nil
	}
	return o.BundleAutomationLastRunState
}

func (o *BundleAutomation) GetBundleAutomationRuleEntitlement() *BundleAutomationRuleEntitlement {
	if o == nil {
		return nil
	}
	return o.BundleAutomationRuleEntitlement
}

func (o *BundleAutomation) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *BundleAutomation) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *BundleAutomation) GetRequestCatalogID() *string {
	if o == nil {
		return nil
	}
	return o.RequestCatalogID
}

func (o *BundleAutomation) GetTenantID() *string {
	if o == nil {
		return nil
	}
	return o.TenantID
}

func (o *BundleAutomation) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
