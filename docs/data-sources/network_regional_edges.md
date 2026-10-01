---
page_title: "xcsh_network_regional_edges landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_regional_edges landing."
---

# xcsh_network_regional_edges landing

<a id="canonical-4b07492253f8f10e61287f15adccadbfc866bfc8be3823bb27bb4e02b5d71089"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1edf1ee0faebe63598a2a904bf75c94a4fe3770d3dd323548e571e1d2de1ae1"></a>

## xcsh_network_regional_edges — xcsh_network_regional_edges / cb3f6efa55b0 / 2

Breadcrumbs:

- xcsh_network_regional_edges

Regional Edge IPv4 networks for origin ingress allowlists. Values are bundled from the pinned
OpenAPI release; this data source performs no network request. Ports and traffic direction are not
encoded in the manifest.

<a id="canonical-2470c00d83d24b1374ce8256fda2593b079cf1d6bb4235a87fda103109f0efa5"></a>

## Prerequisites — xcsh_network_regional_edges / cb3f6efa55b0 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-e33a8f7968d3c58ca269d1f642aaecd85bb1f0654d8eb3167ccde063b0713f47"></a>

## Minimal configuration — xcsh_network_regional_edges / cb3f6efa55b0 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-a495768c1432d374005b527bb11e03b49d1d85b06af962d3e310054e8c0b1eb2"></a>

## Root configuration — xcsh_network_regional_edges / cb3f6efa55b0 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-a0575e19c19107bde05fa6626bdc24b8ee2b6ca380be24b5168124b3155f6be2"></a>

## Next pages — xcsh_network_regional_edges / cb3f6efa55b0 / 6

- [Property reference](../guides/data-sources--network_regional_edges--reference--group-001.md#canonical-e079a6e575809cb77caea1bfb1f018d98eced62e8e5590c688ca25690a4e079a)
- [Examples](../guides/data-sources--network_regional_edges--examples--group-001.md#canonical-69cbc256fdb601aea7a5046bc5c1d7e5f6bd51f16032486f663d435c93baf01e)
