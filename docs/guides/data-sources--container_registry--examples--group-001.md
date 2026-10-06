---
page_title: "xcsh_container_registry examples"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_container_registry examples."
---

# xcsh_container_registry examples

<a id="canonical-2300022233211321-3223202330210002-0331022230100230-1002132001033202-0230103302011023-3102232023121120-1301332301023312-2131310121102020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_container_registry](../data-sources/container_registry.md#canonical-0113331231101032-2331101103012211-0031031222323020-2332013123121012-2010031020013003-3001130201213030-1222020321220312-1223011211021133)
- Examples

<a id="canonical-2101031121221110-3023300033202012-3130312300031022-1021023121113321-0003003132330301-2022011030310310-2321132212103021-1333131010201113"></a>

### Complete configurations for `xcsh_container_registry`

- [Data source](data-sources--container_registry--examples--group-001.md#canonical-1113010313311302-1131103033021122-2130031232233213-2020311233321332-0000303202001013-3203113012013120-1101210133102301-1303022020330121): valid configuration.

<a id="canonical-1113010313311302-1131103033021122-2130031232233213-2020311233321332-0000303202001013-3203113012013120-1101210133102301-1303022020330121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_container_registry](../data-sources/container_registry.md#canonical-0113331231101032-2331101103012211-0031031222323020-2332013123121012-2010031020013003-3001130201213030-1222020321220312-1223011211021133)
- [Examples](data-sources--container_registry--examples--group-001.md#canonical-2300022233211321-3223202330210002-0331022230100230-1002132001033202-0230103302011023-3102232023121120-1301332301023312-2131310121102020)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_container_registry/data-source.tf`; digest `sha256:235790f0ed1e99e7cff01491265f3f6071b78251c3ba60e8925b5d01fe4a9772`.

```terraform
# ContainerRegistry Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ContainerRegistry by name
data "xcsh_container_registry" "example" {
  name      = "example-container-registry"
  namespace = "staging"
}

output "container_registry_id" {
  value = data.xcsh_container_registry.example.id
}
```
