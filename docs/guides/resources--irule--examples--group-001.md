---
page_title: "xcsh_irule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_irule examples."
---

# xcsh_irule examples

<a id="canonical-0102133322030003-2000111231113322-2322313303303010-3032112202330001-1002023033312232-0231210120003301-2003333032030221-0002311200311231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_irule](../resources/irule.md#canonical-2011333112331031-0020123312211213-0023313313132301-1232121102103020-3220220031303222-0310302031122301-3220110033323311-1131002100310203)
- Examples

<a id="canonical-3031032220230202-3212311232011322-0331023120311032-1122102021033023-0001103200230232-0310120001210133-2233022313320300-1313211021020122"></a>

### Complete configurations for `xcsh_irule`

- [Resource](resources--irule--examples--group-001.md#canonical-0121033310322201-2113232211001202-2212002230230323-1023202203302203-0131202322320320-0022210011030332-2011130333023221-3113320313120331): valid configuration.

<a id="canonical-0121033310322201-2113232211001202-2212002230230323-1023202203302203-0131202322320320-0022210011030332-2011130333023221-3113320313120331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_irule](../resources/irule.md#canonical-2011333112331031-0020123312211213-0023313313132301-1232121102103020-3220220031303222-0310302031122301-3220110033323311-1131002100310203)
- [Examples](resources--irule--examples--group-001.md#canonical-0102133322030003-2000111231113322-2322313303303010-3032112202330001-1002023033312232-0231210120003301-2003333032030221-0002311200311231)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_irule/resource.tf`; digest `sha256:42cb09305e24144e92e6cea8fc9300e3823e0b03a09b749cb4b4a6120a2dca4b`.

```terraform
# Irule Resource Example
# Manages iRule in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Irule configuration
resource "xcsh_irule" "example" {
  name      = "example-irule"
  namespace = "staging"

  description_spec = "example-value"
  irule            = "example-value"
}
```
