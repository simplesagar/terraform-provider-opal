---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_resource_reviewers_list Data Source - terraform-provider-opal"
subcategory: ""
description: |-
  ResourceReviewersList DataSource
---

# opal_resource_reviewers_list (Data Source)

ResourceReviewersList DataSource

## Example Usage

```terraform
data "opal_resource_reviewers_list" "my_resourcereviewerslist" {
  resource_id = "6b9881a6-b52c-41c9-b4d2-5038b6404764"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `resource_id` (String) The ID of the resource.

### Read-Only

- `data` (List of String) The IDs of owners that are reviewers for this resource.

