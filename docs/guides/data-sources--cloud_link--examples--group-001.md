---
page_title: "xcsh_cloud_link examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_link examples."
---

# xcsh_cloud_link examples

<a id="canonical-0132232132113111-3321222211103031-1212310102011310-1000030122323222-0011333223100213-3213303310122310-1013321333031321-0301211111333201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- Examples

<a id="canonical-0331333323330103-0223033003213303-3210023102302131-0021320030332333-0213101320020121-1211330220321233-1211000223321030-3010002133330332"></a>

### Complete configurations for `xcsh_cloud_link`

- [Data source](data-sources--cloud_link--examples--group-001.md#canonical-3000320003122322-1312212032232010-2110000333131332-2011230231131301-0213030222303311-0323312200322132-0310022312320102-1310332232200132): valid configuration.

<a id="canonical-3000320003122322-1312212032232010-2110000333131332-2011230231131301-0213030222303311-0323312200322132-0310022312320102-1310332232200132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Examples](data-sources--cloud_link--examples--group-001.md#canonical-0132232132113111-3321222211103031-1212310102011310-1000030122323222-0011333223100213-3213303310122310-1013321333031321-0301211111333201)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cloud_link/data-source.tf`; digest `sha256:a7eb053580e156c686886504937855a78a126659f9e2d4b1412c03b42464b698`.

```terraform
# CloudLink Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudLink by name
data "xcsh_cloud_link" "example" {
  name      = "example-cloud-link"
  namespace = "staging"
}

output "cloud_link_id" {
  value = data.xcsh_cloud_link.example.id
}
```
