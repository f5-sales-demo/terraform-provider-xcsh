---
page_title: "xcsh_ike_phase2_profile examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase2_profile examples."
---

# xcsh_ike_phase2_profile examples

<a id="canonical-d3ab23a37ce3c18427daf6a8491990bcf34c0dec6fe624562c0766270cddc014"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff17e49209ecaeaf5926597aac30b02b430a622060367e769d89779cbdb81c22"></a>

## Examples — Examples / e54aec477cf0 / 2

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)
- Examples

<a id="canonical-8d40da32e5fe13e31a828beba06a93ac749ecffdfb51790bebdc918ed34b89d3"></a>

## Complete configurations — Examples / e54aec477cf0 / 3

- [Resource](resources--ike_phase2_profile--examples--group-001.md#canonical-68b6d18c8c8e8abbdf994f88000c0e61def3bc03be4b8861a4bfc9e89131b3e6): valid configuration.

<a id="canonical-c3536568eddac5c63938d3b60597dfa91f518df1831f18be09e19383aee887d8"></a>

## Next pages — Examples / e54aec477cf0 / 4

- [Resource](resources--ike_phase2_profile--examples--group-001.md#canonical-68b6d18c8c8e8abbdf994f88000c0e61def3bc03be4b8861a4bfc9e89131b3e6)
- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)

<a id="canonical-68b6d18c8c8e8abbdf994f88000c0e61def3bc03be4b8861a4bfc9e89131b3e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f09c9acde82575e6a9ae311130e332d742a05f2d451b893894d624c75834a97d"></a>

## Resource — Resource / da511eb715de / 2

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)
- [Examples](resources--ike_phase2_profile--examples--group-001.md#canonical-d3ab23a37ce3c18427daf6a8491990bcf34c0dec6fe624562c0766270cddc014)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_ike_phase2_profile/resource.tf`; digest `sha256:80e0e1a4beab03b8fb465c2b172d0e27a9d0b25515c600ec2364548afb3d3ffd`.

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

<a id="canonical-8953151b5da4c26175ba23f2c5402dcb1a937b2bb6e0084e76575919eeb930be"></a>

## Next pages — Resource / da511eb715de / 3

- [Examples](resources--ike_phase2_profile--examples--group-001.md#canonical-d3ab23a37ce3c18427daf6a8491990bcf34c0dec6fe624562c0766270cddc014)
- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)
