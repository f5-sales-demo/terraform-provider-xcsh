---
page_title: "xcsh_alert_policy examples"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_alert_policy examples."
---

# xcsh_alert_policy examples

<a id="canonical-0203231323102311-3112102032120023-3022320002010022-1132233023231020-3202031330311210-1313302001121023-0323003222313312-3133221332223323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- Examples

<a id="canonical-2201302331222032-2321233211202031-3012303113231330-1012033312120303-2310121100302202-3102001131131021-1010130101310200-3333133231001221"></a>

### Complete configurations for `xcsh_alert_policy`

- [Resource](resources--alert_policy--examples--group-001.md#canonical-1222003231100012-1132113102332331-2310331010301201-1020303102110002-3222210113331200-3021303323321120-0222023310120020-1101032113332303): valid configuration.

<a id="canonical-1222003231100012-1132113102332331-2310331010301201-1020303102110002-3222210113331200-3021303323321120-0222023310120020-1101032113332303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Examples](resources--alert_policy--examples--group-001.md#canonical-0203231323102311-3112102032120023-3022320002010022-1132233023231020-3202031330311210-1313302001121023-0323003222313312-3133221332223323)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_alert_policy/resource.tf`; digest `sha256:02e1db745e00ae3c15ffd89e3e1addc3f550309288f544ced75c1c7b1645161e`.

```terraform
# AlertPolicy Resource Example
# Manages new Alert Policy Object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertPolicy configuration
resource "xcsh_alert_policy" "example" {
  name      = "example-alert-policy"
  namespace = "staging"
}
```
