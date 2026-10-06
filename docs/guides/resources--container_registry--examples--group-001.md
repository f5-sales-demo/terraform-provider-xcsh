---
page_title: "xcsh_container_registry examples"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_container_registry examples."
---

# xcsh_container_registry examples

<a id="canonical-0132120323120123-0223213303233201-1021012301310010-0103313121312020-3001233203211120-1133301233201321-3003031002201103-2201223212203310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_container_registry](../resources/container_registry.md#canonical-3121031110330001-1021233032301331-1200313112012333-2030023322330002-2021231201222110-1023032102203110-1103221200322013-1102022332013332)
- Examples

<a id="canonical-3220321003012303-3013213333033011-2133302031220310-1232101133302112-2222111132222313-2213003333033030-2310100301112232-1231300111110112"></a>

### Complete configurations for `xcsh_container_registry`

- [Resource](resources--container_registry--examples--group-001.md#canonical-2130301231230320-1301231032231233-0032311020313122-3013012322020212-0130212101221232-0012132331221100-0232022323323110-2033011011101232): valid configuration.

<a id="canonical-2130301231230320-1301231032231233-0032311020313122-3013012322020212-0130212101221232-0012132331221100-0232022323323110-2033011011101232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_container_registry](../resources/container_registry.md#canonical-3121031110330001-1021233032301331-1200313112012333-2030023322330002-2021231201222110-1023032102203110-1103221200322013-1102022332013332)
- [Examples](resources--container_registry--examples--group-001.md#canonical-0132120323120123-0223213303233201-1021012301310010-0103313121312020-3001233203211120-1133301233201321-3003031002201103-2201223212203310)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_container_registry/resource.tf`; digest `sha256:333d8100537c966abe2c10d8b88b6ea4b034dc0115a736110297064d1894694d`.

```terraform
# ContainerRegistry Resource Example
# Manages a Container Registry resource in F5 Distributed Cloud for container image registry configuration.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ContainerRegistry configuration
resource "xcsh_container_registry" "example" {
  name      = "example-container-registry"
  namespace = "staging"

  registry  = "example-value"
  user_name = "example-value"
}
```
