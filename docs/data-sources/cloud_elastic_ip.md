---
page_title: "xcsh_cloud_elastic_ip landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_elastic_ip landing."
---

# xcsh_cloud_elastic_ip landing

<a id="canonical-2b2f30e89bfb277ce22595cc89aff6edb2de7c5d6f5d72c5d631bef5fa2ab292"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d5ca3ac65dca6952e10fed01b5d019996af39157084fbf53a76dd8c9d617df95"></a>

## xcsh_cloud_elastic_ip — xcsh_cloud_elastic_ip / 03abe8e0e00d / 2

Breadcrumbs:

- xcsh_cloud_elastic_ip

Manages Cloud Elastic IP creates Cloud Elastic IP object Object is attached to a site in F5
Distributed Cloud.

<a id="canonical-a4fa942f1ceed1082d66481a66d382c9298ce453221eab7b5a30897bc1d0b5f3"></a>

## Prerequisites — xcsh_cloud_elastic_ip / 03abe8e0e00d / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-4e09f931cce60539e469c22d7d0b6adb7c1b02e6ea1b19d69c174a867717b99c"></a>

## Minimal configuration — xcsh_cloud_elastic_ip / 03abe8e0e00d / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudElasticIP Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudElasticIP by name
data "xcsh_cloud_elastic_ip" "example" {
  name      = "example-cloud-elastic-ip"
  namespace = "staging"
}

output "cloud_elastic_ip_id" {
  value = data.xcsh_cloud_elastic_ip.example.id
}
```

<a id="canonical-6e021f7285826b78c46e5f60a61c3ad20aac7d536e770c980098a86cb2c4b99d"></a>

## Root configuration — xcsh_cloud_elastic_ip / 03abe8e0e00d / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-4933436a8e191e1bce8c827d61c179c654ff14ddcec826737b0f4d116106477f"></a>

## Next pages — xcsh_cloud_elastic_ip / 03abe8e0e00d / 6

- [Property reference](../guides/data-sources--cloud_elastic_ip--reference--group-001.md#canonical-ef15d26eed8ba06029d76ff217be5d95cee1296437e2bcdfc5add660013e8dcb)
- [Examples](../guides/data-sources--cloud_elastic_ip--examples--group-001.md#canonical-076152532c256f4736b20a7aad09ccaa23aa6c26028f74c80e32ace794cb1a1e)
