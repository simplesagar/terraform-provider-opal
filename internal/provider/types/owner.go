// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Owner struct {
	AccessRequestEscalationPeriod types.Int64  `tfsdk:"access_request_escalation_period"`
	Description                   types.String `tfsdk:"description"`
	ID                            types.String `tfsdk:"id"`
	Name                          types.String `tfsdk:"name"`
	ReviewerMessageChannelID      types.String `tfsdk:"reviewer_message_channel_id"`
	SourceGroupID                 types.String `tfsdk:"source_group_id"`
}