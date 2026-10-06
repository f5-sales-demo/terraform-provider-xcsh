---
page_title: "xcsh_fast_acl_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl_rule examples."
---

# xcsh_fast_acl_rule examples

<a id="canonical-2332303311010220-2331002222330022-3222333230321022-3200311020123320-1001311302002013-0210110221030102-3230030311031021-3002010323233311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- Examples

<a id="canonical-3232200132302022-1221332133003123-2210113323113111-0031001200333332-1032200011320310-0223202303330210-2200032100030011-0301202020230211"></a>

### Complete configurations for `xcsh_fast_acl_rule`

- [Resource](resources--fast_acl_rule--examples--group-001.md#canonical-3301000010031022-3000300300221222-0132212232322321-1202033210222023-3303231031322210-2032320123312232-3223111103230211-0332323310033032): valid configuration.

<a id="canonical-3301000010031022-3000300300221222-0132212232322321-1202033210222023-3303231031322210-2032320123312232-3223111103230211-0332323310033032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md#canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302)
- [Examples](resources--fast_acl_rule--examples--group-001.md#canonical-2332303311010220-2331002222330022-3222333230321022-3200311020123320-1001311302002013-0210110221030102-3230030311031021-3002010323233311)
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
