---
page_title: "xcsh_kubernetes_manifests examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_kubernetes_manifests examples."
---

# xcsh_kubernetes_manifests examples

<a id="canonical-3322303333023100-0233011223022113-0321002001333230-3022113113012233-1012213220103312-0221022122302132-3102202230320101-0230032011113332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_kubernetes_manifests](../ephemeral-resources/kubernetes_manifests.md#canonical-3020030201123201-1030130333132003-2002102001021330-1320220001031302-0001300203331022-2100320312131223-1211321010102033-2012323132111020)
- Examples

<a id="canonical-1001031103132220-1211231323121003-1000000213332223-3301202231031022-3101230001112333-1031122102110333-1011200003120130-2133102023012003"></a>

### Complete configurations for `xcsh_kubernetes_manifests`

- [Ephemeral](ephemeral-resources--kubernetes_manifests--examples--group-001.md#canonical-1222000303323313-0032220310022301-0211213320021023-0202310230003111-1122012310302201-0133201000110312-1123110111231202-3312023110332013): valid configuration.

<a id="canonical-1222000303323313-0032220310022301-0211213320021023-0202310230003111-1122012310302201-0133201000110312-1123110111231202-3312023110332013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Ephemeral example

Breadcrumbs:

- [xcsh_kubernetes_manifests](../ephemeral-resources/kubernetes_manifests.md#canonical-3020030201123201-1030130333132003-2002102001021330-1320220001031302-0001300203331022-2100320312131223-1211321010102033-2012323132111020)
- [Examples](ephemeral-resources--kubernetes_manifests--examples--group-001.md#canonical-3322303333023100-0233011223022113-0321002001333230-3022113113012233-1012213220103312-0221022122302132-3102202230320101-0230032011113332)
- Ephemeral

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/ephemeral-resources/xcsh_kubernetes_manifests/ephemeral.tf`; digest `sha256:e7611464c5404b77a4b6c720f8e02924f4419518332607f8e54efc9f31aad54e`.

```terraform
# KubernetesManifests EphemeralResource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

ephemeral "xcsh_kubernetes_manifests" "example" {
  site = "example-value"
}
```
