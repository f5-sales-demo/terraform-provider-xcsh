---
page_title: "xcsh_network_customer_edge_egress examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_customer_edge_egress examples."
---

# xcsh_network_customer_edge_egress examples

<a id="canonical-1333101203321301-3201210223302323-2122211312232210-3013232221331311-2120221213012230-0300222231222110-1301202011101010-2223303330333200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_customer_edge_egress](../data-sources/network_customer_edge_egress.md#canonical-3220221103330020-3110221000233302-0211101212002003-1012001301133302-1131222331211321-2333221222011322-0330222303230333-2322220210010320)
- Examples

<a id="canonical-1212011100313212-3111203133020113-0023320033031121-1300101200320213-2030012213312022-1003120110130220-1031331310022131-3022133320312312"></a>

### Complete configurations for `xcsh_network_customer_edge_egress`

- [Data source](data-sources--network_customer_edge_egress--examples--group-001.md#canonical-3010201100132110-0030003012023201-2323030012013011-0022022200121230-2222110121212311-0121121102300132-2310013211133020-2201000201132211): valid configuration.

<a id="canonical-3010201100132110-0030003012023201-2323030012013011-0022022200121230-2222110121212311-0121121102300132-2310013211133020-2201000201132211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_network_customer_edge_egress](../data-sources/network_customer_edge_egress.md#canonical-3220221103330020-3110221000233302-0211101212002003-1012001301133302-1131222331211321-2333221222011322-0330222303230333-2322220210010320)
- [Examples](data-sources--network_customer_edge_egress--examples--group-001.md#canonical-1333101203321301-3201210223302323-2122211312232210-3013232221331311-2120221213012230-0300222231222110-1301202011101010-2223303330333200)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_customer_edge_egress/data-source.tf`; digest `sha256:1a958054768616b5beeebffa91db22fe699af5f80112367c4fcd9c1d753d7368`.

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

data "xcsh_network_customer_edge_egress" "secure_mesh_v2" {}

# Use the domains with an FQDN-aware control. The legacy CE branch is not
# included in this data source.
output "secure_mesh_v2_https_egress" {
  value = {
    direction              = "egress"
    protocol               = "tcp"
    port                   = 443
    registration_addresses = data.xcsh_network_customer_edge_egress.secure_mesh_v2.registration_addresses
    domains                = data.xcsh_network_customer_edge_egress.secure_mesh_v2.domains
  }
}
```
