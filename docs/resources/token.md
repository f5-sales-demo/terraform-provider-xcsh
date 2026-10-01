---
page_title: "xcsh_token landing"
subcategory: "Identity"
description: "Complete grouped canonical reference for xcsh_token landing."
---

# xcsh_token landing

<a id="canonical-4f14709ee0bf8c6a1f43d547a150dbc9ce3bfe820f4d53be2e91fa13ae4cb3c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebff6a917114a5e0fa10fe5b0a9a69c50c58ec6e5a50788872a526a2c90acb3d"></a>

## xcsh_token — xcsh_token / 3253c5f47f4e / 2

Breadcrumbs:

- xcsh_token

Manages new token. Token object is used to manage site admission. User must generate token before
provisioning and pass this token to site during it's registration in F5 Distributed Cloud.

<a id="canonical-101b87703dd58aee447f9c1da796e22c1c9198d420c0ffe76bfd317bcdbfdeb1"></a>

## Prerequisites — xcsh_token / 3253c5f47f4e / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-472f4070097d336c603a03d3816de6eab3009f804c787c32e2faea24def51dc8"></a>

## Minimal configuration — xcsh_token / 3253c5f47f4e / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Token Resource Example
# Manages new token.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Token configuration
resource "xcsh_token" "example" {
  name      = "example-token"
  namespace = "system"
  type      = 1
  site_name = "example-securemesh-site"
}
```

<a id="canonical-aa33b3c67059451d23633d600c25fb9114c9d3d27134dc0157cb793438420eef"></a>

## Root configuration — xcsh_token / 3253c5f47f4e / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-fcdc5eb5d890a74b64dc4624822d2f05e9e58121b8c799324f16fff60fc4c370"></a>

## Next pages — xcsh_token / 3253c5f47f4e / 6

- [Property reference](../guides/resources--token--reference--group-001.md#canonical-05a33cba933cd1a2d79fe19b098f45465ae3f47fe73462661d6025e7ca495e14)
- [Examples](../guides/resources--token--examples--group-001.md#canonical-6d07d12677cc1288658a0cbde21f944c8cff33dfa0d3d0f5cd2ee61b235e0a92)
- [Import](../guides/resources--token--lifecycle--group-001.md#canonical-ddaae35c5256906174eec7e27442bc009e7a292e1a26e5fb83bda72749f68ffe)
- [Timeouts](../guides/resources--token--lifecycle--group-001.md#canonical-5ce1e18a261d5d1d8dd6dc7609702e439fc1df6a757f6a57be02c932f8c30dce)
