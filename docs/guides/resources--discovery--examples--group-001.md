---
page_title: "xcsh_discovery examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery examples."
---

# xcsh_discovery examples

<a id="canonical-0302301231230113-2120313302220013-1021221011331111-2001300103031013-3110221312211320-0012310022011122-0031200230200333-0110131331330233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- Examples

<a id="canonical-1300103010110131-3213101220220101-1221010030011222-3100303112232333-1001313220303301-1213023111100023-1233022121111312-0310102302030010"></a>

### Complete configurations for `xcsh_discovery`

- [Resource](resources--discovery--examples--group-001.md#canonical-3231332330332133-3012022130101021-0231011222223303-3001221030232022-3032130300020103-1220010012220120-2033233200021330-3222223122312033): valid configuration.

<a id="canonical-3231332330332133-3012022130101021-0231011222223303-3001221030232022-3032130300020103-1220010012220120-2033233200021330-3222223122312033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md#canonical-3201131300321232-1303201100223230-3111121223110010-2301312203313103-1133001130100300-0132321013103013-3122130213231330-3003233233303202)
- [Examples](resources--discovery--examples--group-001.md#canonical-0302301231230113-2120313302220013-1021221011331111-2001300103031013-3110221312211320-0012310022011122-0031200230200333-0110131331330233)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_discovery/resource.tf`; digest `sha256:639a573707f4cc4152b42dbe4ad48b3edaa1017b066d38b4ee06ee6d80b4f03c`.

```terraform
# Discovery Resource Example
# Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site or virtual site in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Discovery configuration
resource "xcsh_discovery" "example" {
  name      = "example-discovery"
  namespace = "staging"
}
```
