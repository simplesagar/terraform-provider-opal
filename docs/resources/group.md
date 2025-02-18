---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_group Resource - terraform-provider-opal"
subcategory: ""
description: |-
  Group Resource
---

# opal_group (Resource)

Group Resource

## Example Usage

```terraform
resource "opal_group" "my_group" {
  admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
  app_id         = "f454d283-ca87-4a8a-bdbb-df212eca5353"
  description    = "Engineering team Okta group."
  group_type     = "OPAL_GROUP"
  message_channel_ids = [
    "e931861e-f161-4b54-bb18-98e51c0009dd",
  ]
  name = "mongo-db-prod"
  on_call_schedule_ids = [
    "10c42893-326c-48d3-b654-3ad1053f385d",
  ]
  request_configurations = [
    {
      allow_requests = true
      auto_approval  = false
      condition = {
        group_ids = [
          "ce27d58f-2561-447d-92ea-246933c124e1",
        ]
        role_remote_ids = [
          "...",
        ]
      }
      max_duration           = 120
      priority               = 1
      recommended_duration   = 120
      request_template_id    = "06851574-e50d-40ca-8c78-f72ae6ab4304"
      require_mfa_to_request = false
      require_support_ticket = false
      reviewer_stages = [
        {
          operator = "AND"
          owner_ids = [
            "53ec25e2-8706-4e50-ad65-59d94490f519",
          ]
          require_manager_approval = false
        },
      ]
    },
  ]
  require_mfa_to_approve = false
  visibility             = "GLOBAL"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The ID of the app for the group. Requires replacement if changed.
- `group_type` (String) The type of the group. Requires replacement if changed. ; must be one of ["ACTIVE_DIRECTORY_GROUP", "AWS_SSO_GROUP", "DUO_GROUP", "GIT_HUB_TEAM", "GIT_LAB_GROUP", "GOOGLE_GROUPS_GROUP", "LDAP_GROUP", "OKTA_GROUP", "OPAL_GROUP", "AZURE_AD_SECURITY_GROUP", "AZURE_AD_MICROSOFT_365_GROUP"]
- `message_channel_ids` (List of String)
- `name` (String) The name of the remote group.
- `on_call_schedule_ids` (List of String)
- `request_configurations` (Attributes List) The request configuration list of the configuration template. If not provided, the default request configuration will be used. (see [below for nested schema](#nestedatt--request_configurations))
- `visibility` (String) The visibility level of the entity. must be one of ["GLOBAL", "LIMITED"]

### Optional

- `admin_owner_id` (String) The ID of the owner of the group.
- `description` (String) A description of the remote group.
- `remote_info` (Attributes) Information that defines the remote group. This replaces the deprecated remote_id and metadata fields. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info))
- `require_mfa_to_approve` (Boolean) A bool representing whether or not to require MFA for reviewers to approve requests for this group.
- `visibility_group_ids` (List of String)

### Read-Only

- `group_binding_id` (String) The ID of the associated group binding.
- `id` (String) The ID of the group.
- `message_channels` (Attributes) The audit and reviewer message channels attached to the group. (see [below for nested schema](#nestedatt--message_channels))
- `oncall_schedules` (Attributes) The on call schedules attached to the group. (see [below for nested schema](#nestedatt--oncall_schedules))
- `remote_name` (String) The name of the remote.

<a id="nestedatt--request_configurations"></a>
### Nested Schema for `request_configurations`

Optional:

- `allow_requests` (Boolean) A bool representing whether or not to allow requests for this resource. Not Null
- `auto_approval` (Boolean) A bool representing whether or not to automatically approve requests for this resource. Not Null
- `condition` (Attributes) (see [below for nested schema](#nestedatt--request_configurations--condition))
- `max_duration` (Number) The maximum duration for which the resource can be requested (in minutes).
- `priority` (Number) The priority of the request configuration. Not Null
- `recommended_duration` (Number) The recommended duration for which the resource should be requested (in minutes). -1 represents an indefinite duration.
- `request_template_id` (String) The ID of the associated request template.
- `require_mfa_to_request` (Boolean) A bool representing whether or not to require MFA for requesting access to this resource. Not Null
- `require_support_ticket` (Boolean) A bool representing whether or not access requests to the resource require an access ticket. Not Null
- `reviewer_stages` (Attributes List) The list of reviewer stages for the request configuration. Not Null (see [below for nested schema](#nestedatt--request_configurations--reviewer_stages))

<a id="nestedatt--request_configurations--condition"></a>
### Nested Schema for `request_configurations.condition`

Optional:

- `group_ids` (List of String) The list of group IDs to match.
- `role_remote_ids` (List of String) The list of role remote IDs to match.


<a id="nestedatt--request_configurations--reviewer_stages"></a>
### Nested Schema for `request_configurations.reviewer_stages`

Optional:

- `operator` (String) The operator of the reviewer stage. must be one of ["AND", "OR"]; Default: "AND"
- `owner_ids` (List of String) Not Null
- `require_manager_approval` (Boolean) Whether this reviewer stage should require manager approval. Not Null



<a id="nestedatt--remote_info"></a>
### Nested Schema for `remote_info`

Optional:

- `active_directory_group` (Attributes) Remote info for Active Directory group. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--active_directory_group))
- `azure_ad_microsoft_365_group` (Attributes) Remote info for Azure AD Microsoft 365 group. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--azure_ad_microsoft_365_group))
- `azure_ad_security_group` (Attributes) Remote info for Azure AD Security group. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--azure_ad_security_group))
- `duo_group` (Attributes) Remote info for Duo Security group. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--duo_group))
- `github_team` (Attributes) Remote info for GitHub team. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--github_team))
- `gitlab_group` (Attributes) Remote info for Gitlab group. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--gitlab_group))
- `google_group` (Attributes) Remote info for Google group. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--google_group))
- `ldap_group` (Attributes) Remote info for LDAP group. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--ldap_group))
- `okta_group` (Attributes) Remote info for Okta Directory group. Requires replacement if changed. (see [below for nested schema](#nestedatt--remote_info--okta_group))

<a id="nestedatt--remote_info--active_directory_group"></a>
### Nested Schema for `remote_info.active_directory_group`

Optional:

- `group_id` (String) The id of the Google group. Requires replacement if changed. ; Not Null


<a id="nestedatt--remote_info--azure_ad_microsoft_365_group"></a>
### Nested Schema for `remote_info.azure_ad_microsoft_365_group`

Optional:

- `group_id` (String) The id of the Azure AD Microsoft 365 group. Requires replacement if changed. ; Not Null


<a id="nestedatt--remote_info--azure_ad_security_group"></a>
### Nested Schema for `remote_info.azure_ad_security_group`

Optional:

- `group_id` (String) The id of the Azure AD Security group. Requires replacement if changed. ; Not Null


<a id="nestedatt--remote_info--duo_group"></a>
### Nested Schema for `remote_info.duo_group`

Optional:

- `group_id` (String) The id of the Duo Security group. Requires replacement if changed. ; Not Null


<a id="nestedatt--remote_info--github_team"></a>
### Nested Schema for `remote_info.github_team`

Optional:

- `team_slug` (String) The slug of the GitHub team. Requires replacement if changed. ; Not Null


<a id="nestedatt--remote_info--gitlab_group"></a>
### Nested Schema for `remote_info.gitlab_group`

Optional:

- `group_id` (String) The id of the Gitlab group. Requires replacement if changed. ; Not Null


<a id="nestedatt--remote_info--google_group"></a>
### Nested Schema for `remote_info.google_group`

Optional:

- `group_id` (String) The id of the Google group. Requires replacement if changed. ; Not Null


<a id="nestedatt--remote_info--ldap_group"></a>
### Nested Schema for `remote_info.ldap_group`

Optional:

- `group_id` (String) The id of the LDAP group. Requires replacement if changed. ; Not Null


<a id="nestedatt--remote_info--okta_group"></a>
### Nested Schema for `remote_info.okta_group`

Optional:

- `group_id` (String) The id of the Okta Directory group. Requires replacement if changed. ; Not Null



<a id="nestedatt--message_channels"></a>
### Nested Schema for `message_channels`

Read-Only:

- `channels` (Attributes List) (see [below for nested schema](#nestedatt--message_channels--channels))

<a id="nestedatt--message_channels--channels"></a>
### Nested Schema for `message_channels.channels`

Read-Only:

- `id` (String) The ID of the message channel.
- `is_private` (Boolean) A bool representing whether or not the message channel is private.
- `name` (String) The name of the message channel.
- `remote_id` (String) The remote ID of the message channel
- `third_party_provider` (String) The third party provider of the message channel. must be one of ["SLACK"]



<a id="nestedatt--oncall_schedules"></a>
### Nested Schema for `oncall_schedules`

Read-Only:

- `id` (String) The ID of the on-call schedule.
- `name` (String) The name of the on call schedule.
- `remote_id` (String) The remote ID of the on call schedule
- `third_party_provider` (String) The third party provider of the on call schedule. must be one of ["OPSGENIE", "PAGER_DUTY"]


