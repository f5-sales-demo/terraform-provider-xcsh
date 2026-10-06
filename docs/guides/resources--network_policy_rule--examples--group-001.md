---
page_title: "xcsh_network_policy_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_rule examples."
---

# xcsh_network_policy_rule examples

<a id="canonical-0332032102031030-3122310133212130-3301012133100130-0213200313233111-3110022130033331-2331310200120230-1230030121103030-3020102112122231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)
- Examples

<a id="canonical-0013100132212221-2030103313131131-1021112212212321-0130311333331103-3011203313032332-0323122103213013-3203133223210222-2010311121112201"></a>

### Complete configurations for `xcsh_network_policy_rule`

- [Resource](resources--network_policy_rule--examples--group-001.md#canonical-0230123322120131-3212132211311011-3332122203002302-1322201201212110-0110101331201331-1002301220103202-3221012201100012-0323020302322320): valid configuration.

<a id="canonical-0230123322120131-3212132211311011-3332122203002302-1322201201212110-0110101331201331-1002301220103202-3221012201100012-0323020302322320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_network_policy_rule](../resources/network_policy_rule.md#canonical-2112020201001103-2201031331322001-0223011022132003-2100033230200220-3103222330021332-3321002202302211-2131122200310202-1220303211023221)
- [Examples](resources--network_policy_rule--examples--group-001.md#canonical-0332032102031030-3122310133212130-3301012133100130-0213200313233111-3110022130033331-2331310200120230-1230030121103030-3020102112122231)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_policy_rule/resource.tf`; digest `sha256:0feaf66dfd5662ae652b573d32c61f50aa8eb63e0f1a54a1e0d029c6ad96ace9`.

```terraform
# NetworkPolicyRule Resource Example
# Manages network policy rule with configured parameters in specified namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkPolicyRule configuration
resource "xcsh_network_policy_rule" "example" {
  name      = "example-network-policy-rule"
  namespace = "staging"
}
```
