---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_apps Data Source - terraform-provider-opal"
subcategory: ""
description: |-
  Apps DataSource
---

# opal_apps (Data Source)

Apps DataSource

## Example Usage

```terraform
data "opal_apps" "my_apps" {
  app_type_filter = [
    "OKTA_DIRECTORY",
  ]
  owner_filter = "d8baf069-fe55-4e23-a367-4e7f570e3140"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `app_type_filter` (List of String) A list of app types to filter by.
- `owner_filter` (String) An owner ID to filter by.

### Read-Only

- `apps` (Attributes List) (see [below for nested schema](#nestedatt--apps))

<a id="nestedatt--apps"></a>
### Nested Schema for `apps`

Read-Only:

- `admin_owner_id` (String) The ID of the owner of the app.
- `description` (String) A description of the app.
- `id` (String) The ID of the app.
- `name` (String) The name of the app.
- `type` (String) The type of an app.


