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
var _ datasource.DataSource = &ResourceDataSource{}
var _ datasource.DataSourceWithConfigure = &ResourceDataSource{}

func NewResourceDataSource() datasource.DataSource {
	return &ResourceDataSource{}
}

// ResourceDataSource is the data source implementation.
type ResourceDataSource struct {
	client *sdk.OpalAPI
}

// ResourceDataSourceModel describes the data model.
type ResourceDataSourceModel struct {
	AdminOwnerID          types.String                   `tfsdk:"admin_owner_id"`
	AppID                 types.String                   `tfsdk:"app_id"`
	AutoApproval          types.Bool                     `tfsdk:"auto_approval"`
	Description           types.String                   `tfsdk:"description"`
	ID                    types.String                   `tfsdk:"id"`
	IsRequestable         types.Bool                     `tfsdk:"is_requestable"`
	MaxDuration           types.Int64                    `tfsdk:"max_duration"`
	Name                  types.String                   `tfsdk:"name"`
	ParentResourceID      types.String                   `tfsdk:"parent_resource_id"`
	RecommendedDuration   types.Int64                    `tfsdk:"recommended_duration"`
	RemoteInfo            *tfTypes.ResourceRemoteInfo    `tfsdk:"remote_info"`
	RemoteResourceID      types.String                   `tfsdk:"remote_resource_id"`
	RemoteResourceName    types.String                   `tfsdk:"remote_resource_name"`
	RequestConfigurations []tfTypes.RequestConfiguration `tfsdk:"request_configurations"`
	RequestTemplateID     types.String                   `tfsdk:"request_template_id"`
	RequireMfaToApprove   types.Bool                     `tfsdk:"require_mfa_to_approve"`
	RequireMfaToConnect   types.Bool                     `tfsdk:"require_mfa_to_connect"`
	RequireMfaToRequest   types.Bool                     `tfsdk:"require_mfa_to_request"`
	RequireSupportTicket  types.Bool                     `tfsdk:"require_support_ticket"`
	ResourceType          types.String                   `tfsdk:"resource_type"`
}

// Metadata returns the data source type name.
func (r *ResourceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource"
}

