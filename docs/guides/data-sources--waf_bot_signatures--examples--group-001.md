---
page_title: "xcsh_waf_bot_signatures examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_bot_signatures examples."
---

# xcsh_waf_bot_signatures examples

<a id="canonical-29e9a910b02b7e528f124d51a1a50b78c81a18346aab6821cf1c7a6ae53af637"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3e377bc15315386958ea7cac8dc473c513b67a5d60e484c6aa785ad4e2f556d"></a>

## Examples — Examples / 99c3922930ab / 2

Breadcrumbs:

- [xcsh_waf_bot_signatures](../data-sources/waf_bot_signatures.md#canonical-6e1fb8fd8cfc0c30cdc9be34f2007f846063ac8689adc77b969c99d5d3cdb624)
- Examples

<a id="canonical-90a19baf8d77c0bcbddc8e69dc150fb47c4c75b707f7032ad81c081d9780a972"></a>

## Complete configurations — Examples / 99c3922930ab / 3

- [Data source](data-sources--waf_bot_signatures--examples--group-001.md#canonical-63d958ba7e7ff263687e7791c891de91969de3d7e5db57d087e310242a219f3a): valid configuration.

<a id="canonical-4ef22256314af9da5c401a749190e14289eb7e07448562d7916cf919b58a8026"></a>

## Next pages — Examples / 99c3922930ab / 4

- [Data source](data-sources--waf_bot_signatures--examples--group-001.md#canonical-63d958ba7e7ff263687e7791c891de91969de3d7e5db57d087e310242a219f3a)
- [xcsh_waf_bot_signatures](../data-sources/waf_bot_signatures.md#canonical-6e1fb8fd8cfc0c30cdc9be34f2007f846063ac8689adc77b969c99d5d3cdb624)

<a id="canonical-63d958ba7e7ff263687e7791c891de91969de3d7e5db57d087e310242a219f3a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed475412e5d99646dccbee709ffce7d1f6e701a33d8d21edeff38fdbcd7dacf4"></a>

## Data source — Data source / e7f46e7a4403 / 2

Breadcrumbs:

- [xcsh_waf_bot_signatures](../data-sources/waf_bot_signatures.md#canonical-6e1fb8fd8cfc0c30cdc9be34f2007f846063ac8689adc77b969c99d5d3cdb624)
- [Examples](data-sources--waf_bot_signatures--examples--group-001.md#canonical-29e9a910b02b7e528f124d51a1a50b78c81a18346aab6821cf1c7a6ae53af637)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_waf_bot_signatures/data-source.tf`; digest `sha256:0cea795e7e9cddaecc9f49e84c9c13efa8ac69b7e21b4ee8402220f9f8ec4b3f`.

```terraform
# WAFBotSignatures DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_bot_signatures" "example" {
}

output "waf_bot_signatures_result" {
  value = data.xcsh_waf_bot_signatures.example
}
```

<a id="canonical-0beafec7db1c9a67d3ea3ecadef21adee7274bfd61e55c0bb31cf46fd654ec8b"></a>

## Next pages — Data source / e7f46e7a4403 / 3

- [Examples](data-sources--waf_bot_signatures--examples--group-001.md#canonical-29e9a910b02b7e528f124d51a1a50b78c81a18346aab6821cf1c7a6ae53af637)
- [xcsh_waf_bot_signatures](../data-sources/waf_bot_signatures.md#canonical-6e1fb8fd8cfc0c30cdc9be34f2007f846063ac8689adc77b969c99d5d3cdb624)
