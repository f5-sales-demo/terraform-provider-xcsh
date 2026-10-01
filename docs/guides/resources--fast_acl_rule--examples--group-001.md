---
page_title: "xcsh_fast_acl_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl_rule examples."
---

# xcsh_fast_acl_rule examples

<a id="canonical-becf5128bd0aaf0aeafece4ae0d486f841d7208724529312ec335349c213bbf5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee81ec8a69f9f0dba45fb5d50d060ffe4e805e342b8b3f24a039030531888b25"></a>

## Examples — Examples / c275b9ffd6c8 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)
- Examples

<a id="canonical-49a7569c4660fe60bf34989efa56bb975709bf8c96542d152f230c9698e8e0b8"></a>

## Complete configurations — Examples / c275b9ffd6c8 / 3

- [Resource](resources--fast_acl_rule--examples--group-001.md#canonical-f100434ac0c30a6a1e9aeeb9623e4a8bf3b4dea48ee1bdaeeb553b253eef43ce): valid configuration.

<a id="canonical-b426eeaea1388d2cd9d0beb9abd36836b45a6e25f4a6e1b133abb7089cdb00f3"></a>

## Next pages — Examples / c275b9ffd6c8 / 4

- [Resource](resources--fast_acl_rule--examples--group-001.md#canonical-f100434ac0c30a6a1e9aeeb9623e4a8bf3b4dea48ee1bdaeeb553b253eef43ce)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)

<a id="canonical-f100434ac0c30a6a1e9aeeb9623e4a8bf3b4dea48ee1bdaeeb553b253eef43ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96ba8667c6486220156fef1602a9fb5f2ddb8ed70df625c64ad7c99a48305a54"></a>

## Resource — Resource / 1d4b6d7840b3 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)
- [Examples](resources--fast_acl_rule--examples--group-001.md#canonical-becf5128bd0aaf0aeafece4ae0d486f841d7208724529312ec335349c213bbf5)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_fast_acl_rule/resource.tf`; digest `sha256:c95e30b59cae8b267c4570beb6043ccc798329a2b5e981ef4b126965c5dceb06`.

```terraform
# FastACLRule Resource Example
# Manages new Fast ACL rule, has specification to match source IP, source port and action to apply in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic FastACLRule configuration
resource "xcsh_fast_acl_rule" "example" {
  name      = "example-fast-acl-rule"
  namespace = "staging"
}
```

<a id="canonical-ca4f4e1a928e2e0e6cc96935004b2bc72cf6e3a2338f513fcefc41043ac73f2b"></a>

## Next pages — Resource / 1d4b6d7840b3 / 3

- [Examples](resources--fast_acl_rule--examples--group-001.md#canonical-becf5128bd0aaf0aeafece4ae0d486f841d7208724529312ec335349c213bbf5)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2)
