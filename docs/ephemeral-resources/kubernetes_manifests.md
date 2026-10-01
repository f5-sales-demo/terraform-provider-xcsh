---
page_title: "xcsh_kubernetes_manifests landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_kubernetes_manifests landing."
---

# xcsh_kubernetes_manifests landing

<a id="canonical-3020030201123201-1030130333132003-2002102001021330-1320220001031302-0001300203331022-2100320312131223-1211321010102033-2012323132111020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322300123203200-3003033311220331-0201030101303200-3000122131010301-3032112310310112-1220331321001211-1303111202313231-0203021012320310"></a>

## xcsh_kubernetes_manifests — xcsh_kubernetes_manifests / 230131313131 / 2

Breadcrumbs:

- xcsh_kubernetes_manifests

Kubernetes workload configuration.

<a id="canonical-3022021122313200-2300213210323013-2312200321022113-0210321100220120-2321230330000202-2311212310001213-3130123302021223-3312133120002033"></a>

## Prerequisites — xcsh_kubernetes_manifests / 230131313131 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2333112210221333-0120031130201020-1023310302111203-2221101212322022-2131112302223120-2113332031322200-3111221311010313-1302222311320101"></a>

## Minimal configuration — xcsh_kubernetes_manifests / 230131313131 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-2011230000021013-0103000301233031-3100001113230033-1001312131000200-2112121103332031-3130032213012333-3310110200220000-1032221211200001"></a>

## Root configuration — xcsh_kubernetes_manifests / 230131313131 / 5

Required root properties: `site`. Full root flags and choices appear in the property reference.

<a id="canonical-2333201000033111-0213103233000311-3303122222202002-1330120002030103-0110333303212132-2021000321013123-0211220002110230-2231113212300023"></a>

## Next pages — xcsh_kubernetes_manifests / 230131313131 / 6

- [Property reference](../guides/ephemeral-resources--kubernetes_manifests--reference--group-001.md#canonical-0222301022001323-1002221032030331-0321231112233121-0131212101033121-3023002220100301-3312321000102011-3211011101131320-3232112012103020)
- [Examples](../guides/ephemeral-resources--kubernetes_manifests--examples--group-001.md#canonical-3322303333023100-0233011223022113-0321002001333230-3022113113012233-1012213220103312-0221022122302132-3102202230320101-0230032011113332)
- [Lifecycle](../guides/ephemeral-resources--kubernetes_manifests--lifecycle--group-001.md#canonical-0131231131103223-3323102110011201-2331133102210002-1032333031030023-1132222323221120-0003011021302021-0020120030303000-2130113320101312)
