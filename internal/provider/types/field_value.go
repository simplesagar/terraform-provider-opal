// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type FieldValue struct {
	Str     types.String `tfsdk:"str" tfPlanOnly:"true"`
	Boolean types.Bool   `tfsdk:"boolean" tfPlanOnly:"true"`
}
