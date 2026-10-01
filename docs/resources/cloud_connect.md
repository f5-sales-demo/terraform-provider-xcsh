---
page_title: "xcsh_cloud_connect landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_connect landing."
---

# xcsh_cloud_connect landing

<a id="canonical-e5bdae3085908ed1207476726f7e81e21a7de63d8bc2bf47443b82363f99ac0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6d1d5fd149f772a5cebd077793da8a0624562457f9dd7829924be21aa320678"></a>

## xcsh_cloud_connect — xcsh_cloud_connect / 2e8cdc6b0afd / 2

Breadcrumbs:

- xcsh_cloud_connect

Manages a Cloud Connect resource in F5 Distributed Cloud for establishing connectivity to cloud
provider networks.

<a id="canonical-f9f8c5929c42300cf97eb7389bdde912921eec0ebf31b7a0931f5fb4110885d5"></a>

## Prerequisites — xcsh_cloud_connect / 2e8cdc6b0afd / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3b7cea73e4e37c0e289d88bee2ef092efff1397aee2e81394dea4a0309da31a0"></a>

## Minimal configuration — xcsh_cloud_connect / 2e8cdc6b0afd / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudConnect Resource Example
# Manages a Cloud Connect resource in F5 Distributed Cloud for establishing connectivity to cloud provider networks.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudConnect configuration
resource "xcsh_cloud_connect" "example" {
  name      = "example-cloud-connect"
  namespace = "staging"
}
```

<a id="canonical-54a120888aeecd0fe9fcbc2eae8fa863f887aa5d5cef2628e7af6212dba26ae0"></a>

## Root configuration — xcsh_cloud_connect / 2e8cdc6b0afd / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-e1357d6f5d084fdd4f1c65c23a18a2dbd9386ce69d79b9e2b023877db90f228b"></a>

## Next pages — xcsh_cloud_connect / 2e8cdc6b0afd / 6

- [Property reference](../guides/resources--cloud_connect--reference--group-001.md#canonical-c03db216fcd6eb05553183cb382fe87c6f5edf7c51ec269e7b462a17cb540a29)
- [Examples](../guides/resources--cloud_connect--examples--group-001.md#canonical-1b4515988a7c81c8322c907aa5b22dee9666424be27bf8fc5d830756c34a7f99)
- [Import](../guides/resources--cloud_connect--lifecycle--group-001.md#canonical-3351377a3d862db15edddd28f6830ef555bb45e442c25b470c963cb08928ba2a)
- [Timeouts](../guides/resources--cloud_connect--lifecycle--group-001.md#canonical-31291f3cd7669a0bcfbdb8c901b30a8638a3526b91aa222845d869decae00422)
