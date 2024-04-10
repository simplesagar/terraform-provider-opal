// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &UARSListDataSource{}
var _ datasource.DataSourceWithConfigure = &UARSListDataSource{}

func NewUARSListDataSource() datasource.DataSource {
	return &UARSListDataSource{}
}

// UARSListDataSource is the data source implementation.
type UARSListDataSource struct {
	client *sdk.OpalAPI
}

// UARSListDataSourceModel describes the data model.
type UARSListDataSourceModel struct {
	Cursor   types.String  `tfsdk:"cursor"`
	Next     types.String  `tfsdk:"next"`
	PageSize types.Int64   `tfsdk:"page_size"`
	Previous types.String  `tfsdk:"previous"`
	Results  []tfTypes.Uar `tfsdk:"results"`
}

// Metadata returns the data source type name.
func (r *UARSListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_uars_list"
}

// Schema defines the schema for the data source.
func (r *UARSListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "UARSList DataSource",

		Attributes: map[string]schema.Attribute{
			"cursor": schema.StringAttribute{
				Optional:    true,
				Description: `The pagination cursor value.`,
			},
			"next": schema.StringAttribute{
				Computed:    true,
				Description: `The cursor with which to continue pagination if additional result pages exist.`,
			},
			"page_size": schema.Int64Attribute{
				Optional:    true,
				Description: `Number of results to return per page. Default is 200.`,
			},
			"previous": schema.StringAttribute{
				Computed:    true,
				Description: `The cursor used to obtain the current result page.`,
			},
			"results": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"deadline": schema.StringAttribute{
							Computed:    true,
							Description: `The last day for reviewers to complete their access reviews.`,
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: `The name of the UAR.`,
						},
						"reviewer_assignment_policy": schema.StringAttribute{
							Computed:    true,
							Description: `A policy for auto-assigning reviewers. If auto-assignment is on, specific assignments can still be manually adjusted after the access review is started. Default is Manually. must be one of ["MANUALLY", "BY_OWNING_TEAM_ADMIN", "BY_MANAGER"]`,
						},
						"self_review_allowed": schema.BoolAttribute{
							Computed:    true,
							Description: `A bool representing whether to present a warning when a user is the only reviewer for themself. Default is False.`,
						},
						"send_reviewer_assignment_notification": schema.BoolAttribute{
							Computed:    true,
							Description: `A bool representing whether to send a notification to reviewers when they're assigned a new review. Default is False.`,
						},
						"time_zone": schema.StringAttribute{
							Computed:    true,
							Description: `The time zone name (as defined by the IANA Time Zone database) used in the access review deadline and exported audit report. Default is America/Los_Angeles.`,
						},
						"uar_id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the UAR.`,
						},
						"uar_scope": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"admins": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `This access review will include resources and groups who are owned by one of the owners corresponding to the given IDs.`,
								},
								"names": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `This access review will include resources and groups whose name contains one of the given strings.`,
								},
								"tags": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Computed:    true,
												Description: `The key of the tag.`,
											},
											"value": schema.StringAttribute{
												Computed:    true,
												Description: `The value of the tag.`,
											},
										},
									},
									Description: `This access review will include resources and groups who are tagged with one of the given tags.`,
								},
							},
							Description: `If set, the access review will only contain resources and groups that match at least one of the filters in scope.`,
						},
					},
				},
			},
		},
	}
}

func (r *UARSListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.OpalAPI)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.OpalAPI, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *UARSListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *UARSListDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	cursor := new(string)
	if !data.Cursor.IsUnknown() && !data.Cursor.IsNull() {
		*cursor = data.Cursor.ValueString()
	} else {
		cursor = nil
	}
	pageSize := new(int64)
	if !data.PageSize.IsUnknown() && !data.PageSize.IsNull() {
		*pageSize = data.PageSize.ValueInt64()
	} else {
		pageSize = nil
	}
	request := operations.GetUARsRequest{
		Cursor:   cursor,
		PageSize: pageSize,
	}
	res, err := r.client.Uars.Get(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.PaginatedUARsList == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPaginatedUARsList(res.PaginatedUARsList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
