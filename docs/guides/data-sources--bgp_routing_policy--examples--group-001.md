---
page_title: "xcsh_bgp_routing_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp_routing_policy examples."
---

# xcsh_bgp_routing_policy examples

<a id="canonical-0020013202220313-1110310122121231-1203301102111331-0112321221000103-2101220000303302-1202101110220103-3131301330003222-3113021201320212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- Examples

<a id="canonical-0213022013111110-2231012223231032-0323113030202320-3211323020320103-3123110222100131-2301331300011302-1303201120121122-2131331032133010"></a>

### Complete configurations for `xcsh_bgp_routing_policy`

- [Data source](data-sources--bgp_routing_policy--examples--group-001.md#canonical-1132013131121233-1203223013000330-1030133200231322-2301021203322013-0001012113023232-1332301210223102-3231101103003203-2222101010212112): valid configuration.

<a id="canonical-1132013131121233-1203223013000330-1030133200231322-2301021203322013-0001012113023232-1332301210223102-3231101103003203-2222101010212112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_bgp_routing_policy](../data-sources/bgp_routing_policy.md#canonical-3102200322103130-1203312020030213-3201133013110033-0210200032321012-3332210321102311-1131021302223032-2111203312332010-3101223110212011)
- [Examples](data-sources--bgp_routing_policy--examples--group-001.md#canonical-0020013202220313-1110310122121231-1203301102111331-0112321221000103-2101220000303302-1202101110220103-3131301330003222-3113021201320212)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bgp_routing_policy/data-source.tf`; digest `sha256:b3d6f35be703edce6abe0703f58285587174e0deae9414b75725a102a2aa63cb`.

```terraform
# BGPRoutingPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGPRoutingPolicy by name
data "xcsh_bgp_routing_policy" "example" {
  name      = "example-bgp-routing-policy"
  namespace = "staging"
}

output "bgp_routing_policy_id" {
  value = data.xcsh_bgp_routing_policy.example.id
}
```
