---
page_title: "xcsh_forwarding_class examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_forwarding_class examples."
---

# xcsh_forwarding_class examples

<a id="canonical-1121111133103312-1312010112003212-2002312232032321-0022321033300030-1300301001320300-2220313031122222-2232333122231221-0111320013023222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212)
- Examples

<a id="canonical-2002130330310123-2132333011130310-3332223211223330-3202312113303313-0030003112321213-3033113320010211-2321300310333100-0320023310211331"></a>

### Complete configurations for `xcsh_forwarding_class`

- [Resource](resources--forwarding_class--examples--group-001.md#canonical-3033303113103310-0131012220003123-2322112013321232-0102023201210323-0131131012012333-0020000210022302-2023212223111123-2210111103001110): valid configuration.

<a id="canonical-3033303113103310-0131012220003123-2322112013321232-0102023201210323-0131131012012333-0020000210022302-2023212223111123-2210111103001110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md#canonical-2311030223121110-0200300102200301-0213002201300200-1210221011312013-2203303003202231-0033212100202321-0010322230031303-1000220313130212)
- [Examples](resources--forwarding_class--examples--group-001.md#canonical-1121111133103312-1312010112003212-2002312232032321-0022321033300030-1300301001320300-2220313031122222-2232333122231221-0111320013023222)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_forwarding_class/resource.tf`; digest `sha256:eb45941e172f8b977e57cad7117bbdf4795101f471776f657bf6e2faa09a4be8`.

```terraform
# ForwardingClass Resource Example
# Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ForwardingClass configuration
resource "xcsh_forwarding_class" "example" {
  name      = "example-forwarding-class"
  namespace = "staging"
}
```
