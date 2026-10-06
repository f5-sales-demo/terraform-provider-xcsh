---
page_title: "xcsh_cminstance examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cminstance examples."
---

# xcsh_cminstance examples

<a id="canonical-1211223003023100-1212001021113111-0212321133311311-1111123102100211-1030030320013203-0001211003021032-2211300222200022-3223233112010302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-3030301003002223-1121312121232112-1232303010012000-0213330232003222-0030121133101123-1130200110322121-0013030301203132-0203031312323211)
- Examples

<a id="canonical-2003100001130013-2300312113130121-3323033312203333-3332100030213313-1210223212012031-2132222301212021-0113212213313230-3323123300211002"></a>

### Complete configurations for `xcsh_cminstance`

- [Resource](resources--cminstance--examples--group-001.md#canonical-0030121000332330-0011313010131230-2210212231210232-2200230213002310-0011030312023330-1113011221032011-2000100201032103-0020101333302020): valid configuration.

<a id="canonical-0030121000332330-0011313010131230-2210212231210232-2200230213002310-0011030312023330-1113011221032011-2000100201032103-0020101333302020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-3030301003002223-1121312121232112-1232303010012000-0213330232003222-0030121133101123-1130200110322121-0013030301203132-0203031312323211)
- [Examples](resources--cminstance--examples--group-001.md#canonical-1211223003023100-1212001021113111-0212321133311311-1111123102100211-1030030320013203-0001211003021032-2211300222200022-3223233112010302)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cminstance/resource.tf`; digest `sha256:8160520d3332724273c3557478a3147a7b31a2fd2927166c5bb6f2b977df1b95`.

```terraform
# Cminstance Resource Example
# Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Cminstance configuration
resource "xcsh_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "staging"

  port     = 1
  username = "example-value"
}
```
