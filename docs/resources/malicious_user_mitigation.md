---
page_title: "xcsh_malicious_user_mitigation landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_malicious_user_mitigation landing."
---

# xcsh_malicious_user_mitigation landing

<a id="canonical-51762279e9e56351e964d117a51a3a003255979ed9c5abdf478c23f3f58615db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79e4581b82ba12dc491cd9eae1d772ae34eaf5b27c752696ae261270766a3606"></a>

## xcsh_malicious_user_mitigation — xcsh_malicious_user_mitigation / 83611296b153 / 2

Breadcrumbs:

- xcsh_malicious_user_mitigation

Manages malicious\_user\_mitigation creates a new object in the storage backend for
metadata.namespace in F5 Distributed Cloud.

<a id="canonical-71c7bb53ffc5655e27a153b3c0191301d3c9045a2064b904568a945325bc502c"></a>

## Prerequisites — xcsh_malicious_user_mitigation / 83611296b153 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3527b2b454e859bcc6529bd91bda983cba345a5af81d34ee652e09e89585d87d"></a>

## Minimal configuration — xcsh_malicious_user_mitigation / 83611296b153 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# MaliciousUserMitigation Resource Example
# Manages malicious_user_mitigation creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic MaliciousUserMitigation configuration
resource "xcsh_malicious_user_mitigation" "example" {
  name      = "example-malicious-user-mitigation"
  namespace = "staging"
}
```

<a id="canonical-816180e19a368d5686905b85c3eff0ab0a681ad39c835553a8e88195913e7a6e"></a>

## Root configuration — xcsh_malicious_user_mitigation / 83611296b153 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-bb65cfaccbdc0071538dfc37766e6dc333aa72336ff501eeacb5c5654e9389d8"></a>

## Next pages — xcsh_malicious_user_mitigation / 83611296b153 / 6

- [Property reference](../guides/resources--malicious_user_mitigation--reference--group-001.md#canonical-54042ff2b05e427630a3f23dd4b8756d9fae52d749e9e4b2663ad76de59b89d1)
- [Examples](../guides/resources--malicious_user_mitigation--examples--group-001.md#canonical-1afd96599fbc6ecd90e916f9d11197ea42dff6c59e02f3f73802395c5e6d414b)
- [Import](../guides/resources--malicious_user_mitigation--lifecycle--group-001.md#canonical-21fed318093c3c1eb1dcd7993306fd0e3a121c917f35e5c2746620f3c7978e5a)
- [Timeouts](../guides/resources--malicious_user_mitigation--lifecycle--group-001.md#canonical-83d7c70839b183e837aa51c48c849f7fe48882caebe1e1514739c814aacd3ad5)
