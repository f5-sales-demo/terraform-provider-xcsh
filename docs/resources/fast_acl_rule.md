---
page_title: "xcsh_fast_acl_rule landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl_rule landing."
---

# xcsh_fast_acl_rule landing

<a id="canonical-e591f185c7baa1942bcdc8fed5e2fd12321882858b9f3acc449ad83625ca98f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3874a85f3567f23361573f137f691064cf6dd22cdffde578b68effc79e3bb09"></a>

## xcsh_fast_acl_rule — xcsh_fast_acl_rule / f7eedb1710d4 / 2

Breadcrumbs:

- xcsh_fast_acl_rule

Manages new Fast ACL rule, has specification to match source IP, source port and action to apply in
F5 Distributed Cloud.

<a id="canonical-47283d708d2f07c7a9de59a3086b007b69ca222cbab7b59bffbd9ee115a4dc73"></a>

## Prerequisites — xcsh_fast_acl_rule / f7eedb1710d4 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-26c7f878fa6239e3a3953549f6a3b40fe4fd0646e6049b225b450f09aba8c24a"></a>

## Minimal configuration — xcsh_fast_acl_rule / f7eedb1710d4 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-c7c7599aed802c938e864ea1bea5c32c380ce850d2ca4f63b9176221687b1b9d"></a>

## Root configuration — xcsh_fast_acl_rule / f7eedb1710d4 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-ef2a359d93fcf6158441b06ecfd8cbb70fc893e54a27d1ce8413c7d4d444baeb"></a>

## Next pages — xcsh_fast_acl_rule / f7eedb1710d4 / 6

- [Property reference](../guides/resources--fast_acl_rule--reference--group-001.md#canonical-05ffc1846a674576d914dc565ae9b17a8d1f27478ca00fe1efbfeb29a301022e)
- [Examples](../guides/resources--fast_acl_rule--examples--group-001.md#canonical-becf5128bd0aaf0aeafece4ae0d486f841d7208724529312ec335349c213bbf5)
- [Import](../guides/resources--fast_acl_rule--lifecycle--group-001.md#canonical-d213b4f187cfcc430f199d77bc8ece0b74e8663db4c0688ec0aa0edeee2a5b20)
- [Timeouts](../guides/resources--fast_acl_rule--lifecycle--group-001.md#canonical-69888151c9c625ae2d34166c876385a7d05d27459eff5bb9436733059b50ebf7)
