---
page_title: "xcsh_tmm_session_metrics landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tmm_session_metrics landing."
---

# xcsh_tmm_session_metrics landing

<a id="canonical-01b1407ad5dbf5565ef4e414bd5969c62fd5daad0530979b89d7a1917fc74573"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-371816fb0dad6c580987b571a4a886f8fd45987ee158c66748fc4d1da92ae45e"></a>

## xcsh_tmm_session_metrics — xcsh_tmm_session_metrics / e61e5a9fda8a / 2

Breadcrumbs:

- xcsh_tmm_session_metrics

Resource creation operation.

<a id="canonical-4b6829f6dcf015c826425e2e1d613e861d7c222d38ab0860492a814e6ec28829"></a>

## Prerequisites — xcsh_tmm_session_metrics / e61e5a9fda8a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-cd93d95708a087acf0db495ac1d5c6265611e2f95837a137d0a1463ff4c00847"></a>

## Minimal configuration — xcsh_tmm_session_metrics / e61e5a9fda8a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# TmmSessionMetrics DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_tmm_session_metrics" "example" {
  namespace = "example-value"
}

output "tmm_session_metrics_result" {
  value = data.xcsh_tmm_session_metrics.example
}
```

<a id="canonical-7c7a93a410378b51f9999243faccb88e378ac0587389349b6316d137ccaaee6d"></a>

## Root configuration — xcsh_tmm_session_metrics / e61e5a9fda8a / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-a84a044ce6849588e40ef5f607b821db0dd635a1ef8354866d8f1b8b2ae8d950"></a>

## Next pages — xcsh_tmm_session_metrics / e61e5a9fda8a / 6

- [Property reference](../guides/data-sources--tmm_session_metrics--reference--group-001.md#canonical-527d0176eb1fcf0d1ab47bc9f224535c861a92810150e2e9a25333b1b6c3fe84)
- [Examples](../guides/data-sources--tmm_session_metrics--examples--group-001.md#canonical-6f96b71270f7e9fc6cf197d842041cff18ad64580f208f0b4e07724b8ba98add)
