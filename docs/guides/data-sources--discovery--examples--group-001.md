---
page_title: "xcsh_discovery examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery examples."
---

# xcsh_discovery examples

<a id="canonical-2233113310202003-3011212032322120-0321110101103330-1110220301021333-2113230300122223-0032111232321131-3031100223030032-3233311203131303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- Examples

<a id="canonical-2002011131121100-1113031010012320-2221133322102010-3101032332331012-0202301203310030-2112101320221111-2213010130310313-2011230123020213"></a>

### Complete configurations for `xcsh_discovery`

- [Data source](data-sources--discovery--examples--group-001.md#canonical-1121032231312020-2323313120122112-3033101332110102-2122233330232211-2330102211332322-3020033233123213-2121023211303312-0120030000011231): valid configuration.

<a id="canonical-1121032231312020-2323313120122112-3033101332110102-2122233330232211-2330102211332322-3020033233123213-2121023211303312-0120030000011231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Examples](data-sources--discovery--examples--group-001.md#canonical-2233113310202003-3011212032322120-0321110101103330-1110220301021333-2113230300122223-0032111232321131-3031100223030032-3233311203131303)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_discovery/data-source.tf`; digest `sha256:de2200d3e23389756392f1e95f9b7544c16ad111ecd7a837e1de4f6caed18861`.

```terraform
# Discovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Discovery by name
data "xcsh_discovery" "example" {
  name      = "example-discovery"
  namespace = "staging"
}

output "discovery_id" {
  value = data.xcsh_discovery.example.id
}
```
