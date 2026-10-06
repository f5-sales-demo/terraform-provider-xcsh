---
page_title: "xcsh_bgp examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp examples."
---

# xcsh_bgp examples

<a id="canonical-0011333200121323-1202013222010111-0303012103322321-1132212333033031-2001120120213230-0113023223130122-0310112310322311-2002111203001003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- Examples

<a id="canonical-3021012221333030-3221330032330301-1233132211000012-0001303332221131-3130022211132312-2121310122222301-3323032122320033-2230120200031313"></a>

### Complete configurations for `xcsh_bgp`

- [Data source](data-sources--bgp--examples--group-001.md#canonical-0321223223233121-0312330200233031-1130101030022030-2323111212200121-0203112130013111-3111001012023220-2322313032020103-0201102003110302): valid configuration.

<a id="canonical-0321223223233121-0312330200233031-1130101030022030-2323111212200121-0203112130013111-3111001012023220-2322313032020103-0201102003110302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_bgp](../data-sources/bgp.md#canonical-0330230310333303-2102033020231123-2303123101223012-3213333103231032-1203320012331320-3031311033301202-3020310230132101-0011223030112231)
- [Examples](data-sources--bgp--examples--group-001.md#canonical-0011333200121323-1202013222010111-0303012103322321-1132212333033031-2001120120213230-0113023223130122-0310112310322311-2002111203001003)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bgp/data-source.tf`; digest `sha256:0e624b6da5c866c4f991e89616b34b5f99775ffa87bb8832e89714aa1bc2809b`.

```terraform
# BGP Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGP by name
data "xcsh_bgp" "example" {
  name      = "example-bgp"
  namespace = "staging"
}

output "bgp_id" {
  value = data.xcsh_bgp.example.id
}
```
