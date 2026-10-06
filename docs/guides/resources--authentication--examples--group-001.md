---
page_title: "xcsh_authentication examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authentication examples."
---

# xcsh_authentication examples

<a id="canonical-2222002030332030-2211020200323212-2310130011032302-0013021012021022-2211200022113033-1213312013012133-0022330111032222-3012301012201312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- Examples

<a id="canonical-2301020120000131-3000012100003203-1211000322131133-1231331031302210-1023302120310133-0313201330311030-0031120221332132-3320320132310200"></a>

### Complete configurations for `xcsh_authentication`

- [Resource](resources--authentication--examples--group-001.md#canonical-3102111333213010-3213112102323101-2302032301012330-3033323221301032-3101111301120310-2321221030132032-0200003312022213-0123200321031200): valid configuration.

<a id="canonical-3102111333213010-3213112102323101-2302032301012330-3033323221301032-3101111301120310-2321221030132032-0200003312022213-0123200321031200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210)
- [Examples](resources--authentication--examples--group-001.md#canonical-2222002030332030-2211020200323212-2310130011032302-0013021012021022-2211200022113033-1213312013012133-0022330111032222-3012301012201312)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_authentication/resource.tf`; digest `sha256:43532e046dd0e813a92a49d472a5eaa915d3ad0e96dcd6b49097f16328170876`.

```terraform
# Authentication Resource Example
# Manages a Authentication resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Authentication configuration
resource "xcsh_authentication" "example" {
  name      = "example-authentication"
  namespace = "staging"
}
```
