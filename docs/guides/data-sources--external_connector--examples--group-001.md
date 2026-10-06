---
page_title: "xcsh_external_connector examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_external_connector examples."
---

# xcsh_external_connector examples

<a id="canonical-1123210132313011-1012021123322001-3100001220223320-0001132020312132-3132011113011320-0130333103100101-3033032111330123-0030221132300301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- Examples

<a id="canonical-3121331332011030-2212103222222011-2201312032002013-1320131312010122-3131111302211200-2023211313102123-3333312320013011-2201100200233100"></a>

### Complete configurations for `xcsh_external_connector`

- [Data source](data-sources--external_connector--examples--group-001.md#canonical-0213100032310001-2023013202121123-2101132021002000-1201312213320131-0203100102021030-0200133133320113-0102210020102300-1113312122112103): valid configuration.

<a id="canonical-0213100032310001-2023013202121123-2101132021002000-1201312213320131-0203100102021030-0200133133320113-0102210020102300-1113312122112103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Examples](data-sources--external_connector--examples--group-001.md#canonical-1123210132313011-1012021123322001-3100001220223320-0001132020312132-3132011113011320-0130333103100101-3033032111330123-0030221132300301)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_external_connector/data-source.tf`; digest `sha256:8660d99e20bf0369dc735e047075e9bae757c75b8856770f0f1f3584ea3dc187`.

```terraform
# ExternalConnector Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ExternalConnector by name
data "xcsh_external_connector" "example" {
  name      = "example-external-connector"
  namespace = "staging"
}

output "external_connector_id" {
  value = data.xcsh_external_connector.example.id
}
```
