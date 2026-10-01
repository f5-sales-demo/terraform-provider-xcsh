---
page_title: "xcsh_malicious_user_mitigation landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_malicious_user_mitigation landing."
---

# xcsh_malicious_user_mitigation landing

<a id="canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1d291aedd70762d8629e8d40a1b4442fad168ec9c2a9836ec5d314a91f22abf"></a>

## xcsh_malicious_user_mitigation — xcsh_malicious_user_mitigation / cef7e1613003 / 2

Breadcrumbs:

- xcsh_malicious_user_mitigation

Manages malicious\_user\_mitigation creates a new object in the storage backend for
metadata.namespace in F5 Distributed Cloud.

<a id="canonical-44d1d1109e8c8b1044d1f978d90f9e6dc72461f7fa5b4972a52d41e74644d060"></a>

## Prerequisites — xcsh_malicious_user_mitigation / cef7e1613003 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-e4f4accfcf96fc07394615f4140739ec16716a4b761e86dab5dc4f68d0a6a3fe"></a>

## Minimal configuration — xcsh_malicious_user_mitigation / cef7e1613003 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# MaliciousUserMitigation Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing MaliciousUserMitigation by name
data "xcsh_malicious_user_mitigation" "example" {
  name      = "example-malicious-user-mitigation"
  namespace = "staging"
}

output "malicious_user_mitigation_id" {
  value = data.xcsh_malicious_user_mitigation.example.id
}
```

<a id="canonical-3b5fbd6c1226f3e232fbd79eb1978f65c745cb36e1826717bffc5a926f4d0a9c"></a>

## Root configuration — xcsh_malicious_user_mitigation / cef7e1613003 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-683c0cbee60e6eeccc54c130de1d0b478ee2ffa5d9f787805b0550632456cd42"></a>

## Next pages — xcsh_malicious_user_mitigation / cef7e1613003 / 6

- [Property reference](../guides/data-sources--malicious_user_mitigation--reference--group-001.md#canonical-e9b58fb6b75336611879fe7e9edb5bd112895f844e6a7f6e57cc31e81a6453ac)
- [Examples](../guides/data-sources--malicious_user_mitigation--examples--group-001.md#canonical-bc550721c1f7b8962ff6e6e12da3d5ce2a5e3dd5e0e1105f9540738bee71be27)
