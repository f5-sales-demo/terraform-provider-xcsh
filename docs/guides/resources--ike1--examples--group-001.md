---
page_title: "xcsh_ike1 examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike1 examples."
---

# xcsh_ike1 examples

<a id="canonical-2221002230320112-3330322031333012-1003220211230121-0013321303220131-0320313301233213-2101130002222202-3222220033130000-1222323203032201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)
- Examples

<a id="canonical-3121020223131303-0012023301233213-3123320002322213-3233031003330311-1212032123030020-0021331011131013-0202033131201020-1000203333210030"></a>

### Complete configurations for `xcsh_ike1`

- [Resource](resources--ike1--examples--group-001.md#canonical-0212202023121020-3323020003311120-1210331132100101-3102212110230210-3030213020330112-1112320212102311-2002202120100300-0012001023112221): valid configuration.

<a id="canonical-0212202023121020-3323020003311120-1210331132100101-3102212110230210-3030213020330112-1112320212102311-2002202120100300-0012001023112221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_ike1](../resources/ike1.md#canonical-2001111300231301-1222011331231002-3213222003233302-1201002031332313-2211200021100302-3302222110022322-0332202221312311-1313230022101130)
- [Examples](resources--ike1--examples--group-001.md#canonical-2221002230320112-3330322031333012-1003220211230121-0013321303220131-0320313301233213-2101130002222202-3222220033130000-1222323203032201)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_ike1/resource.tf`; digest `sha256:3f9c75e223b14e98fb8be4f0bccd9b3e08736eec2e962a472124a5092f03950e`.

```terraform
# Ike1 Resource Example
# Manages a Ike1 resource in F5 Distributed Cloud for ike phase1 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Ike1 configuration
resource "xcsh_ike1" "example" {
  name      = "example-ike1"
  namespace = "staging"
}
```
