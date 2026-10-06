---
page_title: "xcsh_access_active_session_terminate examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_session_terminate examples."
---

# xcsh_access_active_session_terminate examples

<a id="canonical-1133133321103132-3331000032332003-2020202330010032-1212111000232113-3101033331232202-1333332030010121-2032311020210022-3013131203021012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_access_active_session_terminate](../actions/access_active_session_terminate.md#canonical-0233012110120301-2112002303202012-1202012200302211-1131130003102111-0002212030201223-3120132220230001-3123332221233323-3312211113313020)
- Examples

<a id="canonical-3130222331121333-3201220021103121-2010022223330332-2210121100112322-3313001103002312-0031030132303201-3100120313121003-1123113231303201"></a>

### Complete configurations for `xcsh_access_active_session_terminate`

- [Action](actions--access_active_session_terminate--examples--group-001.md#canonical-2110223031311333-3310300201232301-3301310220122231-2200000222230101-2013031200123120-1223030202323330-1122220301101011-2132332032023110): valid configuration.

<a id="canonical-2110223031311333-3310300201232301-3301310220122231-2200000222230101-2013031200123120-1223030202323330-1122220301101011-2132332032023110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Action example

Breadcrumbs:

- [xcsh_access_active_session_terminate](../actions/access_active_session_terminate.md#canonical-0233012110120301-2112002303202012-1202012200302211-1131130003102111-0002212030201223-3120132220230001-3123332221233323-3312211113313020)
- [Examples](actions--access_active_session_terminate--examples--group-001.md#canonical-1133133321103132-3331000032332003-2020202330010032-1212111000232113-3101033331232202-1333332030010121-2032311020210022-3013131203021012)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_access_active_session_terminate/action.tf`; digest `sha256:1bbb807c7b4d5333250c2b20cf18f88631715519e818829cb7766f98ee751aa7`.

```terraform
# AccessActiveSessionTerminate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_access_active_session_terminate" "example" {
  config {
    id        = "example-value"
    namespace = "example-value"
  }
}
```
