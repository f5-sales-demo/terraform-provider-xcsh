---
page_title: "xcsh_bigip_virtual_server examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_virtual_server examples."
---

# xcsh_bigip_virtual_server examples

<a id="canonical-2000033203001131-3130300110310212-0300112012011003-2122112212020211-2211220000232300-3001213332210023-3101102332301333-2331331103233121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bigip_virtual_server](../data-sources/bigip_virtual_server.md#canonical-1131322003302203-3133232130331303-3111132332133231-2012331201133203-0022310012333001-0003330233211211-0333312201021000-3230332012332020)
- Examples

<a id="canonical-3310303211120223-1330110212221122-1231100232102011-1133303132010021-3100301020202120-2300000222133212-2031121233132110-1331122020221012"></a>

### Complete configurations for `xcsh_bigip_virtual_server`

- [Data source](data-sources--bigip_virtual_server--examples--group-001.md#canonical-2301101033301122-1331212031100101-2131221013222020-1311301011213222-1011211133313023-0022122010200311-3222111102330012-0122213121111022): valid configuration.

<a id="canonical-2301101033301122-1331212031100101-2131221013222020-1311301011213222-1011211133313023-0022122010200311-3222111102330012-0122213121111022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_bigip_virtual_server](../data-sources/bigip_virtual_server.md#canonical-1131322003302203-3133232130331303-3111132332133231-2012331201133203-0022310012333001-0003330233211211-0333312201021000-3230332012332020)
- [Examples](data-sources--bigip_virtual_server--examples--group-001.md#canonical-2000033203001131-3130300110310212-0300112012011003-2122112212020211-2211220000232300-3001213332210023-3101102332301333-2331331103233121)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bigip_virtual_server/data-source.tf`; digest `sha256:4b98a932b2c1a0b5e2b9c608b25b6b3f74aea2df606bdee9675da50e166ce11d`.

```terraform
# BigIPVirtualServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BigIPVirtualServer by name
data "xcsh_bigip_virtual_server" "example" {
  name      = "example-bigip-virtual-server"
  namespace = "staging"
}

output "bigip_virtual_server_id" {
  value = data.xcsh_bigip_virtual_server.example.id
}
```
