// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ConfigurationTemplate - # Configuration Template Object
// ### Description
// The `ConfigurationTemplate` object is used to represent a configuration template.
//
// ### Usage Example
// Returned from the `GET Configuration Templates` endpoint.
type ConfigurationTemplate struct {
	// The ID of the owner of the configuration template.
	AdminOwnerID *string `json:"admin_owner_id,omitempty"`
	// The IDs of the break glass users linked to the configuration template.
	BreakGlassUserIds []string `json:"break_glass_user_ids,omitempty"`
	// The ID of the configuration template.
	ConfigurationTemplateID *string `json:"configuration_template_id,omitempty"`
	// The IDs of the audit message channels linked to the configuration template.
	LinkedAuditMessageChannelIds []string `json:"linked_audit_message_channel_ids,omitempty"`
	// The IDs of the on-call schedules linked to the configuration template.
	MemberOncallScheduleIds []string `json:"member_oncall_schedule_ids,omitempty"`
	// The name of the configuration template.
	Name *string `json:"name,omitempty"`
	// The ID of the request configuration linked to the configuration template.
	RequestConfigurationID *string `json:"request_configuration_id,omitempty"`
	// A bool representing whether or not to require MFA for reviewers to approve requests for this configuration template.
	RequireMfaToApprove *bool `json:"require_mfa_to_approve,omitempty"`
	// A bool representing whether or not to require MFA to connect to resources associated with this configuration template.
	RequireMfaToConnect *bool `json:"require_mfa_to_connect,omitempty"`
	// Visibility infomation of an entity.
	Visibility *VisibilityInfo `json:"visibility,omitempty"`
}

func (o *ConfigurationTemplate) GetAdminOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.AdminOwnerID
}

func (o *ConfigurationTemplate) GetBreakGlassUserIds() []string {
	if o == nil {
		return nil
	}
	return o.BreakGlassUserIds
}

func (o *ConfigurationTemplate) GetConfigurationTemplateID() *string {
	if o == nil {
		return nil
	}
	return o.ConfigurationTemplateID
}

func (o *ConfigurationTemplate) GetLinkedAuditMessageChannelIds() []string {
	if o == nil {
		return nil
	}
	return o.LinkedAuditMessageChannelIds
}

func (o *ConfigurationTemplate) GetMemberOncallScheduleIds() []string {
	if o == nil {
		return nil
	}
	return o.MemberOncallScheduleIds
}

func (o *ConfigurationTemplate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConfigurationTemplate) GetRequestConfigurationID() *string {
	if o == nil {
		return nil
	}
	return o.RequestConfigurationID
}

func (o *ConfigurationTemplate) GetRequireMfaToApprove() *bool {
	if o == nil {
		return nil
	}
	return o.RequireMfaToApprove
}

func (o *ConfigurationTemplate) GetRequireMfaToConnect() *bool {
	if o == nil {
		return nil
	}
	return o.RequireMfaToConnect
}

func (o *ConfigurationTemplate) GetVisibility() *VisibilityInfo {
	if o == nil {
		return nil
	}
	return o.Visibility
}