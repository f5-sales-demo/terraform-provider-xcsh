---
page_title: "xcsh_access_active_sessions_terminate examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_sessions_terminate examples."
---

# xcsh_access_active_sessions_terminate examples

<a id="canonical-1233211303213321-3101030230133003-1313322330111221-1220211022000320-3122330010001302-0000122213122000-2213020333311111-1103212212202312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_access_active_sessions_terminate](../actions/access_active_sessions_terminate.md#canonical-2232333201313313-0211033213132122-1112020213103320-3212121022102212-1220232111200021-2230200120313012-2130013111233211-1310221022123111)
- Examples

<a id="canonical-1330113223202233-1223112001321030-1030220320320110-2302021133230100-3030032322210330-0213020010011120-2300323233213202-1302122312123313"></a>

### Complete configurations for `xcsh_access_active_sessions_terminate`

- [Action](actions--access_active_sessions_terminate--examples--group-001.md#canonical-0323022220300023-0020313113010222-2300010131030002-0211131012302230-2213331102112032-2023220223020302-0230331230022310-0020130211311311): valid configuration.

<a id="canonical-0323022220300023-0020313113010222-2300010131030002-0211131012302230-2213331102112032-2023220223020302-0230331230022310-0020130211311311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Action example

Breadcrumbs:

- [xcsh_access_active_sessions_terminate](../actions/access_active_sessions_terminate.md#canonical-2232333201313313-0211033213132122-1112020213103320-3212121022102212-1220232111200021-2230200120313012-2130013111233211-1310221022123111)
- [Examples](actions--access_active_sessions_terminate--examples--group-001.md#canonical-1233211303213321-3101030230133003-1313322330111221-1220211022000320-3122330010001302-0000122213122000-2213020333311111-1103212212202312)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_access_active_sessions_terminate/action.tf`; digest `sha256:e3e7c43f13f1ace8b437c1c8885aa667a9add78ae921c28c7a137377165d7c29`.

```terraform
# AccessActiveSessionsTerminate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_access_active_sessions_terminate" "example" {
  config {
    namespace = "example-value"
  }
}
```
