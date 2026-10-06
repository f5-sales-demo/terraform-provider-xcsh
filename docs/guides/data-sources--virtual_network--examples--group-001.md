---
page_title: "xcsh_virtual_network examples"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_virtual_network examples."
---

# xcsh_virtual_network examples

<a id="canonical-2222201120130020-1313132202212131-3101301212113322-0001333200121000-2022312231231133-2330203232100230-0333321232313201-2322110020333313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-2112023310021033-2012123022032122-0030112331222000-3101313021113310-0300312330333113-3110103330132201-2330310010223213-3013312303120122)
- Examples

<a id="canonical-0233310333122222-2122233131232010-1210110313222233-1132000303320310-0210332012022201-3302211230322332-2233320210003301-2220103223323111"></a>

### Complete configurations for `xcsh_virtual_network`

- [Data source](data-sources--virtual_network--examples--group-001.md#canonical-1302322113330001-3031103023210212-2332230121232333-0131301001300302-1101321011330003-0303030300130231-0323030122313222-2302102332031230): valid configuration.

<a id="canonical-1302322113330001-3031103023210212-2332230121232333-0131301001300302-1101321011330003-0303030300130231-0323030122313222-2302102332031230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-2112023310021033-2012123022032122-0030112331222000-3101313021113310-0300312330333113-3110103330132201-2330310010223213-3013312303120122)
- [Examples](data-sources--virtual_network--examples--group-001.md#canonical-2222201120130020-1313132202212131-3101301212113322-0001333200121000-2022312231231133-2330203232100230-0333321232313201-2322110020333313)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_virtual_network/data-source.tf`; digest `sha256:82f2d94c3d011a8d200fcd015e8a3406aea69e2a77a8f9d9eea9796b4b19e3bd`.

```terraform
# VirtualNetwork Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualNetwork by name
data "xcsh_virtual_network" "example" {
  name      = "example-virtual-network"
  namespace = "system"
}

output "virtual_network_id" {
  value = data.xcsh_virtual_network.example.id
}
```
