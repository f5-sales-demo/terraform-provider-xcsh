---
page_title: "xcsh_data_type examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_type examples."
---

# xcsh_data_type examples

<a id="canonical-3312333023112312-2102203021221000-0300101230031301-2301320200100131-0023130210313113-0230300203013330-3023002011013221-3311031012001100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-0130333301110020-0132111313313303-3030102230323221-2030031322131321-0220013133320111-2333311320030320-2201101033200331-2212312313233130)
- Examples

<a id="canonical-3013311022132010-3123230313210212-3322100000310011-3123210110133300-0020302321101130-0311132223032202-2133312223112303-0303102022311301"></a>

### Complete configurations for `xcsh_data_type`

- [Resource](resources--data_type--examples--group-001.md#canonical-3123010221000010-1300110032110002-0020113210323203-3323003033130211-3223132211302022-1011103232001133-3101011012212000-2123200123233222): valid configuration.

<a id="canonical-3123010221000010-1300110032110002-0020113210323203-3323003033130211-3223132211302022-1011103232001133-3101011012212000-2123200123233222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-0130333301110020-0132111313313303-3030102230323221-2030031322131321-0220013133320111-2333311320030320-2201101033200331-2212312313233130)
- [Examples](resources--data_type--examples--group-001.md#canonical-3312333023112312-2102203021221000-0300101230031301-2301320200100131-0023130210313113-0230300203013330-3023002011013221-3311031012001100)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_data_type/resource.tf`; digest `sha256:292681f3b6c18ec2ffbe2be1a0e17beb9c57b44e1777befb42567a6c3a8f12ee`.

```terraform
# DataType Resource Example
# Manages data_type creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DataType configuration
resource "xcsh_data_type" "example" {
  name      = "example-data-type"
  namespace = "staging"
}
```
