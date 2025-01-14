---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_group_resource_list Data Source - terraform-provider-opal"
subcategory: ""
description: |-
  GroupResourceList DataSource
---

# opal_group_resource_list (Data Source)

GroupResourceList DataSource

## Example Usage

```terraform
data "opal_group_resource_list" "my_groupresourcelist" {
  group_id = "0c1ba08b-7f3a-43e6-96b6-f0baa3e2cf50"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_id` (String) The ID of the group.

### Read-Only

- `group_resources` (Attributes List) (see [below for nested schema](#nestedatt--group_resources))

<a id="nestedatt--group_resources"></a>
### Nested Schema for `group_resources`

Read-Only:

- `access_level` (Attributes) # Access Level Object
### Description
The `ResourceAccessLevel` object is used to represent the level of access that a user has to a resource or a resource has to a group. The "default" access
level is a `ResourceAccessLevel` object whose fields are all empty strings.

### Usage Example
View the `ResourceAccessLevel` of a resource/user or resource/group pair to see the level of access granted to the resource. (see [below for nested schema](#nestedatt--group_resources--access_level))
- `group_id` (String) The ID of the group.
- `resource_id` (String) The ID of the resource.

<a id="nestedatt--group_resources--access_level"></a>
### Nested Schema for `group_resources.access_level`

Read-Only:

- `access_level_name` (String) The human-readable name of the access level.
- `access_level_remote_id` (String) The machine-readable identifier of the access level.


