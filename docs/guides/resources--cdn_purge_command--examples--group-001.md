---
page_title: "xcsh_cdn_purge_command examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_purge_command examples."
---

# xcsh_cdn_purge_command examples

<a id="canonical-67e622085485d9bee851ad3e118817e9c995ea98ce48604db20bdd0c540bacfe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9194b732190b06997028b89b0875b143af5ff3dd65e595c2f9cd95ded12a48df"></a>

## Examples — Examples / 9db5d7c1b862 / 2

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0b2a857dd2a4345307f28b391793f77c7f11e893b8493fb53ba1965b5e2bf425)
- Examples

<a id="canonical-6c48341f3058bdf4d903bd7b0df2f8a1102b5a7ee14b6ec7687b4092c29e0a5a"></a>

## Complete configurations — Examples / 9db5d7c1b862 / 3

- [Resource](resources--cdn_purge_command--examples--group-001.md#canonical-ff8894c9b4674a71ef6db4538df097bb0a8f0218e1a5e529e44e220864a33416): valid configuration.

<a id="canonical-eaf37aa3b35dc2109d8e759ed61115fdb3aa4aab537bcb84389758dde768f39a"></a>

## Next pages — Examples / 9db5d7c1b862 / 4

- [Resource](resources--cdn_purge_command--examples--group-001.md#canonical-ff8894c9b4674a71ef6db4538df097bb0a8f0218e1a5e529e44e220864a33416)
- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0b2a857dd2a4345307f28b391793f77c7f11e893b8493fb53ba1965b5e2bf425)

<a id="canonical-ff8894c9b4674a71ef6db4538df097bb0a8f0218e1a5e529e44e220864a33416"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d22444e7e9f8a4681cecde439ff36083bb34c6e36a3bae1f37bb994bda401390"></a>

## Resource — Resource / a5932c70b32f / 2

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0b2a857dd2a4345307f28b391793f77c7f11e893b8493fb53ba1965b5e2bf425)
- [Examples](resources--cdn_purge_command--examples--group-001.md#canonical-67e622085485d9bee851ad3e118817e9c995ea98ce48604db20bdd0c540bacfe)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cdn_purge_command/resource.tf`; digest `sha256:07fdb3a832901b7ce2b8f7a3292052a09eb46eebbc60fc38d6035bae249459aa`.

```terraform
# CDNPurgeCommand Resource Example
# Manages a CDN Purge Command resource in F5 Distributed Cloud for cdn purge command specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNPurgeCommand configuration
resource "xcsh_cdn_purge_command" "example" {
  name      = "example-cdn-purge-command"
  namespace = "staging"
}
```

<a id="canonical-fcc4090d76253cd7c01c2e418574b92a8566b4bf1b9964a625d3143bd3df83ed"></a>

## Next pages — Resource / a5932c70b32f / 3

- [Examples](resources--cdn_purge_command--examples--group-001.md#canonical-67e622085485d9bee851ad3e118817e9c995ea98ce48604db20bdd0c540bacfe)
- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0b2a857dd2a4345307f28b391793f77c7f11e893b8493fb53ba1965b5e2bf425)
