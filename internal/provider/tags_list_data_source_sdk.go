// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"time"
)

func (r *TagsListDataSourceModel) RefreshFromSharedPaginatedTagsList(resp *shared.PaginatedTagsList) {
	if resp != nil {
		r.Next = types.StringPointerValue(resp.Next)
		r.Previous = types.StringPointerValue(resp.Previous)
		if len(r.Results) > len(resp.Results) {
			r.Results = r.Results[:len(resp.Results)]
		}
		for resultsCount, resultsItem := range resp.Results {
			var results1 tfTypes.Tag
			if resultsItem.CreatedAt != nil {
				results1.CreatedAt = types.StringValue(resultsItem.CreatedAt.Format(time.RFC3339Nano))
			} else {
				results1.CreatedAt = types.StringNull()
			}
			results1.ID = types.StringPointerValue(resultsItem.ID)
			results1.Key = types.StringPointerValue(resultsItem.Key)
			if resultsItem.UpdatedAt != nil {
				results1.UpdatedAt = types.StringValue(resultsItem.UpdatedAt.Format(time.RFC3339Nano))
			} else {
				results1.UpdatedAt = types.StringNull()
			}
			results1.UserCreatorID = types.StringPointerValue(resultsItem.UserCreatorID)
			results1.Value = types.StringPointerValue(resultsItem.Value)
			if resultsCount+1 > len(r.Results) {
				r.Results = append(r.Results, results1)
			} else {
				r.Results[resultsCount].CreatedAt = results1.CreatedAt
				r.Results[resultsCount].ID = results1.ID
				r.Results[resultsCount].Key = results1.Key
				r.Results[resultsCount].UpdatedAt = results1.UpdatedAt
				r.Results[resultsCount].UserCreatorID = results1.UserCreatorID
				r.Results[resultsCount].Value = results1.Value
			}
		}
	}
}
