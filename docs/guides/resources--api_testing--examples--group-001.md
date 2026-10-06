---
page_title: "xcsh_api_testing examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_testing examples."
---

# xcsh_api_testing examples

<a id="canonical-3122233230003213-2011003033202121-2331112103320233-2321012312111213-1010101031002022-2203213121202220-1032111212002233-2230021233210122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- Examples

<a id="canonical-2300331111220122-0031003321212123-3311030231030223-1323330123023303-2232122100012200-0233133321322102-0331002333111202-1322022302002330"></a>

### Complete configurations for `xcsh_api_testing`

- [Resource](resources--api_testing--examples--group-001.md#canonical-3321221100122300-1131211313201230-1111200101001123-3300032223033202-2131230212231332-2031330321223113-3200002103032101-2132332220323233): valid configuration.

<a id="canonical-3321221100122300-1131211313201230-1111200101001123-3300032223033202-2131230212231332-2031330321223113-3200002103032101-2132332220323233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_api_testing](../resources/api_testing.md#canonical-1320313112233301-2001010323023332-0002303200021012-0212000001111221-2303001322120031-0333103011303201-3210321332031231-3311000110203202)
- [Examples](resources--api_testing--examples--group-001.md#canonical-3122233230003213-2011003033202121-2331112103320233-2321012312111213-1010101031002022-2203213121202220-1032111212002233-2230021233210122)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_api_testing/resource.tf`; digest `sha256:c1227be913c1052c8bc552ce9ebe6740f44bc74215b29a565808e69bd8cea2fb`.

```terraform
# APITesting Resource Example
# Manages a API Testing resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APITesting configuration
resource "xcsh_api_testing" "example" {
  name      = "example-api-testing"
  namespace = "staging"
}
```
