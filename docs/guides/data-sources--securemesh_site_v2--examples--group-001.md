---
page_title: "xcsh_securemesh_site_v2 examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 examples."
---

# xcsh_securemesh_site_v2 examples

<a id="canonical-3131012331030122-2111321213311303-0210231321111013-3130330201023003-3230320203231321-3301312023320210-0310032311033233-3332203200033313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- Examples

<a id="canonical-0330131221310220-1220021323222130-2223202010333003-3021123300333010-1011210221131222-3133312033101003-2012100120103201-1330223301302113"></a>

### Complete configurations for `xcsh_securemesh_site_v2`

- [Data source](data-sources--securemesh_site_v2--examples--group-001.md#canonical-2320121111203001-0201203100230213-1333311001213133-0333212022121321-0231200132112320-3310031302030330-3321100321022321-2222131112131300): valid configuration.

<a id="canonical-2320121111203001-0201203100230213-1333311001213133-0333212022121321-0231200132112320-3310031302030330-3321100321022321-2222131112131300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Examples](data-sources--securemesh_site_v2--examples--group-001.md#canonical-3131012331030122-2111321213311303-0210231321111013-3130330201023003-3230320203231321-3301312023320210-0310032311033233-3332203200033313)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_securemesh_site_v2/data-source.tf`; digest `sha256:2abaac1742f086ab988ef94e8d836fe6f538c549072b6b6dd846213e69e59239`.

```terraform
# SecuremeshSiteV2 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecuremeshSiteV2 by name
data "xcsh_securemesh_site_v2" "example" {
  name      = "example-securemesh-site-v2"
  namespace = "system"
}

output "securemesh_site_v2_id" {
  value = data.xcsh_securemesh_site_v2.example.id
}
```
