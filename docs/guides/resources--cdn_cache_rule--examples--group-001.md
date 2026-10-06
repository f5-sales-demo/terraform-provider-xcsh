---
page_title: "xcsh_cdn_cache_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_cache_rule examples."
---

# xcsh_cdn_cache_rule examples

<a id="canonical-0212331020322220-2311332101310031-2231102121022101-3011012031230323-2223331103003223-1202213101232312-0332331023100121-0133133331332211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- Examples

<a id="canonical-1201100113013032-1323210112021011-2301201030210330-3102200120022013-2331331312112103-2220100312333202-1011000130112133-1111013012023011"></a>

### Complete configurations for `xcsh_cdn_cache_rule`

- [Resource](resources--cdn_cache_rule--examples--group-001.md#canonical-0212021203112122-1321202003113303-1102113122112103-1201233203003222-2123203031122302-0022223302321100-3001231230330001-1311112111231021): valid configuration.

<a id="canonical-0212021203112122-1321202003113303-1102113122112103-1201233203003222-2123203031122302-0022223302321100-3001231230330001-1311112111231021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_cdn_cache_rule](../resources/cdn_cache_rule.md#canonical-3200322102222221-2020010330311323-2101331321122233-0110000212110022-0103013000123233-1320233321200001-0102123303130130-3123130022002001)
- [Examples](resources--cdn_cache_rule--examples--group-001.md#canonical-0212331020322220-2311332101310031-2231102121022101-3011012031230323-2223331103003223-1202213101232312-0332331023100121-0133133331332211)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cdn_cache_rule/resource.tf`; digest `sha256:a602b3888b8d6105a3cf057267d0d93b04e1fb4a8ec89a85c41a5bb720a38a3e`.

```terraform
# CDNCacheRule Resource Example
# Manages a CDN Cache Rule resource in F5 Distributed Cloud for cdn loadbalancer specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNCacheRule configuration
resource "xcsh_cdn_cache_rule" "example" {
  name      = "example-cdn-cache-rule"
  namespace = "staging"
}
```
