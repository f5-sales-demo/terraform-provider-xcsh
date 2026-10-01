---
page_title: "xcsh_filter_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_filter_set examples."
---

# xcsh_filter_set examples

<a id="canonical-ea95c2ddcc082e916032a6c060bd5936a855ab6e5ac9587a6bc454afd9d21723"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7f0d17a44cf42e24a01b5a7b3cb157efb7f60719e3b3dc45439783e7f9d0e81"></a>

## Examples — Examples / 816b1af56619 / 2

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)
- Examples

<a id="canonical-efa3ffb3b8fcf5c924d231980edcad249bc80f06a13d2e88dc26b6418759c312"></a>

## Complete configurations — Examples / 816b1af56619 / 3

- [Resource](resources--filter_set--examples--group-001.md#canonical-b71a09f4b6101f63b835d811c4d39f0184c046d522d520d8c9d45b910659a5df): valid configuration.

<a id="canonical-60d43a40de50f2a74e86bee9a84bbe555ea19ba33381707794d7798afd585df3"></a>

## Next pages — Examples / 816b1af56619 / 4

- [Resource](resources--filter_set--examples--group-001.md#canonical-b71a09f4b6101f63b835d811c4d39f0184c046d522d520d8c9d45b910659a5df)
- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)

<a id="canonical-b71a09f4b6101f63b835d811c4d39f0184c046d522d520d8c9d45b910659a5df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-004951be4e6b5d7ca77a7ebc33185c2dcbc99371db64499c76b8d8408badeaaf"></a>

## Resource — Resource / a20b15d1fcf2 / 2

Breadcrumbs:

- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)
- [Examples](resources--filter_set--examples--group-001.md#canonical-ea95c2ddcc082e916032a6c060bd5936a855ab6e5ac9587a6bc454afd9d21723)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_filter_set/resource.tf`; digest `sha256:034749c02446cbe5e781e600fc1e24249a070ab78e205de7268da87fdb24c610`.

```terraform
# FilterSet Resource Example
# Manages specification in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic FilterSet configuration
resource "xcsh_filter_set" "example" {
  name      = "example-filter-set"
  namespace = "staging"

  context_key = "example-value"
}
```

<a id="canonical-b98cbb6380bec7b3fc87d315f63d4fc10d0ab854046f0a57fbd6402ff9c066ea"></a>

## Next pages — Resource / a20b15d1fcf2 / 3

- [Examples](resources--filter_set--examples--group-001.md#canonical-ea95c2ddcc082e916032a6c060bd5936a855ab6e5ac9587a6bc454afd9d21723)
- [xcsh_filter_set](../resources/filter_set.md#canonical-cfea2dffc9695306cb9a6f1b3049551ef60d6c696a1c33096488292434537441)
