---
page_title: "xcsh_dc_cluster_group examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dc_cluster_group examples."
---

# xcsh_dc_cluster_group examples

<a id="canonical-2f4264098e0152fbf2779a6bbae0471ac64aac33e9f2333b107598f53731b774"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74e074a1b643021d10958abb189212586461e65e67aa5c15adc3962265a13069"></a>

## Examples — Examples / bc2bd82b909b / 2

Breadcrumbs:

- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0b65328cb220f5a226ff8b6ec19608451af2bb841451b04a1b840ddbffd72279)
- Examples

<a id="canonical-693877df50feabf7bbd56886d21d9dd0a99ec6ff3a92a70f51a0f3924698346c"></a>

## Complete configurations — Examples / bc2bd82b909b / 3

- [Resource](resources--dc_cluster_group--examples--group-001.md#canonical-2be6cd3ba8228c451b80f923a69a426558b1e76d110fe9c9d64cd46b4dfef5fb): valid configuration.

<a id="canonical-a11cfdda6c6ea9eb4eae8945ce9479ef36f444e6757156cd4503fda2d91e6a35"></a>

## Next pages — Examples / bc2bd82b909b / 4

- [Resource](resources--dc_cluster_group--examples--group-001.md#canonical-2be6cd3ba8228c451b80f923a69a426558b1e76d110fe9c9d64cd46b4dfef5fb)
- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0b65328cb220f5a226ff8b6ec19608451af2bb841451b04a1b840ddbffd72279)

<a id="canonical-2be6cd3ba8228c451b80f923a69a426558b1e76d110fe9c9d64cd46b4dfef5fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abca0dd9b1b55e0e14f5afe74be69741f24733532e07c9dc4d8c09085ba294de"></a>

## Resource — Resource / e6bd60c37ca1 / 2

Breadcrumbs:

- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0b65328cb220f5a226ff8b6ec19608451af2bb841451b04a1b840ddbffd72279)
- [Examples](resources--dc_cluster_group--examples--group-001.md#canonical-2f4264098e0152fbf2779a6bbae0471ac64aac33e9f2333b107598f53731b774)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dc_cluster_group/resource.tf`; digest `sha256:8a0876385980d3e39be1eb4ce88a5f57c0c3917ba68858a952fc5d4f99aaef53`.

```terraform
# DcClusterGroup Resource Example
# Manages DC Cluster group in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DcClusterGroup configuration
resource "xcsh_dc_cluster_group" "example" {
  name      = "example-dc-cluster-group"
  namespace = "system"
}
```

<a id="canonical-1b339e07448b80f7bacf0cb718b0f3f6ae3e54fd8c07e31617c1af447354512a"></a>

## Next pages — Resource / e6bd60c37ca1 / 3

- [Examples](resources--dc_cluster_group--examples--group-001.md#canonical-2f4264098e0152fbf2779a6bbae0471ac64aac33e9f2333b107598f53731b774)
- [xcsh_dc_cluster_group](../resources/dc_cluster_group.md#canonical-0b65328cb220f5a226ff8b6ec19608451af2bb841451b04a1b840ddbffd72279)
