---
page_title: "xcsh_app_api_group examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_api_group examples."
---

# xcsh_app_api_group examples

<a id="canonical-72d930a44e683dcfdf007e37a08cca2ce04cbe9821f13751dffb51b2b504ad90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-591eb55b96d6e340fc0cc254877bd9e0470a42a662a8d99792645445e67048ce"></a>

## Examples — Examples / b0ac6cbfaba3 / 2

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)
- Examples

<a id="canonical-59dba445679cb9813ae99cc8fa79e1ee61b2a6b0cc4358a2a0de5ce62274ff05"></a>

## Complete configurations — Examples / b0ac6cbfaba3 / 3

- [Resource](resources--app_api_group--examples--group-001.md#canonical-f180291d49da7229ad3b16abbe644b3e124bead481677aa6f365f36686b134d0): valid configuration.

<a id="canonical-fc65828a2b65a985d9c1ed32044edfba60892e0dc503227a0e0176354cd11b4a"></a>

## Next pages — Examples / b0ac6cbfaba3 / 4

- [Resource](resources--app_api_group--examples--group-001.md#canonical-f180291d49da7229ad3b16abbe644b3e124bead481677aa6f365f36686b134d0)
- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)

<a id="canonical-f180291d49da7229ad3b16abbe644b3e124bead481677aa6f365f36686b134d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c04234112d3c20e6dd6221d38f27915a28c0784b9f035cb6cd07edc17c639024"></a>

## Resource — Resource / 39e778d3b264 / 2

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)
- [Examples](resources--app_api_group--examples--group-001.md#canonical-72d930a44e683dcfdf007e37a08cca2ce04cbe9821f13751dffb51b2b504ad90)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_api_group/resource.tf`; digest `sha256:f4976402381de0085eaa3ddefc5b729e1a883d869e4792bc2f8f960a6f9b6252`.

```terraform
# AppAPIGroup Resource Example
# Manages app_api_group creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppAPIGroup configuration
resource "xcsh_app_api_group" "example" {
  name      = "example-app-api-group"
  namespace = "staging"
}
```

<a id="canonical-1f0ad2d011748c4d6e94425236594dea99944c3782c033fbdc917bbb007fba1e"></a>

## Next pages — Resource / 39e778d3b264 / 3

- [Examples](resources--app_api_group--examples--group-001.md#canonical-72d930a44e683dcfdf007e37a08cca2ce04cbe9821f13751dffb51b2b504ad90)
- [xcsh_app_api_group](../resources/app_api_group.md#canonical-25df71841b832fff4daf25b51b7b804695bf4735432215f3290cd3188a095887)
