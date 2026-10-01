---
page_title: "xcsh_ike_phase2_profile landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase2_profile landing."
---

# xcsh_ike_phase2_profile landing

<a id="canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1bd166b91b7df326e82b58462a2561f6a0e01ab7179c156c6d321c9777f24dd5"></a>

## xcsh_ike_phase2_profile — xcsh_ike_phase2_profile / 56449a1b1b2f / 2

Breadcrumbs:

- xcsh_ike_phase2_profile

Manages a IKE Phase2 Profile resource in F5 Distributed Cloud for ike phase2 profile specification.
configuration.

<a id="canonical-adabf7cd1962e9d66ff997a9f15dd8c58f38703f5df3d8b933ce6bba63ce7837"></a>

## Prerequisites — xcsh_ike_phase2_profile / 56449a1b1b2f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-68c52bbcd4080ddd2d4b05d918c4c65870499837b554f1ca9b9b79119a124d78"></a>

## Minimal configuration — xcsh_ike_phase2_profile / 56449a1b1b2f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# IKEPhase2Profile Resource Example
# Manages a IKE Phase2 Profile resource in F5 Distributed Cloud for ike phase2 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic IKEPhase2Profile configuration
resource "xcsh_ike_phase2_profile" "example" {
  name      = "example-ike-phase2-profile"
  namespace = "staging"

  authentication_algos = ["example-value"]
  encryption_algos     = ["example-value"]
}
```

<a id="canonical-f07d45318108caa297f723a37a3df7936ed5dea8b54027dc28b2c4702e518bda"></a>

## Root configuration — xcsh_ike_phase2_profile / 56449a1b1b2f / 5

Required root properties: `authentication_algos`, `encryption_algos`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-b0edb46f78ae81415d71f6ed3815a15a381fcc1b4ecc3f7b8f7b19bd0a212dbd"></a>

## Next pages — xcsh_ike_phase2_profile / 56449a1b1b2f / 6

- [Property reference](../guides/resources--ike_phase2_profile--reference--group-001.md#canonical-880ac98477290434e8fa578e0ae87dda91ef4fd852b03e4243db262ef0076126)
- [Examples](../guides/resources--ike_phase2_profile--examples--group-001.md#canonical-d3ab23a37ce3c18427daf6a8491990bcf34c0dec6fe624562c0766270cddc014)
- [Import](../guides/resources--ike_phase2_profile--lifecycle--group-001.md#canonical-b051adbbedf0f476d9ea177b0cb00b02b8b0a1158cb131de12a88e9dcd7339fd)
- [Timeouts](../guides/resources--ike_phase2_profile--lifecycle--group-001.md#canonical-f3163bfba0d5b46ceae75a36dc02051c313d95be0bbc89ec0e4a7ae472d62be4)
