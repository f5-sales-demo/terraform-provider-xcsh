---
page_title: "xcsh_access_active_session_terminate examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_access_active_session_terminate examples."
---

# xcsh_access_active_session_terminate examples

<a id="canonical-5f7f94defd00ef83888bc10e66540b97d13fdba27ff8c1198ed4890ac7763246"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dcabd67fe1a094d9842abf3ea46505baf70530b60d31ece1d06376435b5edce1"></a>

## Examples — Examples / 429b7029d220 / 2

Breadcrumbs:

- [xcsh_access_active_session_terminate](../actions/access_active_session_terminate.md#canonical-2f194631960b3886621a0ca55d7034950298c86bd87a8b01dbfa9bfbf6957dc8)
- Examples

<a id="canonical-4eaf9543f965e91ec34b45faf864fa99e08a79baea0111605a92ebc6c5642ee7"></a>

## Complete configurations — Examples / 429b7029d220 / 3

- [Action](actions--access_active_session_terminate--examples--group-001.md#canonical-94acdd7ff4c21bb1f1d286ada002ab11873606d86b322efc5aa314459ef8e2d4): valid configuration.

<a id="canonical-8a0bb939662a3ffd79968b7724bed3449dd6fd522cd73b07a6ae68c6e0c4ba1f"></a>

## Next pages — Examples / 429b7029d220 / 4

- [Action](actions--access_active_session_terminate--examples--group-001.md#canonical-94acdd7ff4c21bb1f1d286ada002ab11873606d86b322efc5aa314459ef8e2d4)
- [xcsh_access_active_session_terminate](../actions/access_active_session_terminate.md#canonical-2f194631960b3886621a0ca55d7034950298c86bd87a8b01dbfa9bfbf6957dc8)

<a id="canonical-94acdd7ff4c21bb1f1d286ada002ab11873606d86b322efc5aa314459ef8e2d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00180fcafffd1b73a233d00393734a90840d0f9993d5d53c40e1ace2ad57c748"></a>

## Action — Action / 912e1305b7fa / 2

Breadcrumbs:

- [xcsh_access_active_session_terminate](../actions/access_active_session_terminate.md#canonical-2f194631960b3886621a0ca55d7034950298c86bd87a8b01dbfa9bfbf6957dc8)
- [Examples](actions--access_active_session_terminate--examples--group-001.md#canonical-5f7f94defd00ef83888bc10e66540b97d13fdba27ff8c1198ed4890ac7763246)
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

<a id="canonical-2986507207484e5b38d4c714f9c2c1220c0472f02b45025561f36d8e5f017ec8"></a>

## Next pages — Action / 912e1305b7fa / 3

- [Examples](actions--access_active_session_terminate--examples--group-001.md#canonical-5f7f94defd00ef83888bc10e66540b97d13fdba27ff8c1198ed4890ac7763246)
- [xcsh_access_active_session_terminate](../actions/access_active_session_terminate.md#canonical-2f194631960b3886621a0ca55d7034950298c86bd87a8b01dbfa9bfbf6957dc8)
