---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "squadcast_runbook Data Source - terraform-provider-squadcast"
subcategory: ""
description: |-
  A Runbook is a compilation of routine procedures and operations that are documented for reference while working on a critical incident. Sometimes, it can also be referred to as a Playbook.Use this data source to get information about a specific Runbook that you can use for other Squadcast resources.
---

# squadcast_runbook (Data Source)

A Runbook is a compilation of routine procedures and operations that are documented for reference while working on a critical incident. Sometimes, it can also be referred to as a Playbook.Use this data source to get information about a specific Runbook that you can use for other Squadcast resources.

## Example Usage

```terraform
data "squadcast_runbook" "test" {
  name    = squadcast_runbook.test.name
  team_id = "teamObjectId"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the Runbook
- `team_id` (String) Team id.

### Read-Only

- `id` (String) Runbook id.
- `steps` (List of Object) Step by Step instructions, you can add as many steps as you want, supports markdown formatting. (see [below for nested schema](#nestedatt--steps))

<a id="nestedatt--steps"></a>
### Nested Schema for `steps`

Read-Only:

- `content` (String)


