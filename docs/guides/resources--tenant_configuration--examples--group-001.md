---
page_title: "xcsh_tenant_configuration examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tenant_configuration examples."
---

# xcsh_tenant_configuration examples

<a id="canonical-f14acc1257a6c1413786a908104e7dc1d113d5ed4b44ad0345eaf33abc2cf8e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3608f0236d883e92702e7133183d045c8b3156e3e272ca87e65a4046c09fec3f"></a>

## Examples — Examples / 7c3c456e6971 / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)
- Examples

<a id="canonical-bcc7fd5c11fc61a3fe81c55e1d62e53d2631e07fccf224dd6da327b863ef35f1"></a>

## Complete configurations — Examples / 7c3c456e6971 / 3

- [Resource](resources--tenant_configuration--examples--group-001.md#canonical-1fd446d35755812ba3ace3c2d8ad10c0aca23cc624dfbb6cc68e3f89ca62cd13): valid configuration.

<a id="canonical-2f91b2bb5d2e2eee1e547a9f6ea55d0768eda88fcd5c44217636ccbbfb924bd1"></a>

## Next pages — Examples / 7c3c456e6971 / 4

- [Resource](resources--tenant_configuration--examples--group-001.md#canonical-1fd446d35755812ba3ace3c2d8ad10c0aca23cc624dfbb6cc68e3f89ca62cd13)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)

<a id="canonical-1fd446d35755812ba3ace3c2d8ad10c0aca23cc624dfbb6cc68e3f89ca62cd13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-898acade907c47deba7f49a60cb9c29ed8f6694fc60b1f8d6cfa093612c4ad12"></a>

## Resource — Resource / 24dc068b751e / 2

Breadcrumbs:

- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)
- [Examples](resources--tenant_configuration--examples--group-001.md#canonical-f14acc1257a6c1413786a908104e7dc1d113d5ed4b44ad0345eaf33abc2cf8e7)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_tenant_configuration/resource.tf`; digest `sha256:649c10c3cc2ff997128da4475c6c82fbfad1826ae6d86d2d3e3332f24b98496b`.

```terraform
# TenantConfiguration Resource Example
# Manages a Tenant Configuration resource in F5 Distributed Cloud for tenant configuration specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic TenantConfiguration configuration
resource "xcsh_tenant_configuration" "example" {
  name      = "example-tenant-configuration"
  namespace = "staging"
}
```

<a id="canonical-3ddb05d3d6a64979d9177dc340ef0deaa0934467abcb7104d0857610a97611a6"></a>

## Next pages — Resource / 24dc068b751e / 3

- [Examples](resources--tenant_configuration--examples--group-001.md#canonical-f14acc1257a6c1413786a908104e7dc1d113d5ed4b44ad0345eaf33abc2cf8e7)
- [xcsh_tenant_configuration](../resources/tenant_configuration.md#canonical-7e081faf644af0a986f55a7e090dca819919ecfa557bfc7618a12896eb23dc34)
