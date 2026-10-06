---
page_title: "xcsh_tunnel examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tunnel examples."
---

# xcsh_tunnel examples

<a id="canonical-2313003312031200-0221021012003310-1331120231133012-0101333030122330-2020203002311013-1011200200333123-1332322203132332-2103103231221131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- Examples

<a id="canonical-0012233103223313-2313011312312121-3202312000332120-1033332021011332-1011001202231323-0230220232130212-3121011221000020-2331112332011012"></a>

### Complete configurations for `xcsh_tunnel`

- [Resource](resources--tunnel--examples--group-001.md#canonical-3120322120112022-0023131222020213-3313233331303131-1021012313112101-0200300131221330-0222133001202000-3310231301110112-3323213303332023): valid configuration.

<a id="canonical-3120322120112022-0023131222020213-3313233331303131-1021012313112101-0200300131221330-0222133001202000-3310231301110112-3323213303332023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_tunnel](../resources/tunnel.md#canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122)
- [Examples](resources--tunnel--examples--group-001.md#canonical-2313003312031200-0221021012003310-1331120231133012-0101333030122330-2020203002311013-1011200200333123-1332322203132332-2103103231221131)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_tunnel/resource.tf`; digest `sha256:d8686dd99ea2784f4ce424a8e06432fe8502eabcb7549c1f57d3a841fff62174`.

```terraform
# Tunnel Resource Example
# Manages tunnel in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Tunnel configuration
resource "xcsh_tunnel" "example" {
  name      = "example-tunnel"
  namespace = "staging"
}
```
