---
page_title: "xcsh_authentication examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authentication examples."
---

# xcsh_authentication examples

<a id="canonical-aa08cf8ca5220ee6b47053b20724624aa580a5cf67d8719f0af153aac6c46876"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b121801dc01900e36503a75f6df4dca44bc98d1f3787cd4c0d629f9ef8e1ed20"></a>

## Examples — Examples / 50d06795eccc / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- Examples

<a id="canonical-e7238d45f0b5214d10c14442caabcefddf6d3f1acad24562f3e2234934e3d4b9"></a>

## Complete configurations — Examples / 50d06795eccc / 3

- [Resource](resources--authentication--examples--group-001.md#canonical-d257f9c4e7592ed1b23b11bccfee9c4ed1571634b9a4c78e200f62a71b839360): valid configuration.

<a id="canonical-b6ab6c419ebfb590914f3c57eb4d120ab23063692447ca94ddf7655565c95c3e"></a>

## Next pages — Examples / 50d06795eccc / 4

- [Resource](resources--authentication--examples--group-001.md#canonical-d257f9c4e7592ed1b23b11bccfee9c4ed1571634b9a4c78e200f62a71b839360)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)

<a id="canonical-d257f9c4e7592ed1b23b11bccfee9c4ed1571634b9a4c78e200f62a71b839360"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4bbabf49b556cde8024d28c797820ceb41a08fc99255a7893343203afc5976b2"></a>

## Resource — Resource / bebc6d1a2e7b / 2

Breadcrumbs:

- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
- [Examples](resources--authentication--examples--group-001.md#canonical-aa08cf8ca5220ee6b47053b20724624aa580a5cf67d8719f0af153aac6c46876)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_authentication/resource.tf`; digest `sha256:43532e046dd0e813a92a49d472a5eaa915d3ad0e96dcd6b49097f16328170876`.

```terraform
# Authentication Resource Example
# Manages a Authentication resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Authentication configuration
resource "xcsh_authentication" "example" {
  name      = "example-authentication"
  namespace = "staging"
}
```

<a id="canonical-ad6e6c2a304c0b1cb66a9784296dac5f1c2841760bdb5676290bef314136128f"></a>

## Next pages — Resource / bebc6d1a2e7b / 3

- [Examples](resources--authentication--examples--group-001.md#canonical-aa08cf8ca5220ee6b47053b20724624aa580a5cf67d8719f0af153aac6c46876)
- [xcsh_authentication](../resources/authentication.md#canonical-417d73493652f5a1771ef46f280f23ffd5c91667377a08f64262e5ac9ba40724)
