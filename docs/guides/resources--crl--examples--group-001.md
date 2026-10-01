---
page_title: "xcsh_crl examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_crl examples."
---

# xcsh_crl examples

<a id="canonical-f4b8e81c172d26ed2416d6b64fa100e009639d05b196f13650b51e15b65a999e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-850b33c86bd726e386397d407d6ca75bc55636b695cdb27671ff23b90dcd07ab"></a>

## Examples — Examples / 3eab776d4f03 / 2

Breadcrumbs:

- [xcsh_crl](../resources/crl.md#canonical-96a63453b75e823f71de872f51de2a086a64ac241ba947a0a9c61f4a01cb3404)
- Examples

<a id="canonical-528a6edf59cbc67f0b854c5fd4a8bdd4b08fe38d26e70a47535923a1a3b9f2b0"></a>

## Complete configurations — Examples / 3eab776d4f03 / 3

- [Resource](resources--crl--examples--group-001.md#canonical-6abefb38ab636e169493847842caa7982cbff245f67ff2341ee9dc3286c0acf7): valid configuration.

<a id="canonical-049d24c0babb72f447e1135e0f0c4ee6ce55294f1de7602fd40ad93cac625424"></a>

## Next pages — Examples / 3eab776d4f03 / 4

- [Resource](resources--crl--examples--group-001.md#canonical-6abefb38ab636e169493847842caa7982cbff245f67ff2341ee9dc3286c0acf7)
- [xcsh_crl](../resources/crl.md#canonical-96a63453b75e823f71de872f51de2a086a64ac241ba947a0a9c61f4a01cb3404)

<a id="canonical-6abefb38ab636e169493847842caa7982cbff245f67ff2341ee9dc3286c0acf7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-626829411a488befbed2e62f766ec456a018e6514490f20537f160965cc2273f"></a>

## Resource — Resource / a61a0ce78579 / 2

Breadcrumbs:

- [xcsh_crl](../resources/crl.md#canonical-96a63453b75e823f71de872f51de2a086a64ac241ba947a0a9c61f4a01cb3404)
- [Examples](resources--crl--examples--group-001.md#canonical-f4b8e81c172d26ed2416d6b64fa100e009639d05b196f13650b51e15b65a999e)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_crl/resource.tf`; digest `sha256:787cb67d26f54e510168be44dfb1793c35d97676385e49506a4c31f4d8662bf3`.

```terraform
# CRL Resource Example
# Manages a CRL resource in F5 Distributed Cloud for api to create crl object.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CRL configuration
resource "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"

  refresh_interval = 6
  server_address   = "example-value"
  server_port      = 1
  timeout          = 1
}
```

<a id="canonical-629eea6dbf66af9cc2ba960ad2c4858faebc216717899aec717ec8a1982abe30"></a>

## Next pages — Resource / a61a0ce78579 / 3

- [Examples](resources--crl--examples--group-001.md#canonical-f4b8e81c172d26ed2416d6b64fa100e009639d05b196f13650b51e15b65a999e)
- [xcsh_crl](../resources/crl.md#canonical-96a63453b75e823f71de872f51de2a086a64ac241ba947a0a9c61f4a01cb3404)
