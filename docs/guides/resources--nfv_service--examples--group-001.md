---
page_title: "xcsh_nfv_service examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service examples."
---

# xcsh_nfv_service examples

<a id="canonical-fb601c4a70a8553becbfe72e87df89c17d140f83c5c1549cbb177631c1d2b373"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57017d3930a0ccec23b7ccfdb544973eabdfee5f40f0891519082275512f221a"></a>

## Examples — Examples / ed497c549096 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- Examples

<a id="canonical-62e21b63189688a1d693ca1205ebd03487d7e6cfcaccdf912214cfd2526ebeb8"></a>

## Complete configurations — Examples / ed497c549096 / 3

- [Resource](resources--nfv_service--examples--group-001.md#canonical-67a5baa10cad0674a4c6aaccae0b898f0d7432a598c2c97f7f14ca1fdfa05ce7): valid configuration.

<a id="canonical-8c79abcb2ed66b456ef9139742c8e5dedbaf0685b8ccb4d07a5482415a42aba2"></a>

## Next pages — Examples / ed497c549096 / 4

- [Resource](resources--nfv_service--examples--group-001.md#canonical-67a5baa10cad0674a4c6aaccae0b898f0d7432a598c2c97f7f14ca1fdfa05ce7)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-67a5baa10cad0674a4c6aaccae0b898f0d7432a598c2c97f7f14ca1fdfa05ce7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc5a577c689e6abafc963628a5bd28049e2e41e560b586cf681cd6edd5b8e4a8"></a>

## Resource — Resource / bc59ecd94c89 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Examples](resources--nfv_service--examples--group-001.md#canonical-fb601c4a70a8553becbfe72e87df89c17d140f83c5c1549cbb177631c1d2b373)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_nfv_service/resource.tf`; digest `sha256:7b15b9a747e6ef39c623ab924043149ec709c96c07747307cc5c511433ad09f9`.

```terraform
# NfvService Resource Example
# Manages new NFV service with configured parameters in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NfvService configuration
resource "xcsh_nfv_service" "example" {
  name      = "example-nfv-service"
  namespace = "staging"
}
```

<a id="canonical-0c2bd50b869f5a7e3c4de7cffbaa42adc79acaf307cf8903dd7681536191aea1"></a>

## Next pages — Resource / bc59ecd94c89 / 3

- [Examples](resources--nfv_service--examples--group-001.md#canonical-fb601c4a70a8553becbfe72e87df89c17d140f83c5c1549cbb177631c1d2b373)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
