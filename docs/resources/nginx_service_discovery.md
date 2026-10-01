---
page_title: "xcsh_nginx_service_discovery landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nginx_service_discovery landing."
---

# xcsh_nginx_service_discovery landing

<a id="canonical-9f101755538801ac6046faa20459834b8511ae6b45bee321e3800bd29419268b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bfd1b09898b44112dc0d9e02cc612b03fff15a5daaa8e743a0e1616206df594a"></a>

## xcsh_nginx_service_discovery — xcsh_nginx_service_discovery / 861d2d0c4d8a / 2

Breadcrumbs:

- xcsh_nginx_service_discovery

Manages a Nginx Service Discovery resource in F5 Distributed Cloud for api to create nginx service
discovery object for a site or virtual site in system namespace. configuration.

<a id="canonical-7c88db2ad797c92264b70c3e123bcc8be6061b2119a64945d4154ab8f69af689"></a>

## Prerequisites — xcsh_nginx_service_discovery / 861d2d0c4d8a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-bef62abdf78b48630a769675f7915be9eee4db2c277afcc3d60ff30384ec57f9"></a>

## Minimal configuration — xcsh_nginx_service_discovery / 861d2d0c4d8a / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NginxServiceDiscovery Resource Example
# Manages a Nginx Service Discovery resource in F5 Distributed Cloud for api to create nginx service discovery object for a site or virtual site in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NginxServiceDiscovery configuration
resource "xcsh_nginx_service_discovery" "example" {
  name      = "example-nginx-service-discovery"
  namespace = "staging"
}
```

<a id="canonical-c94a8a142e1a68fa423e7c74fea52dde3281ec0d2e7c2e549ee3be1bfaefd0c5"></a>

## Root configuration — xcsh_nginx_service_discovery / 861d2d0c4d8a / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-c035f4cb648334c9bdb4adcc9f217487ef8db4ffb7555623240692f2b34cbae8"></a>

## Next pages — xcsh_nginx_service_discovery / 861d2d0c4d8a / 6

- [Property reference](../guides/resources--nginx_service_discovery--reference--group-001.md#canonical-7c43a05457b4d262da288872a5e7442922d9f0b5651f9b7c0efeb8a3e37d7c09)
- [Examples](../guides/resources--nginx_service_discovery--examples--group-001.md#canonical-c18b13d10d0226163fc8cf1360174bc98d4ef06d5d4daa9e87bea575065b2b6a)
- [Import](../guides/resources--nginx_service_discovery--lifecycle--group-001.md#canonical-e59531dcf659d7250fd99d0d8a49733700fce4acb7f11c2728fb5b386e15180c)
- [Timeouts](../guides/resources--nginx_service_discovery--lifecycle--group-001.md#canonical-b0e567bcc337829b40f3bb58c9f1012f3a2948dce57b31cfbd14620ee8ab8dd7)
