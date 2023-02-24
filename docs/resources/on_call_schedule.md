---
page_title: "opal_on_call_schedule Resource - terraform-provider-opal"
subcategory: ""
description: |-
  An Opal OnCallSchedule resource.
---

# opal_on_call_schedule (Resource)

An Opal OnCallSchedule resource.

## Example Usage

```terraform
resource "opal_on_call_schedule" "security_oncall_rotation" {
  third_party_provider = "PAGER_DUTY"
  remote_id            = "PNXHVAA"
}

# Example group usage
resource "opal_group" "security" {
  // ...

  on_call_schedule {
    id = opal_on_call_schedule.security_oncall_rotation.id
  }

  // or if an UUID is already present in Opal
  on_call_schedule {
    id = "878ba05b-33f0-4dd5-a199-09efc06abcf7"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `remote_id` (String) The remote ID of the on call schedule.
- `third_party_provider` (String) The provider of the on call schedule (i.e. PAGER_DUTY, OPSGENIE).

### Read-Only

- `id` (String) The ID of the on call schedule.

Please [file a ticket](https://github.com/opalsecurity/terraform-provider-opal/issues) to discuss use cases that are not yet supported in the provider.