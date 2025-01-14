// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// # UpdateGroupInfo Object
// ### Description
// The `UpdateGroupInfo` object is used as an input to the UpdateGroup API.
type UpdateGroupInfo struct {
	// The ID of the owner of the group.
	AdminOwnerID *string `json:"admin_owner_id,omitempty"`
	// A description of the group.
	Description *string `json:"description,omitempty"`
	// The ID of the group.
	ID *string `json:"group_id,omitempty"`
	// The name of the group.
	Name *string `json:"name,omitempty"`
	// The request configuration list of the configuration template. If not provided, the default request configuration will be used.
	RequestConfigurations []RequestConfiguration `json:"request_configurations"`
	// A bool representing whether or not to require MFA for reviewers to approve requests for this group.
	RequireMfaToApprove *bool `json:"require_mfa_to_approve,omitempty"`
}

func (o *UpdateGroupInfo) GetAdminOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.AdminOwnerID
}

func (o *UpdateGroupInfo) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateGroupInfo) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateGroupInfo) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateGroupInfo) GetRequestConfigurations() []RequestConfiguration {
	if o == nil {
		return []RequestConfiguration{}
	}
	return o.RequestConfigurations
}

func (o *UpdateGroupInfo) GetRequireMfaToApprove() *bool {
	if o == nil {
		return nil
	}
	return o.RequireMfaToApprove
}
