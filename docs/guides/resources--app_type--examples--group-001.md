---
page_title: "xcsh_app_type examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_type examples."
---

# xcsh_app_type examples

<a id="canonical-3313232102212003-2110322003232032-1022002110331123-2100110311031302-2201131131312212-0233003322122301-1332112323100030-3033122312231231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-2220100013322301-3233230230220131-2230301203210321-0110213003232211-2331220333323130-0302001023303001-2320103333130202-1122210103122102)
- Examples

<a id="canonical-2021002301213113-3320312123301122-1013030002030320-3310330003010033-0020123133022211-2222320232333133-0012220020033310-3303202213010232"></a>

### Complete configurations for `xcsh_app_type`

- [Resource](resources--app_type--examples--group-001.md#canonical-1223322221000032-2221330233112230-1021112133211022-3011302012032303-3313202032110100-3133201300220323-0311011032013311-0331032330111201): valid configuration.

<a id="canonical-1223322221000032-2221330233112230-1021112133211022-3011302012032303-3313202032110100-3133201300220323-0311011032013311-0331032330111201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md#canonical-2220100013322301-3233230230220131-2230301203210321-0110213003232211-2331220333323130-0302001023303001-2320103333130202-1122210103122102)
- [Examples](resources--app_type--examples--group-001.md#canonical-3313232102212003-2110322003232032-1022002110331123-2100110311031302-2201131131312212-0233003322122301-1332112323100030-3033122312231231)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_type/resource.tf`; digest `sha256:9988ab36f101ce764d3f26103a86f6c75f3dcea74d272aeebe958208da89b71d`.

```terraform
# AppType Resource Example
# Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppType configuration
resource "xcsh_app_type" "example" {
  name      = "example-app-type"
  namespace = "staging"
}
```
