---
page_title: "xcsh_tunnel examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tunnel examples."
---

# xcsh_tunnel examples

<a id="canonical-3300330231121311-2222032231103133-0121230003001020-2233233231213002-1010021013211133-0322111310201012-0020010233021132-1133232313203310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- Examples

<a id="canonical-3211311221320132-1120030333300222-2110303110133012-0101323000330101-0333110102230303-1133232320213022-2123030310133233-1312111032202231"></a>

### Complete configurations for `xcsh_tunnel`

- [Data source](data-sources--tunnel--examples--group-001.md#canonical-2310011211200203-2221333001022301-1121322111133330-0000313213120301-0020230122311012-3011230220120212-3323130220033033-1320130313311222): valid configuration.

<a id="canonical-2310011211200203-2221333001022301-1121322111133330-0000313213120301-0020230122311012-3011230220120212-3323130220033033-1320130313311222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_tunnel](../data-sources/tunnel.md#canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001)
- [Examples](data-sources--tunnel--examples--group-001.md#canonical-3300330231121311-2222032231103133-0121230003001020-2233233231213002-1010021013211133-0322111310201012-0020010233021132-1133232313203310)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_tunnel/data-source.tf`; digest `sha256:6bb09ba6599023728c287794941b0e304024407337d302d5ca4e83c7bdd1f81d`.

```terraform
# Tunnel Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Tunnel by name
data "xcsh_tunnel" "example" {
  name      = "example-tunnel"
  namespace = "staging"
}

output "tunnel_id" {
  value = data.xcsh_tunnel.example.id
}
```
