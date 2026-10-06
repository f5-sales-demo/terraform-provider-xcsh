---
page_title: "xcsh_network_interface examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface examples."
---

# xcsh_network_interface examples

<a id="canonical-2133221333220133-3003232131233011-0323131130131010-0032031132202203-3203132322000010-3200121132023132-3333031110011010-0230013230113030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- Examples

<a id="canonical-2313020122021220-0203101220121221-1001110120332312-2132113223020133-0313130202321131-2330002011220000-1222121231222012-0211201223221302"></a>

### Complete configurations for `xcsh_network_interface`

- [Data source](data-sources--network_interface--examples--group-001.md#canonical-1030333002233230-1030302301300310-0330313123230111-0013312003023032-0230013211000220-3022102220031021-2013001023323221-0200202201121023): valid configuration.

<a id="canonical-1030333002233230-1030302301300310-0330313123230111-0013312003023032-0230013211000220-3022102220031021-2013001023323221-0200202201121023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132)
- [Examples](data-sources--network_interface--examples--group-001.md#canonical-2133221333220133-3003232131233011-0323131130131010-0032031132202203-3203132322000010-3200121132023132-3333031110011010-0230013230113030)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_interface/data-source.tf`; digest `sha256:2b23ed57ede50484e6ffd1c4ce6fdf9ba70d1bbaddf2b7e8ceb68afb488d14e2`.

```terraform
# NetworkInterface Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkInterface by name
data "xcsh_network_interface" "example" {
  name      = "example-network-interface"
  namespace = "staging"
}

output "network_interface_id" {
  value = data.xcsh_network_interface.example.id
}
```
