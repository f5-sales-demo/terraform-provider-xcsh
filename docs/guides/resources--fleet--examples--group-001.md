---
page_title: "xcsh_fleet examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet examples."
---

# xcsh_fleet examples

<a id="canonical-f088ffb8e7ae87e96ee1781aef5f7a9d6ced2ed410db553b5a7fe67146c45fa5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-566e8e3e981be6efbcc76740e82f4df002b16297e4fc60abcf31338736fd2736"></a>

## Examples — Examples / d3e9bd3c2eee / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- Examples

<a id="canonical-86bbfb887c8f1ec3ec54043e701a080051564674c054f23036493994e505ed56"></a>

## Complete configurations — Examples / d3e9bd3c2eee / 3

- [Resource](resources--fleet--examples--group-001.md#canonical-a676c245982b241ae7a93bb4badd6fd61581cc38824578ac225735bf746beeba): valid configuration.

<a id="canonical-7bf33d45dafad13c33c355ebafda3799a26f7bb37bde835ef3ab81ae5f751834"></a>

## Next pages — Examples / d3e9bd3c2eee / 4

- [Resource](resources--fleet--examples--group-001.md#canonical-a676c245982b241ae7a93bb4badd6fd61581cc38824578ac225735bf746beeba)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)

<a id="canonical-a676c245982b241ae7a93bb4badd6fd61581cc38824578ac225735bf746beeba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbd03afb92302057989153ede7846a4ed2074bb74b99e07dfb703a68a631cc60"></a>

## Resource — Resource / 5ccac4620a5e / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
- [Examples](resources--fleet--examples--group-001.md#canonical-f088ffb8e7ae87e96ee1781aef5f7a9d6ced2ed410db553b5a7fe67146c45fa5)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_fleet/resource.tf`; digest `sha256:4d32cab1b63bbbd4184a815cf7049069724ae120ad87141da3ed9def32e6d8f3`.

```terraform
# Fleet Resource Example
# Manages fleet will create a fleet object in 'system' namespace of the user in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Fleet configuration
resource "xcsh_fleet" "example" {
  name      = "example-fleet"
  namespace = "staging"

  fleet_label = "example-value"
}
```

<a id="canonical-24722706e97822e41fe2d3621da698b10179471d61c6231cf87f38cb19d14bfc"></a>

## Next pages — Resource / 5ccac4620a5e / 3

- [Examples](resources--fleet--examples--group-001.md#canonical-f088ffb8e7ae87e96ee1781aef5f7a9d6ced2ed410db553b5a7fe67146c45fa5)
- [xcsh_fleet](../resources/fleet.md#canonical-796406566bcae64745ba2b16932161d375fb7e68f997fdc7783547fe8611aa9d)
