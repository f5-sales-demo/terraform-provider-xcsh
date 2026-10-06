---
page_title: "xcsh_network_regional_edges examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_regional_edges examples."
---

# xcsh_network_regional_edges examples

<a id="canonical-1221302330021112-3331231200012232-2213221100101223-3011300131133211-3312233111013301-1200030210201233-1212033110031130-2103232233000132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_regional_edges](../data-sources/network_regional_edges.md#canonical-1023001310210202-1103332033010032-1201022013330111-2231303022312333-3020121223333020-2332032002032323-0213232310320002-2311311301002021)
- Examples

<a id="canonical-2111313112222121-0232122120112313-2133322101312222-1113002332221132-1301123133012300-1031110121202220-1030213210330212-2233202212000022"></a>

### Complete configurations for `xcsh_network_regional_edges`

- [Data source](data-sources--network_regional_edges--examples--group-001.md#canonical-3012131233211333-0113131211111320-3011322331211210-2200313333300101-3232010110232233-2021013200123312-1311122020120212-2110002123211023): valid configuration.

<a id="canonical-3012131233211333-0113131211111320-3011322331211210-2200313333300101-3232010110232233-2021013200123312-1311122020120212-2110002123211023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_network_regional_edges](../data-sources/network_regional_edges.md#canonical-1023001310210202-1103332033010032-1201022013330111-2231303022312333-3020121223333020-2332032002032323-0213232310320002-2311311301002021)
- [Examples](data-sources--network_regional_edges--examples--group-001.md#canonical-1221302330021112-3331231200012232-2213221100101223-3011300131133211-3312233111013301-1200030210201233-1212033110031130-2103232233000132)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_regional_edges/data-source.tf`; digest `sha256:1744edb321716a35476a9a25cade01cfd9671a46b90a9dfb2f90ccff555099d6`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

# Select the Regional Edge source networks that may initiate HTTPS connections
# to an origin. Omit regions to return all published regions.
data "xcsh_network_regional_edges" "origin_ingress" {
  regions = ["americas", "europe"]
}

output "https_origin_ingress" {
  value = {
    direction   = "ingress"
    protocol    = "tcp"
    port        = 443
    cidr_blocks = data.xcsh_network_regional_edges.origin_ingress.cidr_blocks
  }
}
```
