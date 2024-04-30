---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_users Data Source - terraform-provider-opal"
subcategory: ""
description: |-
  Users DataSource
---

# opal_users (Data Source)

Users DataSource

## Example Usage

```terraform
data "opal_users" "my_users" {
  cursor    = "...my_cursor..."
  page_size = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `cursor` (String) The pagination cursor value.
- `page_size` (Number) Number of results to return per page. Default is 200.

### Read-Only

- `next` (String) The cursor with which to continue pagination if additional result pages exist.
- `previous` (String) The cursor used to obtain the current result page.
- `results` (Attributes List) (see [below for nested schema](#nestedatt--results))

<a id="nestedatt--results"></a>
### Nested Schema for `results`

Read-Only:

- `email` (String) The email of the user.
- `first_name` (String) The first name of the user.
- `hr_idp_status` (String) User status pulled from an HR/IDP provider. must be one of ["ACTIVE", "SUSPENDED", "DEPROVISIONED", "DELETED", "NOT_FOUND"]
- `id` (String) The ID of the user.
- `last_name` (String) The last name of the user.
- `name` (String) The full name of the user.
- `position` (String) The user's position.

