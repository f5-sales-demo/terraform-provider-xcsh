---
page_title: "xcsh_namespace examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_namespace examples."
---

# xcsh_namespace examples

<a id="canonical-0123202330031033-0303010003032300-1110000210011033-0332010112301132-1111030222330030-0220310211111100-2310320200233000-2302132023320230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_namespace](../resources/namespace.md#canonical-1131103113212120-1020211231300120-1030023112132023-3013023011120122-0313023322131212-0001202033301111-3031320130100032-1113113021202310)
- Examples

<a id="canonical-2103102222221231-3323103232031322-0231001213132123-3301313212002311-1320003222201120-2230322123231311-3130223002132102-0232112222021212"></a>

### Complete configurations for `xcsh_namespace`

- [Resource](resources--namespace--examples--group-001.md#canonical-1000320232300203-3003333110003100-1310301310102123-3031221121110122-1122112210013322-2320202220323123-0003330023101033-2000233332213300): valid configuration.

<a id="canonical-1000320232300203-3003333110003100-1310301310102123-3031221121110122-1122112210013322-2320202220323123-0003330023101033-2000233332213300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_namespace](../resources/namespace.md#canonical-1131103113212120-1020211231300120-1030023112132023-3013023011120122-0313023322131212-0001202033301111-3031320130100032-1113113021202310)
- [Examples](resources--namespace--examples--group-001.md#canonical-0123202330031033-0303010003032300-1110000210011033-0332010112301132-1111030222330030-0220310211111100-2310320200233000-2302132023320230)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_namespace/resource.tf`; digest `sha256:95871f5bf448147bdd129adaa3690bae9af907a5a988941728cc5c87a1b9e514`.

```terraform
# Namespace Resource Example
# Manages new namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Credentials are supplied externally.
provider "xcsh" {}

# Basic Namespace configuration
resource "xcsh_namespace" "this" {
  name = "example-namespace"
}
```
