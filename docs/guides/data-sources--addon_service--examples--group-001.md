---
page_title: "xcsh_addon_service examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_addon_service examples."
---

# xcsh_addon_service examples

<a id="canonical-2200032121332322-2232122133010113-2030311301331333-0012110031130233-2230232102303223-1021131330013322-3112031330020320-3332012230232033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_addon_service](../data-sources/addon_service.md#canonical-2330122101113210-0031213211000013-0113320231112112-2023300013333011-1030201010131303-2332021031023321-1030302210322030-1132200333322121)
- Examples

<a id="canonical-3132322203221200-2101110100222022-2121210111233120-0100323221312211-1031213313100220-0133220313333011-0203130301001212-3203032003012102"></a>

### Complete configurations for `xcsh_addon_service`

- [Data source](data-sources--addon_service--examples--group-001.md#canonical-2131011112100033-2103211030113023-0312103132202320-3213200310222301-1103122223130032-0301321223022231-0011221023322310-3333233023013213): valid configuration.

<a id="canonical-2131011112100033-2103211030113023-0312103132202320-3213200310222301-1103122223130032-0301321223022231-0011221023322310-3333233023013213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_addon_service](../data-sources/addon_service.md#canonical-2330122101113210-0031213211000013-0113320231112112-2023300013333011-1030201010131303-2332021031023321-1030302210322030-1132200333322121)
- [Examples](data-sources--addon_service--examples--group-001.md#canonical-2200032121332322-2232122133010113-2030311301331333-0012110031130233-2230232102303223-1021131330013322-3112031330020320-3332012230232033)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_addon_service/data-source.tf`; digest `sha256:809191858b0cd3ac6042ff132999dd4d2b0f0773a32efcde70b3a2eff7c8fa98`.

```terraform
# AddonService Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AddonService by name
data "xcsh_addon_service" "example" {
  name      = "example-addon-service"
  namespace = "staging"
}

output "addon_service_id" {
  value = data.xcsh_addon_service.example.id
}
```
