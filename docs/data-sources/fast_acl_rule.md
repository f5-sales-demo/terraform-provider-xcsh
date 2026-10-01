---
page_title: "xcsh_fast_acl_rule landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl_rule landing."
---

# xcsh_fast_acl_rule landing

<a id="canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c34a9e9371889fc52c508357d7c2095bf98d37d84772e2eea1184c52c1fe6ae7"></a>

## xcsh_fast_acl_rule — xcsh_fast_acl_rule / 58139c18cade / 2

Breadcrumbs:

- xcsh_fast_acl_rule

Manages new Fast ACL rule, has specification to match source IP, source port and action to apply in
F5 Distributed Cloud.

<a id="canonical-ba11dd6a19a7691edee11521ddb9747d285d61bc7662a268d2ea27474262573a"></a>

## Prerequisites — xcsh_fast_acl_rule / 58139c18cade / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-e12b9f124d6791f8e4e357aa5efbec8cd9e75ba382606f91394fe3e0be569814"></a>

## Minimal configuration — xcsh_fast_acl_rule / 58139c18cade / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FastACLRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FastACLRule by name
data "xcsh_fast_acl_rule" "example" {
  name      = "example-fast-acl-rule"
  namespace = "staging"
}

output "fast_acl_rule_id" {
  value = data.xcsh_fast_acl_rule.example.id
}
```

<a id="canonical-265da474c4ae36411f08913e7f4fa3f30d039df6488dd9a37edce018d18a50a9"></a>

## Root configuration — xcsh_fast_acl_rule / 58139c18cade / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-afd2c571dc467b28797b005c807c70cc6219a5b8b01509c0b4f91383724ca5a0"></a>

## Next pages — xcsh_fast_acl_rule / 58139c18cade / 6

- [Property reference](../guides/data-sources--fast_acl_rule--reference--group-001.md#canonical-1c3c85be785c735afb4ad85cb212fa043eaa753afc4f7076378bdd310dc0b8dc)
- [Examples](../guides/data-sources--fast_acl_rule--examples--group-001.md#canonical-e5b9c4665cbf18fbdc5dded70558f8fea91828cc91f8581313bd641cf26d9b2c)
