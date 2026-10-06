---
page_title: "xcsh_app_setting examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_setting examples."
---

# xcsh_app_setting examples

<a id="canonical-3320122113230000-3221120120232111-0331321100100320-1223112322232312-0212003023000333-3013212020121300-1022310012331022-1323330321021011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- Examples

<a id="canonical-2302131300113233-0133133033000323-0102113012213320-2010320232300301-1213122032022312-2023201111003132-1313333313101000-1303032303330102"></a>

### Complete configurations for `xcsh_app_setting`

- [Resource](resources--app_setting--examples--group-001.md#canonical-0300030103112322-2002201203210030-1013031211001003-0301213001221003-3202330000311300-2302322102003332-2322121113311303-3022201333120203): valid configuration.

<a id="canonical-0300030103112322-2002201203210030-1013031211001003-0301213001221003-3202330000311300-2302322102003332-2322121113311303-3022201333120203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md#canonical-0233122232111213-1333023112003320-0002130023110121-2101210112133031-2122112013202200-0320103020032232-0130220110001230-2003323312222110)
- [Examples](resources--app_setting--examples--group-001.md#canonical-3320122113230000-3221120120232111-0331321100100320-1223112322232312-0212003023000333-3013212020121300-1022310012331022-1323330321021011)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_setting/resource.tf`; digest `sha256:d31cbf807fc9dcf209556e05af346ee32ebdbf440fd52fffb2c11e50ae1d1f1a`.

```terraform
# AppSetting Resource Example
# Manages App setting configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppSetting configuration
resource "xcsh_app_setting" "example" {
  name      = "example-app-setting"
  namespace = "staging"
}
```
