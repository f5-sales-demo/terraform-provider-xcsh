---
page_title: "xcsh_nat_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nat_policy examples."
---

# xcsh_nat_policy examples

<a id="canonical-2133002012221220-1223212000131030-1323323013232213-1331030310331231-3031032213110301-1030200010103130-1213033032023003-3130112123011213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- Examples

<a id="canonical-1021323320000020-1022133210010231-0311023313220020-2010102322222212-2213220222023322-2233021232232010-2132103021121133-1001231303333332"></a>

### Complete configurations for `xcsh_nat_policy`

- [Resource](resources--nat_policy--examples--group-001.md#canonical-2302320100223323-1103033101202233-2011332102122012-2321313122020033-1303211313033103-1112322132223331-1321113120223303-0002311000301120): valid configuration.

<a id="canonical-2302320100223323-1103033101202233-2011332102122012-2321313122020033-1303211313033103-1112322132223331-1321113120223303-0002311000301120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Examples](resources--nat_policy--examples--group-001.md#canonical-2133002012221220-1223212000131030-1323323013232213-1331030310331231-3031032213110301-1030200010103130-1213033032023003-3130112123011213)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_nat_policy/resource.tf`; digest `sha256:9976ae5d67503e6c32870094ecdf09b471ea80cfbc484860002722af35c9de08`.

```terraform
# NATPolicy Resource Example
# Manages a NAT Policy resource in F5 Distributed Cloud for nat policy create specification configures nat policy with multiple rules,.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NATPolicy configuration
resource "xcsh_nat_policy" "example" {
  name      = "example-nat-policy"
  namespace = "staging"
}
```
