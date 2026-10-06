---
page_title: "xcsh_fast_acl examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl examples."
---

# xcsh_fast_acl examples

<a id="canonical-1233100313101303-1021322332001030-0313120000210313-3020023113032003-2211001033322132-2302120033032012-0300033322003332-3021212102003323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- Examples

<a id="canonical-0303233011233031-2303013323112213-1113203103312120-2322220112330030-2011322010333103-2203212000103100-0102100020230310-3110103133013213"></a>

### Complete configurations for `xcsh_fast_acl`

- [Data source](data-sources--fast_acl--examples--group-001.md#canonical-3033302120112100-1131213013222223-1301013212132100-1101102233301022-3120003130000031-1201303311111122-1002332203203321-2323321120031222): valid configuration.

<a id="canonical-3033302120112100-1131213013222223-1301013212132100-1101102233301022-3120003130000031-1201303311111122-1002332203203321-2323321120031222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Examples](data-sources--fast_acl--examples--group-001.md#canonical-1233100313101303-1021322332001030-0313120000210313-3020023113032003-2211001033322132-2302120033032012-0300033322003332-3021212102003323)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_fast_acl/data-source.tf`; digest `sha256:8fad5fdbb88c4dc30f2314eb3a4717e4d1278f83525eb02233ddf9ff899b963a`.

```terraform
# FastACL Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FastACL by name
data "xcsh_fast_acl" "example" {
  name      = "example-fast-acl"
  namespace = "system"
}

output "fast_acl_id" {
  value = data.xcsh_fast_acl.example.id
}
```