// Schema defines the schema for the data source.
func (r *ResourceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource DataSource",

		Attributes: map[string]schema.Attribute{
			"admin_owner_id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the owner of the resource.`,
			},
			"app_id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the app.`,
			},
			"auto_approval": schema.BoolAttribute{
				Computed:    true,
				Description: `A bool representing whether or not to automatically approve requests to this resource.`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: `A description of the resource.`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `The ID of the resource.`,
			},
			"is_requestable": schema.BoolAttribute{
				Computed:    true,
				Description: `A bool representing whether or not to allow access requests to this resource.`,
			},
			"max_duration": schema.Int64Attribute{
				Computed:    true,
				Description: `The maximum duration for which the resource can be requested (in minutes).`,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the resource.`,
			},
			"parent_resource_id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the parent resource.`,
			},
			"recommended_duration": schema.Int64Attribute{
				Computed:    true,
				Description: `The recommended duration for which the resource should be requested (in minutes). -1 represents an indefinite duration.`,
			},
			"remote_info": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"aws_account": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"account_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the AWS account.`,
							},
						},
						Description: `Remote info for AWS account.`,
					},
					"aws_ec2_instance": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"account_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the AWS account. Required for AWS Organizations.`,
							},
							"instance_id": schema.StringAttribute{
								Computed:    true,
								Description: `The instanceId of the EC2 instance.`,
							},
							"region": schema.StringAttribute{
								Computed:    true,
								Description: `The region of the EC2 instance.`,
							},
						},
						Description: `Remote info for AWS EC2 instance.`,
					},
					"aws_eks_cluster": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"account_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the AWS account. Required for AWS Organizations.`,
							},
							"arn": schema.StringAttribute{
								Computed:    true,
								Description: `The ARN of the EKS cluster.`,
							},
						},
						Description: `Remote info for AWS EKS cluster.`,
					},
					"aws_iam_role": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"account_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the AWS account. Required for AWS Organizations.`,
							},
							"arn": schema.StringAttribute{
								Computed:    true,
								Description: `The ARN of the IAM role.`,
							},
						},
						Description: `Remote info for AWS IAM role.`,
					},
					"aws_permission_set": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"account_id": schema.StringAttribute{
								Computed:    true,
								Description: `The ID of an AWS account to which this permission set is provisioned.`,
							},
							"arn": schema.StringAttribute{
								Computed:    true,
								Description: `The ARN of the permission set.`,
							},
						},
						Description: `Remote info for AWS Identity Center permission set.`,
					},
					"aws_rds_instance": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"account_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the AWS account. Required for AWS Organizations.`,
							},
							"instance_id": schema.StringAttribute{
								Computed:    true,
								Description: `The instanceId of the RDS instance.`,
							},
							"region": schema.StringAttribute{
								Computed:    true,
								Description: `The region of the RDS instance.`,
							},
							"resource_id": schema.StringAttribute{
								Computed:    true,
								Description: `The resourceId of the RDS instance.`,
							},
						},
						Description: `Remote info for AWS RDS instance.`,
					},
					"gcp_big_query_dataset": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"dataset_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the dataset.`,
							},
							"project_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the project the dataset is in.`,
							},
						},
						Description: `Remote info for GCP BigQuery Dataset.`,
					},
					"gcp_big_query_table": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"dataset_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the dataset the table is in.`,
							},
							"project_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the project the table is in.`,
							},
							"table_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the table.`,
							},
						},
						Description: `Remote info for GCP BigQuery Table.`,
					},
					"gcp_bucket": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"bucket_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the bucket.`,
							},
						},
						Description: `Remote info for GCP bucket.`,
					},
					"gcp_compute_instance": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"instance_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the instance.`,
							},
							"project_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the project the instance is in.`,
							},
							"zone": schema.StringAttribute{
								Computed:    true,
								Description: `The zone the instance is in.`,
							},
						},
						Description: `Remote info for GCP compute instance.`,
					},
					"gcp_folder": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"folder_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the folder.`,
							},
						},
						Description: `Remote info for GCP folder.`,
					},
					"gcp_gke_cluster": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"cluster_name": schema.StringAttribute{
								Computed:    true,
								Description: `The name of the GKE cluster.`,
							},
						},
						Description: `Remote info for GCP GKE cluster.`,
					},
					"gcp_organization": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"organization_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the organization.`,
							},
						},
						Description: `Remote info for GCP organization.`,
					},
					"gcp_project": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"project_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the project.`,
							},
						},
						Description: `Remote info for GCP project.`,
					},
					"gcp_sql_instance": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"instance_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the SQL instance.`,
							},
							"project_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the project the instance is in.`,
							},
						},
						Description: `Remote info for GCP SQL instance.`,
					},
					"github_repo": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"repo_name": schema.StringAttribute{
								Computed:    true,
								Description: `The name of the repository.`,
							},
						},
						Description: `Remote info for GitHub repository.`,
					},
					"gitlab_project": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"project_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the project.`,
							},
						},
						Description: `Remote info for Gitlab project.`,
					},
					"okta_app": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"app_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the app.`,
							},
						},
						Description: `Remote info for Okta directory app.`,
					},
					"okta_custom_role": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"role_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the custom role.`,
							},
						},
						Description: `Remote info for Okta directory custom role.`,
					},
					"okta_standard_role": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"role_type": schema.StringAttribute{
								Computed:    true,
								Description: `The type of the standard role.`,
							},
						},
						Description: `Remote info for Okta directory standard role.`,
					},
					"pagerduty_role": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"role_name": schema.StringAttribute{
								Computed:    true,
								Description: `The name of the role.`,
							},
						},
						Description: `Remote info for Pagerduty role.`,
					},
					"salesforce_permission_set": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"permission_set_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the permission set.`,
							},
						},
						Description: `Remote info for Salesforce permission set.`,
					},
					"salesforce_profile": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"profile_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the permission set.`,
							},
							"user_license_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the user license.`,
							},
						},
						Description: `Remote info for Salesforce profile.`,
					},
					"salesforce_role": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"role_id": schema.StringAttribute{
								Computed:    true,
								Description: `The id of the role.`,
							},
						},
						Description: `Remote info for Salesforce role.`,
					},
					"teleport_role": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"role_name": schema.StringAttribute{
								Computed:    true,
								Description: `The name role.`,
							},
						},
						Description: `Remote info for Teleport role.`,
					},
				},
				Description: `Information that defines the remote resource. This replaces the deprecated remote_id and metadata fields.`,
			},
			"remote_resource_id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the resource on the remote system.`,
			},
			"remote_resource_name": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the resource on the remote system.`,
			},
			"request_configurations": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allow_requests": schema.BoolAttribute{
							Computed:    true,
							Description: `A bool representing whether or not to allow requests for this resource.`,
						},
						"auto_approval": schema.BoolAttribute{
							Computed:    true,
							Description: `A bool representing whether or not to automatically approve requests for this resource.`,
						},
						"condition": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"group_ids": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `The list of group IDs to match.`,
								},
								"role_remote_ids": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `The list of role remote IDs to match.`,
								},
							},
						},
						"max_duration": schema.Int64Attribute{
							Computed:    true,
							Description: `The maximum duration for which the resource can be requested (in minutes).`,
						},
						"priority": schema.Int64Attribute{
							Computed:    true,
							Description: `The priority of the request configuration.`,
						},
						"recommended_duration": schema.Int64Attribute{
							Computed:    true,
							Description: `The recommended duration for which the resource should be requested (in minutes). -1 represents an indefinite duration.`,
						},
						"request_template_id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the associated request template.`,
						},
						"require_mfa_to_request": schema.BoolAttribute{
							Computed:    true,
							Description: `A bool representing whether or not to require MFA for requesting access to this resource.`,
						},
						"require_support_ticket": schema.BoolAttribute{
							Computed:    true,
							Description: `A bool representing whether or not access requests to the resource require an access ticket.`,
						},
						"reviewer_stages": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"operator": schema.StringAttribute{
										Computed:    true,
										Description: `The operator of the reviewer stage. must be one of ["AND", "OR"]`,
									},
									"owner_ids": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
									"require_manager_approval": schema.BoolAttribute{
										Computed:    true,
										Description: `Whether this reviewer stage should require manager approval.`,
									},
								},
							},
							Description: `The list of reviewer stages for the request configuration.`,
						},
					},
				},
				Description: `A list of configurations for requests to this resource.`,
			},
			"request_template_id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of the associated request template.`,
			},
			"require_mfa_to_approve": schema.BoolAttribute{
				Computed:    true,
				Description: `A bool representing whether or not to require MFA for reviewers to approve requests for this resource.`,
			},
			"require_mfa_to_connect": schema.BoolAttribute{
				Computed:    true,
				Description: `A bool representing whether or not to require MFA to connect to this resource.`,
			},
			"require_mfa_to_request": schema.BoolAttribute{
				Computed:    true,
				Description: `A bool representing whether or not to require MFA for requesting access to this resource.`,
			},
			"require_support_ticket": schema.BoolAttribute{
				Computed:    true,
				Description: `A bool representing whether or not access requests to the resource require an access ticket.`,
			},
			"resource_type": schema.StringAttribute{
				Computed:    true,
				Description: `The type of the resource. must be one of ["AWS_IAM_ROLE", "AWS_EC2_INSTANCE", "AWS_EKS_CLUSTER", "AWS_RDS_POSTGRES_INSTANCE", "AWS_RDS_MYSQL_INSTANCE", "AWS_ACCOUNT", "AWS_SSO_PERMISSION_SET", "CUSTOM", "GCP_BUCKET", "GCP_COMPUTE_INSTANCE", "GCP_FOLDER", "GCP_GKE_CLUSTER", "GCP_PROJECT", "GCP_CLOUD_SQL_POSTGRES_INSTANCE", "GCP_CLOUD_SQL_MYSQL_INSTANCE", "GIT_HUB_REPO", "GIT_LAB_PROJECT", "GOOGLE_WORKSPACE_ROLE", "MONGO_INSTANCE", "MONGO_ATLAS_INSTANCE", "OKTA_APP", "OKTA_ROLE", "OPAL_ROLE", "PAGERDUTY_ROLE", "TAILSCALE_SSH", "SALESFORCE_PERMISSION_SET", "SALESFORCE_PROFILE", "SALESFORCE_ROLE", "WORKDAY_ROLE", "MYSQL_INSTANCE", "MARIADB_INSTANCE", "TELEPORT_ROLE"]`,
			},
		},
	}
}

func (r *ResourceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *ResourceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ResourceDataSourceModel
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

	id := data.ID.ValueString()
	request := operations.GetResourceIDRequest{
		ID: id,
	}
	res, err := r.client.Resources.GetID(ctx, request)
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
	if res.Resource == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedResource(res.Resource)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
