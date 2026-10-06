---
page_title: "xcsh_global_log_receiver examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver examples."
---

# xcsh_global_log_receiver examples

<a id="canonical-3023311232200010-0133300122022312-3102212013022001-1130103102302322-3213013111200203-3211331221323312-0213002213312012-2002030133101203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- Examples

<a id="canonical-0311331101322131-0101102200311003-0111112032201201-1331310300230301-2000010013112223-2333321223233323-0331300212311322-1100301301122012"></a>

### Complete configurations for `xcsh_global_log_receiver`

- [Data source](data-sources--global_log_receiver--examples--group-001.md#canonical-1122320213101203-0303313132002303-0020020303122132-1333101010210131-0301300020312010-1010132221110100-1310011322213111-1313212033332003): valid configuration.

<a id="canonical-1122320213101203-0303313132002303-0020020303122132-1333101010210131-0301300020312010-1010132221110100-1310011322213111-1313212033332003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Examples](data-sources--global_log_receiver--examples--group-001.md#canonical-3023311232200010-0133300122022312-3102212013022001-1130103102302322-3213013111200203-3211331221323312-0213002213312012-2002030133101203)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_global_log_receiver/data-source.tf`; digest `sha256:9f73258c5164e08d5a0687330200706260aacebe3599fa365152a1976e19940b`.

```terraform
# GlobalLogReceiver Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing GlobalLogReceiver by name
data "xcsh_global_log_receiver" "example" {
  name      = "example-global-log-receiver"
  namespace = "staging"
}

output "global_log_receiver_id" {
  value = data.xcsh_global_log_receiver.example.id
}
```
