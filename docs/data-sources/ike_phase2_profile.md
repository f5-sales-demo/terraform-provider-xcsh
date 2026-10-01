---
page_title: "xcsh_ike_phase2_profile landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase2_profile landing."
---

# xcsh_ike_phase2_profile landing

<a id="canonical-d4f07d0519ebc7b4b0d60c702b18e26168602089d08803fd25362ad1cabf1c5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa88561fa19c36c23fb6f2cbc91720afc506947b869823b2fd5c8176a414a8dd"></a>

## xcsh_ike_phase2_profile — xcsh_ike_phase2_profile / 5d917b86e68e / 2

Breadcrumbs:

- xcsh_ike_phase2_profile

Manages a IKE Phase2 Profile resource in F5 Distributed Cloud for ike phase2 profile specification.
configuration.

<a id="canonical-f296f70745e377512eddc72ccf72afaa2d3fb4ae86aed31fdea983e486bcc1a8"></a>

## Prerequisites — xcsh_ike_phase2_profile / 5d917b86e68e / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-c8b6c6c067de8c2472e66e5a46fc8b34ea3f86fd6a5458e6de137ceb165587f8"></a>

## Minimal configuration — xcsh_ike_phase2_profile / 5d917b86e68e / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# IKEPhase2Profile Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing IKEPhase2Profile by name
data "xcsh_ike_phase2_profile" "example" {
  name      = "example-ike-phase2-profile"
  namespace = "staging"
}

output "ike_phase2_profile_id" {
  value = data.xcsh_ike_phase2_profile.example.id
}
```

<a id="canonical-71b3efc3d5417b1b07e5083c86564304b0727dd5975accb7748b14b7d2e920a6"></a>

## Root configuration — xcsh_ike_phase2_profile / 5d917b86e68e / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-d1c2bc3808f802bc88144992007c1d09f31f7e15427db3e943e45c3221387beb"></a>

## Next pages — xcsh_ike_phase2_profile / 5d917b86e68e / 6

- [Property reference](../guides/data-sources--ike_phase2_profile--reference--group-001.md#canonical-69208af1f987cbde066b5daf25f7fe566428d0259b66669340f64476744b4c9a)
- [Examples](../guides/data-sources--ike_phase2_profile--examples--group-001.md#canonical-59588c09ef53762a3a76590a75aaee307169f6d19df7a933e040c76d7162e176)
