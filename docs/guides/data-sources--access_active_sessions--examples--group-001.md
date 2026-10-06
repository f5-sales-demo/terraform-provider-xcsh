---
page_title: "xcsh_access_active_sessions examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_sessions examples."
---

# xcsh_access_active_sessions examples

<a id="canonical-2130122221103112-3022213203212002-0223130112330320-2112021120223032-1012230202021021-3012000133320332-2301032201010003-1302301020010310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_access_active_sessions](../data-sources/access_active_sessions.md#canonical-3120002103123020-1021131301111100-3330103332001301-2030202210211031-3012203032110100-2013230030123100-0331113020331110-1012323022312133)
- Examples

<a id="canonical-3212301031213003-3231322102122012-1301131313223020-1130213310101103-0230232032030321-0022031133030321-2211312030032310-2230101211302102"></a>

### Complete configurations for `xcsh_access_active_sessions`

- [Data source](data-sources--access_active_sessions--examples--group-001.md#canonical-1202122232020100-1122112103210323-1133013331211222-3120012223010033-0112221002020012-2230332303101231-1121101202231232-2213301123311001): valid configuration.

<a id="canonical-1202122232020100-1122112103210323-1133013331211222-3120012223010033-0112221002020012-2230332303101231-1121101202231232-2213301123311001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_access_active_sessions](../data-sources/access_active_sessions.md#canonical-3120002103123020-1021131301111100-3330103332001301-2030202210211031-3012203032110100-2013230030123100-0331113020331110-1012323022312133)
- [Examples](data-sources--access_active_sessions--examples--group-001.md#canonical-2130122221103112-3022213203212002-0223130112330320-2112021120223032-1012230202021021-3012000133320332-2301032201010003-1302301020010310)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_access_active_sessions/data-source.tf`; digest `sha256:9f2701c9ce37e0c0584f554dedc7733344fe85678c9dadc11b3997db86cac357`.

```terraform
# AccessActiveSessions DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_access_active_sessions" "example" {
  namespace = "example-value"
}

output "access_active_sessions_result" {
  value = data.xcsh_access_active_sessions.example
}
```
