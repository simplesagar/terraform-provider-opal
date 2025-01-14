---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opal_uar Data Source - terraform-provider-opal"
subcategory: ""
description: |-
  Uar DataSource
---

# opal_uar (Data Source)

Uar DataSource

## Example Usage

```terraform
data "opal_uar" "my_uar" {
  uar_id = "83615f59-182c-4e94-a61d-7e2dadcd9e9a"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `uar_id` (String) The ID of the UAR.

### Read-Only

- `deadline` (String) The last day for reviewers to complete their access reviews.
- `name` (String) The name of the UAR.
- `reviewer_assignment_policy` (String) A policy for auto-assigning reviewers. If auto-assignment is on, specific assignments can still be manually adjusted after the access review is started. Default is Manually. must be one of ["MANUALLY", "BY_OWNING_TEAM_ADMIN", "BY_MANAGER"]
- `self_review_allowed` (Boolean) A bool representing whether to present a warning when a user is the only reviewer for themself. Default is False.
- `send_reviewer_assignment_notification` (Boolean) A bool representing whether to send a notification to reviewers when they're assigned a new review. Default is False.
- `time_zone` (String) The time zone name (as defined by the IANA Time Zone database) used in the access review deadline and exported audit report. Default is America/Los_Angeles.
- `uar_scope` (Attributes) If set, the access review will only contain resources and groups that match at least one of the filters in scope. (see [below for nested schema](#nestedatt--uar_scope))

<a id="nestedatt--uar_scope"></a>
### Nested Schema for `uar_scope`

Read-Only:

- `admins` (List of String) This access review will include resources and groups who are owned by one of the owners corresponding to the given IDs.
- `names` (List of String) This access review will include resources and groups whose name contains one of the given strings.
- `tags` (Attributes List) This access review will include resources and groups who are tagged with one of the given tags. (see [below for nested schema](#nestedatt--uar_scope--tags))

<a id="nestedatt--uar_scope--tags"></a>
### Nested Schema for `uar_scope.tags`

Read-Only:

- `key` (String) The key of the tag.
- `value` (String) The value of the tag.


