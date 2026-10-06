---
page_title: "xcsh_securemesh_site_v2 examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 examples."
---

# xcsh_securemesh_site_v2 examples

<a id="canonical-1303300012312320-3201110132301331-1220203112220033-2121022103321322-3322310212312003-1133112300000111-2012112001331231-3033021103030133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- Examples

<a id="canonical-3002332321021233-0123321130203212-1023212322031120-2132311301123010-3121012210201120-2013123013102233-2022012323223033-0111331033011010"></a>

### Complete configurations for `xcsh_securemesh_site_v2`

- [Resource](resources--securemesh_site_v2--examples--group-001.md#canonical-0023333232013220-2113230021310123-3120301200001322-3032302302200322-0133031223122031-3303032211013301-0202221332221222-3113320232122123): valid configuration.

<a id="canonical-0023333232013220-2113230021310123-3120301200001322-3032302302200322-0133031223122031-3303032211013301-0202221332221222-3113320232122123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Examples](resources--securemesh_site_v2--examples--group-001.md#canonical-1303300012312320-3201110132301331-1220203112220033-2121022103321322-3322310212312003-1133112300000111-2012112001331231-3033021103030133)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_securemesh_site_v2/resource.tf`; digest `sha256:dd55f175aeb63b8fcdedf3ae50d54522d7f389806ca40efe077ad7621a9cf6e2`.

```terraform
# SecuremeshSiteV2 Resource Example
# Manages a Securemesh Site V2 resource in F5 Distributed Cloud for deploying secure mesh edge sites with security and networking controls.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SecuremeshSiteV2 configuration
resource "xcsh_securemesh_site_v2" "example" {
  name      = "example-securemesh-site-v2"
  namespace = "system"
}
```
