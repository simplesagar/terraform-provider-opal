---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_tag Data Source - terraform-provider-opal"
subcategory: ""
description: |-
  Tag DataSource
---

# opal_tag (Data Source)

Tag DataSource

## Example Usage

```terraform
data "opal_tag" "my_tag" {
  id = "37bd5696-00ed-4435-9a92-039280c1b7e2"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The tag ID

### Read-Only

- `created_at` (String) The date the tag was created.
- `key` (String) The key of the tag.
- `updated_at` (String) The date the tag was last updated.
- `user_creator_id` (String) The ID of the user that created the tag.
- `value` (String) The value of the tag.

