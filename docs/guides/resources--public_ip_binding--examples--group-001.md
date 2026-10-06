---
page_title: "xcsh_public_ip_binding examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_public_ip_binding examples."
---

# xcsh_public_ip_binding examples

<a id="canonical-3200320122123303-1130300212320023-3232122301302102-2221120230202213-0322231312213330-3311300112112302-3000101330021310-3012112203332131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_public_ip_binding](../resources/public_ip_binding.md#canonical-0203231102121100-0021222023322132-1033002220322133-0210023131031121-3331122113020013-3010312110213323-0102222031221012-2102213113003031)
- Examples

<a id="canonical-2301232120231121-3110203232030102-2320131230322032-2200013102321100-0233133212100020-3333223130002331-0213310030003011-3323230321101123"></a>

### Complete configurations for `xcsh_public_ip_binding`

- [Resource](resources--public_ip_binding--examples--group-001.md#canonical-2223123313212300-3023032303202012-2012211231031331-1311020231001102-3213120012231210-1013001311303121-0003302020113300-2213222303130231): valid configuration.

<a id="canonical-2223123313212300-3023032303202012-2012211231031331-1311020231001102-3213120012231210-1013001311303121-0003302020113300-2213222303130231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_public_ip_binding](../resources/public_ip_binding.md#canonical-0203231102121100-0021222023322132-1033002220322133-0210023131031121-3331122113020013-3010312110213323-0102222031221012-2102213113003031)
- [Examples](resources--public_ip_binding--examples--group-001.md#canonical-3200320122123303-1130300212320023-3232122301302102-2221120230202213-0322231312213330-3311300112112302-3000101330021310-3012112203332131)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_public_ip_binding/resource.tf`; digest `sha256:193afa89349b62b6fa18ffe2657f80b77590d2dd479611710090adfb88effdd9`.

```terraform
# PublicIPBinding Resource Example
# Manage the regional virtual-site binding of an already allocated public IP.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic PublicIPBinding configuration
resource "xcsh_public_ip_binding" "example" {
  name      = "example-public-ip-binding"
  namespace = "staging"

  expected_ip            = "example-value"
  virtual_site           = "example-value"
  virtual_site_namespace = "example-value"
}
```
