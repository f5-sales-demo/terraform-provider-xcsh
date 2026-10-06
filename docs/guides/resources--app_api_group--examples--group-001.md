---
page_title: "xcsh_app_api_group examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_api_group examples."
---

# xcsh_app_api_group examples

<a id="canonical-1302312103002210-1032122003313033-3133000013320313-2200203030220230-3200103023322120-0201330103131101-3133332311012302-2311001022312100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-0211313313012010-0123200302333333-1031223302112311-0123132320001012-2111233310130311-1003020201113303-0221003031030120-2022002111202013)
- Examples

<a id="canonical-1121013223111123-2112311232031000-3330003030021110-2013132331213200-1013002210022212-1202222031212113-2102121011101011-3212130010203032"></a>

### Complete configurations for `xcsh_app_api_group`

- [Resource](resources--app_api_group--examples--group-001.md#canonical-3301200002210131-1021312213020221-2231032301122223-2332121010230332-0102102332223110-2001121313222212-3303121133031212-2012230103103100): valid configuration.

<a id="canonical-3301200002210131-1021312213020221-2231032301122223-2332121010230332-0102102332223110-2001121313222212-3303121133031212-2012230103103100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_app_api_group](../resources/app_api_group.md#canonical-0211313313012010-0123200302333333-1031223302112311-0123132320001012-2111233310130311-1003020201113303-0221003031030120-2022002111202013)
- [Examples](resources--app_api_group--examples--group-001.md#canonical-1302312103002210-1032122003313033-3133000013320313-2200203030220230-3200103023322120-0201330103131101-3133332311012302-2311001022312100)
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
