---
page_title: "xcsh_cloud_elastic_ip landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_elastic_ip landing."
---

# xcsh_cloud_elastic_ip landing

<a id="canonical-46bf726603a8f595739fbcf7482329f46550f2379ad9a0cdb88b12f1d28bb69c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea14737ad5e4f3978bbe99c2d42d8c2f69586a2260ed6f8333298285f8da2871"></a>

## xcsh_cloud_elastic_ip — xcsh_cloud_elastic_ip / 40409020753f / 2

Breadcrumbs:

- xcsh_cloud_elastic_ip

Manages Cloud Elastic IP creates Cloud Elastic IP object Object is attached to a site in F5
Distributed Cloud.

<a id="canonical-8d80eb51fdaab23f1ed24b96f85b1f7cfc212c0639819f6ca0c8f08531dbcba9"></a>

## Prerequisites — xcsh_cloud_elastic_ip / 40409020753f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-d09df7ca945134eefff6d4737d13456f51e62c2cb1228884a28d6fb65bf5f3f9"></a>

## Minimal configuration — xcsh_cloud_elastic_ip / 40409020753f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudElasticIP Resource Example
# Manages Cloud Elastic IP creates Cloud Elastic IP object Object is attached to a site in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudElasticIP configuration
resource "xcsh_cloud_elastic_ip" "example" {
  name      = "example-cloud-elastic-ip"
  namespace = "staging"

  item_count = 1
}
```

<a id="canonical-916dc07d31f6914d7a1c651730f3e92b9bc612989157d5318b8f39731cf4c7e5"></a>

## Root configuration — xcsh_cloud_elastic_ip / 40409020753f / 5

Required root properties: `item_count`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-c18799441a2f3d32c472f28b3b19eeeac422213007710d98536a322153f95bd0"></a>

## Next pages — xcsh_cloud_elastic_ip / 40409020753f / 6

- [Property reference](../guides/resources--cloud_elastic_ip--reference--group-001.md#canonical-c0f91f91cfb563bfe37a3cba69683a2e56a66d3b7c53c84011ee5b0bf7d437a4)
- [Examples](../guides/resources--cloud_elastic_ip--examples--group-001.md#canonical-1d4686fe67fa6921671b8abc615eb2e9b7327b95438af7b56d449703a9ec52fe)
- [Import](../guides/resources--cloud_elastic_ip--lifecycle--group-001.md#canonical-50ab79f9d297efb219d9209ba664d30cad244e930e36e238f4a6b4b5aae85d7f)
- [Timeouts](../guides/resources--cloud_elastic_ip--lifecycle--group-001.md#canonical-b78a14dd51ae1db0a17b59487786fc9797eefac087ac0958d3d3d8213f3219f1)
