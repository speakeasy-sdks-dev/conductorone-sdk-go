// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The ResponseTest message.
type ResponseTest struct {
	// version contains the constant value "v1". Future versions of the Webhook Response
	//  will use a different string.
	Version *string `json:"version,omitempty"`
}

func (o *ResponseTest) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}
