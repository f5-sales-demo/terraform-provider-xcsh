---
page_title: "xcsh_bgp_asn_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_asn_set examples."
---

# xcsh_bgp_asn_set examples

<a id="canonical-19443d73f29eec1400dbb793db9a20a8b2687458ef1ceba1d5cacf5041ca414d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6accabc4aac12e0d099776442b2c418efbdfbd5620b9b9c10b48648f4d9beaf"></a>

## Examples — Examples / 7ef56b44dc2c / 2

Breadcrumbs:

- [xcsh_bgp_asn_set](../data-sources/bgp_asn_set.md#canonical-117a5d23afbac64f2130fcb3963e2ecf60d3582aab779009e560eb05521c7e3f)
- Examples

<a id="canonical-acfd685103939f4115bdae522032d5a9db9dd3b51f172f6c0f06561f2a75d4af"></a>

## Complete configurations — Examples / 7ef56b44dc2c / 3

- [Data source](data-sources--bgp_asn_set--examples--group-001.md#canonical-1f2be3bed252a0fef27df3294910107587086e131e14575dddb851766350b5ca): valid configuration.

<a id="canonical-a1cc60a89486095cade0bf6ba777c754cdfa6e2b2073abdd21c83a5a691d32c1"></a>

## Next pages — Examples / 7ef56b44dc2c / 4

- [Data source](data-sources--bgp_asn_set--examples--group-001.md#canonical-1f2be3bed252a0fef27df3294910107587086e131e14575dddb851766350b5ca)
- [xcsh_bgp_asn_set](../data-sources/bgp_asn_set.md#canonical-117a5d23afbac64f2130fcb3963e2ecf60d3582aab779009e560eb05521c7e3f)

<a id="canonical-1f2be3bed252a0fef27df3294910107587086e131e14575dddb851766350b5ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c90410828f7f8f67cdbee582c92c1dd15941212c4f4b0778d5d8379d32c134e"></a>

## Data source — Data source / f1035abf356b / 2

Breadcrumbs:

- [xcsh_bgp_asn_set](../data-sources/bgp_asn_set.md#canonical-117a5d23afbac64f2130fcb3963e2ecf60d3582aab779009e560eb05521c7e3f)
- [Examples](data-sources--bgp_asn_set--examples--group-001.md#canonical-19443d73f29eec1400dbb793db9a20a8b2687458ef1ceba1d5cacf5041ca414d)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bgp_asn_set/data-source.tf`; digest `sha256:101ddc7bff960e99df5f4c481b2d71baf9042bf80e41813366a6137bff577b85`.

```terraform
# BGPAsnSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGPAsnSet by name
data "xcsh_bgp_asn_set" "example" {
  name      = "example-bgp-asn-set"
  namespace = "staging"
}

output "bgp_asn_set_id" {
  value = data.xcsh_bgp_asn_set.example.id
}
```

<a id="canonical-bb7620ec55cda51db91a3969e6ee118b7328e3443e4108d7cd4a883b1440faa2"></a>

## Next pages — Data source / f1035abf356b / 3

- [Examples](data-sources--bgp_asn_set--examples--group-001.md#canonical-19443d73f29eec1400dbb793db9a20a8b2687458ef1ceba1d5cacf5041ca414d)
- [xcsh_bgp_asn_set](../data-sources/bgp_asn_set.md#canonical-117a5d23afbac64f2130fcb3963e2ecf60d3582aab779009e560eb05521c7e3f)
