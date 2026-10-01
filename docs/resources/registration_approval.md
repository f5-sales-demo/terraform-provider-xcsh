---
page_title: "xcsh_registration_approval landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_registration_approval landing."
---

# xcsh_registration_approval landing

<a id="canonical-7d4fb674027240378875a2202f657eddfb167f6c0f01b29cddfd793bc70315ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0bc1ce806cefb89c40e4db7b414249887cde27bca09b724f485d45ac95b10c02"></a>

## xcsh_registration_approval — xcsh_registration_approval / 19394161374d / 2

Breadcrumbs:

- xcsh_registration_approval

Manages a Registration Approval resource in F5 Distributed Cloud for request for admission approval.
configuration.

<a id="canonical-769610cca4ce6f445cfa104c7d5d871862ccd88141156d21ace5ab5e0c57886c"></a>

## Prerequisites — xcsh_registration_approval / 19394161374d / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-cbf005311e5ba18223552343e23add9da370e7627d02256d3cff889a8b4b68b5"></a>

## Minimal configuration — xcsh_registration_approval / 19394161374d / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# RegistrationApproval Resource Example
# Manages a Registration Approval resource in F5 Distributed Cloud for request for admission approval.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic RegistrationApproval configuration
resource "xcsh_registration_approval" "example" {
  name      = "example-registration-approval"
  namespace = "staging"

  cluster_size = 1
}
```

<a id="canonical-fb8c065a3e1969a7e3233daebef36679008562e56ca096c4c516384ed8c04f86"></a>

## Root configuration — xcsh_registration_approval / 19394161374d / 5

Required root properties: `cluster_size`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-f945f204d11e70d1b1316b0c46eda7cba4604c89edc40cecb5e397f72555fb10"></a>

## Next pages — xcsh_registration_approval / 19394161374d / 6

- [Property reference](../guides/resources--registration_approval--reference--group-001.md#canonical-2950841dcd2dad45aa090a989e2c5989663f4229329496cccde1786e7269c4b0)
- [Examples](../guides/resources--registration_approval--examples--group-001.md#canonical-18983f6c2d60c35fbc6941ad7957980b8a60e933dc081010b07873198b27c179)
- [Import](../guides/resources--registration_approval--lifecycle--group-001.md#canonical-b650ee81672cb726831c9981d7c0a90bd9a72f19657f487688db316d183e0a33)
